package main

import (
	"fmt"
	"net"
)

func TcpListener(IP string, port int, handler func(net.Conn) error) {
	listener, err := net.Listen("tcp", IP+":"+string(port))

	if err != nil {
		fmt.Println("Error listening", err.Error())
		return // terminate
	}

	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Error accepting", err.Error())
			return
		}

		err = handler(conn)
		if err != nil {
			fmt.Println("Error handling connection", err.Error())
		}
	}
}

func TcpHandler() error {

}
