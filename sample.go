package main

import (
	"fmt"
	"net"
)

// Create a buffer to read data from the listening port
var Buffer = make([]byte, 1024)

// Create a packet forwarder
func fwdPackets(data []byte, destination string) {

	for {
		fwdData, err := net.Dial("udp", destination)
		if err != nil {
			fmt.Printf("Error communicating with destination %s", destination)
		}
		_, err = fwdData.Write(data)
		if err != nil {
			fmt.Printf("Error forwarding packets to destination")
		}
		continue
	}
}

func readToBuffer(conn net.Conn, destination string) {
	defer conn.Close()

	for {
		// Read data from the connection
		n, err := conn.Read(Buffer)
		if err != nil {
			fmt.Println("Error reading:", err)
			continue
		}
		go fwdPackets(Buffer[:n], destination)
	}
}

func listenOnPort(listen net.Listener) net.Conn {
	for {
		// Accept incoming connection
		conn, err := listen.Accept()
		if err != nil {
			fmt.Println("Error accepting connection:", err)
			continue
		}
		return conn
	}
}

func Replicate(src_port, destination string) {

	// Start listening for incoming connections
	listener, err := net.Listen("udp", fmt.Sprintf(":%s", src_port))
	if err != nil {
		fmt.Println("Error listening:", err)
	}
	defer listener.Close()

	fmt.Printf("Server listening on port %s ", src_port)

	conn := listenOnPort(listener)

	go readToBuffer(conn, destination)

}
