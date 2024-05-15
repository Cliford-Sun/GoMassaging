package main

import (
	"fmt"
	"net"
)

type Server struct {
	Ip   string
	Port int
}

// server接口
func NewServer(ip string, port int) *Server {
	server := &Server{
		Ip:   ip,
		Port: port,
	}
	return server
}

func (t *Server) Handerler(conn net.Conn) {
	//当前业务
	fmt.Println("链接建立成功... ")
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

	for {
		//如果监听成功，进入一个无限循环，不断接受新的连接。
		connect, err := listener.Accept()
		if err != nil {
			fmt.Println("listerner.accept err:", err)
			continue
		}
		//对于每个新的连接，它会使用go关键字来异步执行Handerler方法。
		go t.Handerler(connect)

	}
}
