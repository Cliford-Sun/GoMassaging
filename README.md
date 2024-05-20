# GoMassaging

golang学习后的初次项目实战，也是第一次做项目，为了加强自身的理解

下面为课程链接：

[https://www.bilibili.com/video/BV1gf4y1r79E?p=1&amp;vd_source=082140fe62c9e441bcd547ba7e3ae6ff](https://www.bilibili.com/video/BV1gf4y1r79E?p=1&vd_source=082140fe62c9e441bcd547ba7e3ae6ff%E2%80%B8)

---

**下面就是go语言webserver服务器的基本框架**
框架利用了go语言的goroutine和channel的机制，实现了高并发的WebSocket服务器，用于处理多用户实时通信。
![image](https://github.com/Cliford-Sun/GoMassaging/blob/main/graph/WebServer.png)

服务器接收消息后，通过遍历 OnlineMap，将消息放入每个在线用户的消息通道，用户goroutine从通道读取消息并发送给对应的客户端，实现消息的广播。

---

**实现了基本的即使通信聊天功能**
用户加入聊天时候有广播欢迎
直接输入内容可以向所有用户广播内容
输入 `who`能够实现列出所有的在线用户
输入 `rename|newname`就可以将用户名字改成*newname*
输入 `to|username|message`可以对用户*uername*私聊发送消息*message*
当用户在聊天室内不活跃30秒后将会被强制踢出

**在实现上面聊天功能的基础上实现了客户端的建立**
在运行bin文件夹中的server.exe和client.exe后,会有目录提示选择模式,根据不同的模式来实现公聊,私聊,更新用户名等功能
