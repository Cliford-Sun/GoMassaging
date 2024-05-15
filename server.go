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
func (t *Server) Strat() {
	//socket监听
	listener, err := net.Listen("tcp", fmt.Sprintf("%s:%d", t.Ip, t.Port))
	if err != nil {
		fmt.Println("net.listen err:", err)
		return
	}
	//关闭监听
	defer listener.Close()

	for {
		//接收
		connect, err := listener.Accept()
		if err != nil {
			fmt.Println("listerner.accept err:", err)
			continue
		}
		//do hendler
		go t.Handerler(connect)

	}

	//close listen socket

}
