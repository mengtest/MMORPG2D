package main

import (
	"fmt"
	"log"
	"net"
)

//定义结构体全局map的value值，包括每一个用户的姓名，ip地址和私人管道
type client struct {
	name string
	addr string
	c    chan string
}

func writeMsg2Client(client client, conn net.Conn) {
	for m := range client.C {
		conn.Write([]byte(m + "\n"))
	}
}

func makeMsg(name string, addr string, s string) string {
	return "[" + addr + "]" + name + s
}

func handleconn(conn net.Conn) {
	defer conn.Close()

	addr := conn.RemoteAddr().String()
	fmt.Printf("用户%s进入了房间\n", addr)

	client := client{addr, addr, make(chan string)}

	go writeMsg2Client(client, conn)
	onlienmap[addr] = client

	Message <- makeMsg(client.name, addr, "login")

	haschat := make(chan bool)
	ifquit := make(chan bool)
	var flag bool

	go func(){
		buf := make([]byte,4096)
		for{
			n,err := conn.Read(buf)
			if n == 0{
				fmt.Printf("%s离开了房间\n",client.name)
				ifquit<-true
				return
			}

			if string(buf[:7])=="Rename|"{
				client.name = strings.Split(string(buf[:n-1]),"|")[1]
				onlienmap[addr] = client
				conn.Write([]byte("rename success\n"))
			}else if string(buf[:n-1])=="/who"{
				for _,s := range onlinemap{
					conn.Write([]byte(s.name+"online\n"))
				}
			}else if string(buf[:2])=="m|"&&strings.Count(string(buf[:n]),"|")==2{
				//私聊功能实现
				flag =false
				slice := strings.Split(string(buf))
			}

		}
	}

}

func main() {
	// serverSocket :=
	Listening()
}

func Listening() {
	tcpListen, err := net.Listen("tcp", ":8565")
	// msg := "hello client! 你好..."

	if err != nil {
		panic(err)
	}

	for {
		conn, err := tcpListen.Accept()
		if err != nil {
			log.Println(err)
			continue
		}
		go connHandle(conn)
	}
}
func connHandle(conn net.Conn) {
	defer conn.Close()
	readBuff := make([]byte, 14)
	for {
		n, err := conn.Read(readBuff)
		if err != nil {
			return
		}
		fmt.Println(readBuff[:n])

	}
}
