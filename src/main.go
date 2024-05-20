package main

func main() {
	//可以通过telnet localhost 8888来实现链接，localhost通常是指本地机器的回环地址，也就是127.0.0.1
	server := NewServer("127.0.0.1", 8888)
	server.Start()
}
