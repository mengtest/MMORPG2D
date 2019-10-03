package main

import(
	"net"
	"log"
	"fmt"
)

func main(){
	Listening()
}

func Listening(){
	//调用net.listen传入 网络类型和ip端口，返回一个listener
	tcpListen,err := net.Listen("tcp",":8565")

	if err != nil{
		panic(err)
	}

	//进入一个无限循环
	for{
		//执行listener.Accept()方法，这里是阻塞的，
		// 当有客户端连接时返回一个conn对象
		conn,err:=tcpListen.Accept()
		if err !=nil{
			log.Println(err)
			continue
		}
		// 每当一个客户端连接成功，就单独开一个协程处理连接
		go connHandle(conn)
	}
}
func connHandle(conn net.Conn){
	// 确保协程退出时候关闭
	defer conn.Close()
	readBuff := make([]byte,14)
	for{
		n,err:=conn.Read(readBuff)
		if err !=nil{
			return
		}
		// 读取客户端信息，注意这里是阻塞的！
		fmt.Println(readBuff[:n])
	}
}