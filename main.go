package main

import (
	"fmt"
	"net"
	"net/http"
)

const (
	SERVER_HOST = "localhost"
	SERVER_PORT = "9988"
	SERVER_TYPE = "tcp"
)

func headers(w http.ResponseWriter, req *http.Request) {

	for name, headers := range req.Header {
		for _, h := range headers {
			fmt.Fprintf(w, "%v: %v\n", name, h)
		}
	}
}

func main() {

	http.HandleFunc("/headers", headers)

	http.ListenAndServe(SERVER_PORT, nil)
}

// func main() {

// 	fmt.Println("Server Running...")
// 	server, err := net.Listen(SERVER_TYPE, SERVER_HOST+":"+SERVER_PORT)
// 	if err != nil {
// 		fmt.Println("Error listening:", err.Error())
// 		os.Exit(1)
// 	}
// 	defer server.Close()

// 	fmt.Println("Listening on " + SERVER_HOST + ":" + SERVER_PORT)
// 	fmt.Println("Waiting for client...")

// 	for {
// 		// Wait for a connection
// 		connection, err := server.Accept()
// 		if err != nil {
// 			fmt.Println("Error accepting: ", err.Error())
// 			os.Exit(1)
// 		}
// 		fmt.Println("client connected")

// 		// Handle the connection in a new goroutine
// 		go handleConnection(connection)
// 	}
// }

func handleConnection(connection net.Conn) {

	// Make a buffer to hold incoming data
	buffer := make([]byte, 1024)

	// Read the incoming connection into the buffer
	mLen, err := connection.Read(buffer)
	if err != nil {
		fmt.Println("Error reading:", err.Error())
	}
	fmt.Println("Received: ", string(buffer[:mLen]))

	// Send a response back to the client
	_, err = connection.Write([]byte("Thanks! Got your message:" + string(buffer[:mLen])))
	connection.Close()
}
