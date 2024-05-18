package main

import (
	"net"
	"strings"
)

type User struct {
	Name string
	Addr string
	C    chan string
	conn net.Conn

	server *Server
}

// 创建用户api
func NewUser(conn net.Conn, server *Server) *User {
	userAddr := conn.RemoteAddr().String()

	user := &User{
		Name: userAddr,
		Addr: userAddr,
		C:    make(chan string),
		conn: conn,

		server: server,
	}
	//启动监听当前user channel消息的构成
	go user.ListenMessage()
	return user
}

func (t *User) Online() {
	//用户上线，将用户加入到OnlineMap中
	t.server.mapLock.Lock()
	t.server.OnlineMap[t.Name] = t
	t.server.mapLock.Unlock()

	//广播用户上线信息
	t.server.Broadcast(t, "上线")

}

func (t *User) Offline() {
	//用户下线，将用户从OnlineMap中删除
	t.server.mapLock.Lock()
	delete(t.server.OnlineMap, t.Name)
	t.server.mapLock.Unlock()

	//广播用户上线信息
	t.server.Broadcast(t, "下线")
}

// 给当前用户对应的客户端发消息
func (t *User) Sendmsg(msg string) {
	t.conn.Write([]byte(msg))
}

func (t *User) Domessage(msg string) {

	if msg == "who" {
		//当前在线用户的查询
		t.server.mapLock.Lock()
		for _, user := range t.server.OnlineMap {
			onlineMsg := "[" + user.Addr + "]" + user.Name + ":" + "在线...\n"
			t.Sendmsg(onlineMsg)
		}
		t.server.mapLock.Unlock()
	} else if len(msg) > 7 && msg[:7] == "rename|" {
		newName := strings.Split(msg, "|")[1]
		//判断name是否已经存在
		_, ok := t.server.OnlineMap[newName]
		if ok {
			t.Sendmsg("当前用户名已经存在\n")
		} else {
			t.server.mapLock.Lock()
			delete(t.server.OnlineMap, t.Name)
			t.server.OnlineMap[newName] = t
			t.server.mapLock.Unlock()

			t.Name = newName
			t.Sendmsg("已修改用户名为" + newName + "\n")
		}
	} else {
		t.server.Broadcast(t, msg)
	}

}

// 监听当前的User channel的方法，一旦有消息就发送给客户端
func (t *User) ListenMessage() {
	for {
		msg := <-t.C
		t.conn.Write([]byte(msg + "\n"))
	}
}
