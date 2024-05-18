# GoMassaging

golang学习后的初次项目实战，也是第一次做项目，为了加强自身的理解

下面为课程链接：

[https://www.bilibili.com/video/BV1gf4y1r79E?p=1&amp;vd_source=082140fe62c9e441bcd547ba7e3ae6ff](https://www.bilibili.com/video/BV1gf4y1r79E?p=1&vd_source=082140fe62c9e441bcd547ba7e3ae6ff%E2%80%B8)

---

## v0.1：实现了server的基本架构

---

## v0.2：实现了用户上线与广播的功能

**下面就是go语言webserver服务器的基本框架**
框架利用了go语言的goroutine和channel的机制，实现了高并发的WebSocket服务器，用于处理多用户实时通信。
![image](https://github.com/Cliford-Sun/GoMassaging/blob/main/WebServer.png)

服务器接收消息后，通过遍历 OnlineMap，将消息放入每个在线用户的消息通道，用户goroutine从通道读取消息并发送给对应的客户端，实现消息的广播。

---

## v0.3：实现了用户消息的广播

完善了handler处理业务的方法，启动了一个针对客户端的读goroutine

---

## v0.4：实现用户业务层的封装，添加用户的上线，下线，广播用户消息的功能到用户部分
