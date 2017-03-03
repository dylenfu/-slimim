package main

import (
	"net"
	log "github.com/thinkboy/log4go"
	"fmt"
)

const host = "0.0.0.0:9090"

func main() {
	address, err := net.ResolveTCPAddr("tcp4", host)
	if err != nil {
		log.Error("net.ResolveTCPAddr(\"tcp4\", \"%s\") error(%v)", address.String(), err)
	}
	conn, err := net.DialTCP("tcp4", nil, address)
	if err != nil {
		log.Error("net.DailTCP(\"tcp4\", \"%s\") error(%v)", address.String(), err)
	}
	handleConnection(conn)
}

func handleConnection(conn *net.TCPConn) {
	sendBuf := []byte(conn.LocalAddr().String() + " get server message!")
	recvBuf := make([]byte, 128)
	for {
		_, err1 := conn.Write(sendBuf)
		if err1 != nil {
			log.Error("conn.Write(\"%s\") error(%v)", string(sendBuf), err1)
			return
		}
		_, err2 := conn.Read(recvBuf)
		if err2 != nil {
			log.Error("conn.Read(\"%s\") error(%v)", string(sendBuf), err2)
			return
		}
		fmt.Println("client recv:::::" + string(recvBuf))
	}
	defer conn.Close()
}