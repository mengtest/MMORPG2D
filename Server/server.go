package main

import(
	"net"
	"log"
	"fmt"
)

const(
	CONSTHEADER       = "testHeader"
    CONSTHEADERLENGTH = 10
    CONSTMLENGTH      = 4
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

	var errs error
	tempBuff := make([]byte,0)
	readBuff := make([]byte,14)
	data := make([]byte,20)

	for{
		n,err:=conn.Read(readBuff)
		if err !=nil{
			return
		}
		tempBuff = append(tempBuff,readBuff[:n]...)
		// 对缓存器进行分包处理
		tempBuff,data,errs = Depack(tempBuff)
		if errs != nil {
			return
		}

		if len(data)==0 {
			return
		}
		// 读取客户端信息，注意这里是阻塞的！
		fmt.Println(readBuff[:n])
		//伪代码data是一个单独的包，可以进行逻辑处理了
		// do(data)
	}
}

func Enpack(message []byte) []byte{
	return append(append([]byte(CONSTHEADER),IntToBytes(len(message))...),message...)
}

func Depack(buff[]byte)([]byte,[]byte,error){
	length :=len(buff)

	// 如果包长小于header 就直接返回，因为接受的数据不完整
	if length<CONSTHEADERLENGTH+CONSTMLENGTH{
		return buff,nil,nil
	}

	// 如果header不是指定的header说明数据已经被污染，则直接返回
	if string(buff[:CONSTHEADERLENGTH])!=CONSTHEADER {
		return []byte{},nil,error.New("header is not safe")
	}
	
	msgLength := BytesToInt(buff[CONSTMLENGTH:CONSTHEADERLENGTH+CONSTMLENGTH])
	if length < CONSTHEADERLENGTH + CONSTMLENGTH + msgLength{
		return buff,nil,nil
	}

	data := buff[CONSTHEADERLENGTH + CONSTMLENGTH : CONSTHEADERLENGTH + CONSTMLENGTH + msgLength]
	buffs := buff[CONSTHEADERLENGTH + CONSTMLENGTH + msgLength:]
	return buffs,data,nil
}

//将int转成四个字节
func IntToBytes(n int) []byte {
    x := int32(n)
    bytesBuffer := bytes.NewBuffer([]byte{})
    binary.Write(bytesBuffer, binary.BigEndian, x)
    return bytesBuffer.Bytes()
}
//将四个字节转成int
func BytesToInt(b []byte)int {
	bytesBuffer := bytes.NewBuffer(b)
	var x int32
	binary.Read(bytesBuffer,binary.BigEndian,&x)
	return int(x)
}

