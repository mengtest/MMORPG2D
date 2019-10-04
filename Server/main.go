package main

import (
	"log"
	"net"

	msf "./core"
)

var ser = msf.NewMsf(&msf.CommSocket{})

type event struct {
}

func (this event) OnHandel(fd uint32, conn net.Conn) bool {
	log.Println(fd, "链接成功类")
	return true
}

func (this event) OnMessage(fd uint32, msg map[string]string) bool {
	log.Println("这是一个接受消息事件", msg)
	return true
}

type Test struct {
}

func (this Test) Default(fd uint32, data map[string]string) bool {
	log.Println("default")
	return true
}

func (this Test) BeforeRequest(fd uint32, data map[string]string) bool {
	log.Println("before")
	return true
}

func (this Test) AfterRequest(fd uint32, data map[string]string) bool {
	log.Println("after")
	return true
}

func (this Test) Hello(fd uint32, data map[string]string) bool {
	log.Println("收到消息了")
	log.Println(data)
	ser.SessionMaster.WriteByid(fd, []byte("hehehehehehehehe"))
	return true
}

//----------------------------------------------------------------------------------------------------------------------
func main() {
	log.SetFlags(log.Lshortfile | log.LstdFlags | log.Llongfile)
	ser.EventPool.RegisterEvent(&event{})
	ser.EventPool.RegisterStructFun("test", &Test{})
	ser.Listening(":8565")
}
