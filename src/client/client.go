package main

import (
	"flag"
	"fmt"
	"io"
	"net"
	"os"
)

type Client struct {
	ServerIP   string
	ServerPort int
	Name       string
	conn       net.Conn
	flag       int
}

func NewClient(serverIP string, serverPort int) *Client {
	// 创建客户端对象
	client := &Client{
		ServerIP:   serverIP,
		ServerPort: serverPort,
		flag:       999,
	}

	// 链接server
	conn, err := net.Dial("tcp", fmt.Sprintf("%s:%d", serverIP, serverPort))
	if err != nil {
		fmt.Println("net.Dial error: ", err)
		return nil
	}
	client.conn = conn

	// 返回对象
	return client
}

func (t *Client) PublicChat() {
	var msg string
	fmt.Println(">>>>>>请输入聊天内容,输入exit退出公聊模式")
	fmt.Scanln(&msg)
	for msg != "exit" {
		//发送给服务器

		//消息不为空则发送
		if len(msg) != 0 {
			sendmsg := msg + "\n"
			_, err := t.conn.Write([]byte(sendmsg))
			if err != nil {
				fmt.Println("conn Write error:", err)
				break
			}
		}

		msg = ""
		fmt.Println(">>>>>>请输入聊天内容,输入exit退出公聊模式")
		fmt.Scanln(&msg)
	}
}

func (t *Client) SelectUsers() {
	sendmsg := "who\n"

	_, err := t.conn.Write([]byte(sendmsg))
	if err != nil {
		fmt.Println("conn Write error:", err)
		return
	}
}

func (t *Client) PrivateChat() {
	var remoteName string
	var msg string

	t.SelectUsers()
	fmt.Println(">>>>>>请输入私聊对象[用户名],exit退出私聊模式")
	fmt.Scanln(&remoteName)

	for remoteName != "exit" {
		fmt.Println(">>>>>>请输入消息内容,exit退出消息内容输入")
		fmt.Scanln(&msg)

		for msg != "exit" {
			//消息不为空则发送
			if len(msg) != 0 {
				sendmsg := "to|" + remoteName + "|" + msg + "\n"
				_, err := t.conn.Write([]byte(sendmsg))
				if err != nil {
					fmt.Println("conn Write error:", err)
					break
				}
			}
			msg = ""
			fmt.Println(">>>>>>请输入消息内容,exit退出消息内容输入")
			fmt.Scanln(&msg)
		}
		t.SelectUsers()
		fmt.Println(">>>>>>请输入私聊对象[用户名],exit退出私聊模式")
		fmt.Scanln(&remoteName)
	}
}

func (t *Client) UpdateUserName() bool {
	fmt.Println(">>>>>>请输入新用户名:")
	fmt.Scanln(&t.Name)

	sendmsg := "rename|" + t.Name + "\n"
	_, err := t.conn.Write([]byte(sendmsg))
	if err != nil {
		fmt.Println("conn Write error: ", err)
		return false
	}
	return true
}

// 处理server回应的消息内容
func (t *Client) DealResponse() {
	//一旦t.conn有数据,就copy到stdout中,是永久阻塞监听
	io.Copy(os.Stdout, t.conn)
}

// 显示菜单
func (t *Client) menu() bool {
	var flag int
	fmt.Println("1.公聊模式")
	fmt.Println("2.私聊模式")
	fmt.Println("3.更新用户名")
	fmt.Println("0.退出")

	fmt.Scanln(&flag)
	if flag >= 0 && flag <= 3 {
		t.flag = flag
		return true
	} else {
		fmt.Println("输入格式不正确,请输入合法范围内的数字")
		return false
	}
}

func (t *Client) Run() {
	for t.flag != 0 {
		for !t.menu() {
		}
		//根据不同模式执行不同业务
		switch t.flag {
		case 1:
			//公聊
			t.PublicChat()
		case 2:
			//私聊
			t.PrivateChat()
		case 3:
			//更新用户名
			t.UpdateUserName()

		}
	}
}

var serverIp string
var serverPort int

func init() {
	flag.StringVar(&serverIp, "ip", "127.0.0.1", "设置服务器IP地址(默认是127.0.0.1)")
	flag.IntVar(&serverPort, "port", 8888, "设置服务器port端口(默认是8888)")
}

func main() {
	//命令行解析
	flag.Parse()

	client := NewClient(serverIp, serverPort)
	if client == nil {
		fmt.Println(">>>>>>链接服务器失败>>>>>>")
		return
	}

	//开启一个go程回应server的消息
	go client.DealResponse()

	fmt.Println(">>>>>>链接服务器成功>>>>>>")

	//执行业务
	client.Run()
}
