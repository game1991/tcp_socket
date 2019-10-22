package main

import (
	"fmt"
	"log"
	"net"
	"os"
)

func main(){
	//主动连接服务器
	conn,err:=net.Dial("tcp","127.0.0.1:8000")
	if err!=nil{
		log.Fatal(err)
		return
	}
	defer conn.Close()
	//新建协程
	go func() {
		//切片缓存
		str:=make([]byte,2048)
		//从键盘输入内容，发送到服务器
		for{
			n,err:=os.Stdin.Read(str)//从键盘读取内容，放在str
			if err!=nil{
				log.Fatal("键盘读取错误:",err)
				return
			}
			conn.Write(str[:n])
		}
	}()

	//接收服务器回复的数据
	buf:=make([]byte,2048)
	for ; ;  {
		n,err:=conn.Read(buf)//接收服务器的请求
		if err!=nil{
			log.Fatal("conn read err=",err)
			return
		}
		fmt.Println(string(buf[:n]))
	}
}
