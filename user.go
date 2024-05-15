package main

import "net"

type User struct {
	Name string
	Addr string
	C    chan string
	conn net.Conn
}

//创建用户api
func NewUser(conn net.Conn) *User {
	userAddr := conn.RemoteAddr().String()

	user := &User{
		Name: userAddr,
		Addr: userAddr,
		C:    make(chan string),
		conn: conn,
	}
	//启动监听当前user channel消息的构成
	go user.ListenMessage()
	return user
}

//监听当前的User channel的方法，一旦有消息就发送给客户端
func (t *User) ListenMessage() {
	for {
		msg := <-t.C
		t.conn.Write([]byte(msg + "\n"))
	}
}
