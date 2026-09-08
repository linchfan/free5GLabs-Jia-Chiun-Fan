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

	fmt.Println("TCP is listening on ", IP, ":", string(port))

	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Error accepting", err.Error())
			return
		}

		fmt.Println("new client accepted: ", IP, ":", string(port))

		go func() {
			err := handler(conn)
			if err != nil {
				fmt.Println("Error handling connection", err.Error())
				return
			}
		}()
	}
}

func TcpHandler(conn net.Conn) error {
	remoteAddr := conn.RemoteAddr() // Client's remote address
	tcpAddr, ok := remoteAddr.(*net.TCPAddr)

	if !ok {
		fmt.Println("Error getting remote address", err.Error())
		err := fmt.Errorf("failed to get TCP remote address")
		return err
	}

	clientAddr := tcpAddr.IP.String() // Client's remote IP address
	clientPort := tcpAddr.Port        // Client's remote port
	fmt.Println("Handle Request from [", clientAddr, ":", clientPort, "]")
}
