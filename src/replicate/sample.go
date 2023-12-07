package sample

import (
	"fmt"
	"net"
	"os"
	"os/signal"
	"sync"
	"syscall"
)

const bufferSize = 1024

func fwdPackets(data []byte, destination string) {
	fwdData, err := net.Dial("udp", destination)
	if err != nil {
		fmt.Printf("Error communicating with destination %s: %s\n", destination, err)
		return
	}
	defer fwdData.Close()

	_, err = fwdData.Write(data)
	if err != nil {
		fmt.Printf("Error forwarding packets to destination: %s\n", err)
	}

}

func handleConnection(conn *net.UDPConn, destination string, wg *sync.WaitGroup) {
	defer conn.Close()
	defer wg.Done()

	buffer := make([]byte, bufferSize)

	for {
		n, err := conn.Read(buffer)
		if err != nil {
			fmt.Println("Error reading:", err)
			break
		}

		go fwdPackets(buffer[:n], destination)
	}
}

func Replicate(srcPort, destination string) {
	udpAddr, err := net.ResolveUDPAddr("udp", fmt.Sprintf("localhost:%s", srcPort))
	if err != nil {
		fmt.Println("Error resolving UDP address:", err)
		return
	}

	listener, err := net.ListenUDP("udp", udpAddr)
	if err != nil {
		fmt.Println("Error listening:", err)
		return
	}
	defer listener.Close()

	fmt.Printf("Server listening on port %s\n", srcPort)

	var wg sync.WaitGroup
	wg.Add(1)
	go handleConnection(listener, destination, &wg)

	// Wait for a termination signal to gracefully shutdown
	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, os.Interrupt, syscall.SIGTERM)
	<-sigCh

	// Signal received, perform cleanup and exit gracefully
	fmt.Println("Shutting down...")
	listener.Close()
	wg.Wait()
	fmt.Println("Server stopped.")
}

func Test() {
	srcPort := "2055"                 // Set your source port
	destination := "10.42.15.30:2055" // Set your destination address

	Replicate(srcPort, destination)
}
