package main

import (
	"fmt"
	"goClient/util"
	"net"
)

func testClient() {
	// 调用net包中的dial 传入ip 端口 进行拨号连接，通过三次握手之后获取到conn
	conn, err := net.Dial(util.Server_NetWorkType, util.Server_Address)
	if err != nil{
        fmt.Println("Client create conn error err:", err)
	}
	defer conn.Close()
	 //往服务端传递消息
	 util.Write(conn, "aaaa")
	//读取服务端返回的消息
	if str, err := util.Read(conn); err == nil {
		fmt.Println(str)
	}
}

func main() {
	fmt.Println("----------------Client!")

	testClient()

	//接受命令行输入，不做任何事情
    var input string
    fmt.Scanln(&input)
}
