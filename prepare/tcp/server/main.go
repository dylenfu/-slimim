package main

import (
	"fmt"
	"net"

	log "github.com/thinkboy/log4go"
)

const host = "0.0.0.0:9090"

func main() {
	address, err := net.ResolveTCPAddr("tcp4", host)
	if err != nil {
		log.Error(err)
		return
	}

	listener, err := net.ListenTCP("tcp4", address)
	if err != nil {
		log.Error(err)
	}

	acceptTCP(listener)
}

func acceptTCP(listener *net.TCPListener) {
	defer listener.Close()
	for {
		conn, err := listener.AcceptTCP()
		if err != nil {
			log.Error(err)
			continue
		}
		handleConnection(conn)
	}
}

func handleConnection(conn *net.TCPConn) {
	defer conn.Close()
	recvBuf := make([]byte, 128)
	sendBuf := []byte("server get client:" + conn.RemoteAddr().String() + " message!")
	for {
		_, err1 := conn.Read(recvBuf)
		if err1 != nil {
			log.Error(err1)
			return
		}
		_, err2 := conn.Write(sendBuf)
		if err2 != nil {
			log.Error(err2)
			return
		}
		fmt.Println("server recv:::::" + string(recvBuf))
	}
}
