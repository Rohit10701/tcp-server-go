// // Listen for connections: 
// 		You need to listen on a particular TCP port for incoming client connections.
// // Handle client requests: 
// 		Once a connection is established, the server should handle requests from clients (e.g., reading data).
// // Respond back: 
// 		After receiving data, the server can send a response back to the client.
// // Close the connection: 
// 		Once done, the connection is closed.
package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
)

func handleConnection(conn net.Conn) {
	defer conn.Close()

	fmt.Printf("client connected: %s\n", conn.RemoteAddr().String())

	reader := bufio.NewReader(conn)

	for {
		message, err := reader.ReadString('\n')
		if err != nil {
			fmt.Printf("client disconnected: %s\n", conn.RemoteAddr().String())
			break
		}

		fmt.Printf("recived: %s", message)

		response := fmt.Sprintf("Server received: %s", message)
		_, err = conn.Write([]byte(response))
		if err != nil {
			fmt.Println("error writing response:", err)
			break
		}
	}
}

func main() {
	listener, err := net.Listen("tcp", ":3000")
	if err != nil {
		log.Fatal("error starting server:", err)
	}
	defer listener.Close()

	fmt.Println("tcp server is listening on port 3000...")

	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("error accepting connection:", err)
			continue
		}

		go handleConnection(conn)
	}
}
