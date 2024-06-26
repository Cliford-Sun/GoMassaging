package main

import (
	"fmt"
	"io"
	"net"
	"sync"
	"time"
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

// 监听Message广播消息channel的goroutine，一旦有消息就把所有消息传递给在线用户
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
	user := NewUser(conn, t)
	//fmt.Println(user.Name, "链接建立成功... ")
	//用户上线
	user.Online()

	//判断用户是否活跃
	isLive := make(chan bool)

	go func() {
		buf := make([]byte, 4096)
		for {
			n, err := conn.Read(buf)
			if n == 0 {
				user.Offline()
				return
			}
			if err != nil && err != io.EOF {
				fmt.Println("conn Read err:", err)
				return
			}
			//提取用户的消息（去除\n）
			msg := string(buf[:n-1])

			//将的到的消息广播：
			user.Domessage(msg)

			//用户活跃
			isLive <- true

		}
	}()

	//当前handler阻塞
	for {
		select {
		case <-isLive:
			//当前用户是活跃的，重置计时器
			//不做任何事情，为了激活select，更新下面的计时器
		case <-time.After(time.Second * 300):
			//已经超时
			//将当前user强制踢出

			user.Sendmsg("你因不活跃被踢出聊天室")

			//销毁资源
			close(user.C)

			conn.Close()

			//退出房前handler
			return
		}
	}
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
