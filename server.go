package main

import (
	"fmt"
	"net"
	"sync"
)

type Server struct {
	Ip   string
	Port int

	//在线用户列表
	OnlineMap map[string]*User
	mapLock   sync.RWMutex

	//消息广播的channel
	Message chan string
}

// server接口
func NewServer(ip string, port int) *Server {
	server := &Server{
		Ip:        ip,
		Port:      port,
		OnlineMap: make(map[string]*User),
		Message:   make(chan string),
	}
	return server
}

// 监听Message广播小西channel的goroutine，一旦有消息就把所有消息传递给在线用户
func (t *Server) ListenMessager() {
	for {
		msg := <-t.Message

		//将消息发送给所有User
		t.mapLock.Lock()
		for _, cli := range t.OnlineMap {
			cli.C <- msg
		}
		t.mapLock.Unlock()
	}
}

// 广播消息的方法
func (t *Server) Broadcast(User *User, msg string) {
	sendmsg := "[" + User.Addr + "]" + User.Name + ":" + msg
	t.Message <- sendmsg
}

func (t *Server) Handler(conn net.Conn) {
	//当前业务
	fmt.Println("链接建立成功... ")

	user := NewUser(conn)

	//用户上线，将用户加入到OnlineMap中
	t.mapLock.Lock()
	t.OnlineMap[user.Name] = user
	t.mapLock.Unlock()

	//广播用户上线信息
	t.Broadcast(user, "上线")

	//当前handler阻塞
	select {}
}

// 启动服务器
func (t *Server) Start() {
	//socket监听
	listener, err := net.Listen("tcp", fmt.Sprintf("%s:%d", t.Ip, t.Port))
	if err != nil {
		fmt.Println("net.listen err:", err)
		return
	}
	//关闭监听
	defer listener.Close()

	//启动监听masage的goroutine
	go t.ListenMessager()

	for {
		//如果监听成功，进入一个无限循环，不断接受新的连接。
		connect, err := listener.Accept()
		if err != nil {
			fmt.Println("listerner.accept err:", err)
			continue
		}
		//对于每个新的连接，它会使用go关键字来异步执行Handerler方法。
		go t.Handler(connect)
	}
}
