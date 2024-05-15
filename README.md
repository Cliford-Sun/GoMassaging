# GoMassaging

golang学习后的初次项目实战，也是第一次做项目，为了加强自身的理解

下面为课程链接：

[https://www.bilibili.com/video/BV1gf4y1r79E?p=1&amp;vd_source=082140fe62c9e441bcd547ba7e3ae6ff](https://www.bilibili.com/video/BV1gf4y1r79E?p=1&vd_source=082140fe62c9e441bcd547ba7e3ae6ff%E2%80%B8)

---

## 1.实现了server的基本架构
---

## 2.实现了用户上线与广播的功能：

**下面就是go语言webserver服务器的基本框架**

框架利用了go语言的goroutine和channel的机制，实现了高并发的WebSocket服务器，用于处理多用户实时通信。

![image](https://github.com/JSmikasa/GoMassaging/blob/main/image/README/1715783867674.png)

**OnlineMap**：

用来存储所有在线用户。每个用户都有一个名字和对应的用户对象。

**用户goroutine**：

每个在线用户都有一个独立的goroutine。这些goroutine负责处理与客户端的连接，并通过write操作向客户端发送消息。

**Channel**:

每个用户goroutine都与一个消息通道关联。通道用于在服务器和用户goroutine之间传递消息。

当服务器收到一条消息时，它会通过通道将消息传递给对应的用户goroutine。

**消息传递流程** :

服务器接收到一条消息后，消息会通过相关的通道发送给对应的用户goroutine。

用户goroutine从通道中读取消息，并通过连接将消息写入到客户端。

**服务端与客户端的关系**

服务器管理在线用户，并通过通道和goroutine与每个客户端进行通信。

每个客户端与一个用户goroutine关联，通过连接接收服务器发送的消息。


总的来说，这个架构利用Go语言的goroutine和channel机制，实现了一个高并发的WebSocket服务器，用于处理多用户实时通信。
