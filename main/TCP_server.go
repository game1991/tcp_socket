package main

import (
	"fmt"
	"log"
	"net"
	"strings"
)

func main(){
	//监听
	listener,err:=net.Listen("tcp","127.0.0.1:8000")
	if err!=nil{
		log.Fatal("无连接")
		return
	}
	defer listener.Close()
	//接收多个用户
	for {
		conn,err:=listener.Accept()
		if err!=nil{
			log.Fatal("接收连接失败")
			return
		}
		//处理用户请求，新建协程
		go handleConn(conn)
	}
}

func handleConn(conn net.Conn)  {
	//函数调用完毕，自动关闭conn
	defer conn.Close()
	//获取客户端的网络地址信息
	addr:=conn.RemoteAddr().String()
	fmt.Printf("%v connect sucessful!\n",addr)
	buf:=make([]byte,2048)
	//读取用户数据
	for ; ;  {
		n,err:=conn.Read(buf)
		if err!=nil{
			log.Fatal("conn read err=",err)
			return
		}
		//fmt.Println(len(string(buf[:n])))
		fmt.Printf("[%s]:%s",addr,string(buf[:n]))
		if "exit"== string(buf[:n-1]) {
			fmt.Println(addr,"exit connected")
			return
		}
		//把数据转换成大写，再发送给用户
		conn.Write([]byte(strings.ToUpper(string(buf[:n]))))
	}


}