package main

import (
	"fmt"
	"goServer/util"
	"net"
	"time"
)

// func testMySql() {
// 	opend, db := util.OpenDB()
//     if !opend {
// 		fmt.Println("open faile")
// 		return
//     }
// 	// util.InsertToDB(db)
// 	// util.DeleteFromDB(db, 1)
// 	// util.UpdateDB(db, "924ba755021103855ad73dc5faea9de2")
// 	util.QueryFromDB(db)
// }

func testSocketServer() {
	// net listen 函数 传入socket类型和ip端口，返回监听对象
	listener, err := net.Listen(util.Server_NetWorkType, util.Server_Address)
	if err == nil {
		defer listener.Close()
		// 循环等待客户端访问
		for {
			conn, err := listener.Accept()
			if err == nil {
				// 一旦有外部请求，并且没有错误 直接开启异步执行
				go handleConn(conn)
			}
		}
	} else {
		fmt.Println("server error", err)
	}
}

func handleConn(conn net.Conn) {
	for {
		// 设置读取超时时间
		conn.SetReadDeadline(time.Now().Add(time.Second * 2))
		// 调用公用方法read 获取客户端传过来的消息。
		if str, err := util.Read(conn); err == nil {
			fmt.Println("client:", conn.RemoteAddr(), str)
			// 通过write 方法往客户端传递一个消息
			util.Write(conn, "server got:" + str)
		}
	}
}

func main() {
	fmt.Println("----------------Server!")

	// testMySql()
	testSocketServer()

	//接受命令行输入，不做任何事情
    var input string
    fmt.Scanln(&input)
}
