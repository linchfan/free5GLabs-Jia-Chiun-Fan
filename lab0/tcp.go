package lab0

import (
	"fmt"
	"net"
	"strconv"
)

type listenerInterface func(string, int, func(net.Conn) error)
type handlerInterface func(net.Conn) error

func TCPListener(IP string, port int, handler func(net.Conn) error) {
	listener, err := net.Listen("tcp", net.JoinHostPort(IP, strconv.Itoa(port))) // Listen on the specified IP and port

	if err != nil {
		fmt.Println("Error listening", err.Error())
		return // Terminate
	}

	fmt.Println("TCP is listening on ", net.JoinHostPort(IP, strconv.Itoa(port)))

	for {
		conn, err := listener.Accept() // Accept incoming connections
		if err != nil {
			fmt.Println("Error accepting", err.Error())
			return
		}

		fmt.Println("new client accepted: ", net.JoinHostPort(IP, strconv.Itoa(port)))

		go func() {
			err := handler(conn) // Handle the connection using the provided handler function
			if err != nil {
				fmt.Println("Error handling connection", err.Error())
				return
			}
		}()
	}
}

func TCPHandler(conn net.Conn) error {
	defer conn.Close()

	remoteAddr := conn.RemoteAddr() // Client's remote address
	tcpAddr, ok := remoteAddr.(*net.TCPAddr)

	if !ok {
		fmt.Println("Error getting remote address")
		err := fmt.Errorf("failed to get TCP remote address")
		return err
	}

	clientAddr := tcpAddr.IP.String() // Client's remote IP address
	clientPort := tcpAddr.Port        // Client's remote port
	fmt.Println("Handle Request from [", clientAddr, ":", clientPort, "]")

	buf := make([]byte, 1024) // Prepare a buffer to read data from the connection

	for {
		n, err := conn.Read(buf) // Read data from the connection
		if err != nil {          // EOF or other error
			fmt.Println("Error reading", err.Error())
			return err
		}

		if n > 0 { // Correctly read data from the connection
			_, err = conn.Write(buf[:n]) // Echo back the received data to the client
			if err != nil {
				fmt.Println("Error writing", err.Error())
				return err
			}
		} else if n == 0 { // EOF
			fmt.Println("Client [", net.JoinHostPort(clientAddr, strconv.Itoa(clientPort)), "Error: EOF")
			return nil
		} else { // Error reading from the connection
			fmt.Println("Error reading")
			return nil
		}
	}

	return nil
}
