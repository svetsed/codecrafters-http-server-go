package main

import (
	"fmt"
	"net"
	"os"
)

// Ensures gofmt doesn't remove the "net" and "os" imports above (feel free to remove this!)
var _ = net.Listen
var _ = os.Exit

func main() {
	// You can use print statements as follows for debugging, they'll be visible when running tests.
	fmt.Println("Logs from your program will appear here!")

	l, err := net.Listen("tcp", "0.0.0.0:4221")
	if err != nil {
		fmt.Println("Failed to bind to port 4221")
		os.Exit(1)
	}
	defer l.Close()

	conn, err := l.Accept()
	if err != nil {
		fmt.Println("Error accepting connection: ", err.Error())
		os.Exit(1)
	}

	b := make([]byte, 1024)
	if _, err := conn.Read(b); err != nil {
		fmt.Println("Error reading request: ", err.Error())
		os.Exit(1)
	}
	reqStr := string(b)
	strSlice := strings.Split(reqStr, "\r\n")
	reqLine := strings.Split(strSlice[0], " ")

	fmt.Println(reqLine[1])

	var resp string
	if reqLine[1] == "/" {
		resp = "HTTP/1.1 200 OK\r\n\r\n"
	} else {
		resp = "HTTP/1.1 404 Not Found\r\n\r\n"
	}

	_, err = conn.Write([]byte(resp))
	if err != nil {
		fmt.Println("Error writing response: ", err.Error())
		os.Exit(1)
	}

	fmt.Println("Accepted connection from: ", conn.RemoteAddr())
}
