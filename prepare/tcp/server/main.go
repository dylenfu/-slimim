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
		log.Error("net.ResolveTCPAddr(\"tcp4\",\"%s\") error(%v)", host, err)
		return
	}

	listener, err := net.ListenTCP("tcp4", address)
	if err != nil {
		log.Error("net.ListenTCP(\"tcp4\", \"%s\") error(%v)", host, err)
	}

	acceptTCP(listener)
}

func acceptTCP(listener *net.TCPListener) {
	for {
		conn, err := listener.AcceptTCP()
		if err != nil {
			log.Error("listener.AcceptTCP(\"%s\") error(%v)", listener.Addr().String(), err)
			continue
		}
		handleConnection(conn)
	}
	defer listener.Close()
}

func handleConnection(conn *net.TCPConn) {
	recvBuf := make([]byte, 128)
	sendBuf := []byte("server get client:" + conn.RemoteAddr().String() + " message!")
	for {
		_, err1 := conn.Read(recvBuf)
		if err1 != nil {
			log.Error("conn.Read(\"%s\") error(%v)", string(sendBuf), err1)
			return
		}
		_, err2 := conn.Write(sendBuf)
		if err2 != nil {
			log.Error("conn.Write(\"%s\") error(%v)", string(sendBuf), err2)
			return
		}
		fmt.Println("server recv:::::" + string(recvBuf))
	}
	defer conn.Close()
}
