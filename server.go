package main

import (
	"bufio"
	"fmt"
	"io"
	"net"
	"os"
	"strings"
)

var(
	w io.Writer
	// r io.Reader
)

func main() {
	fmt.Println("Logs from your program will appear here!")
	l, err := net.Listen("tcp", "0.0.0.0:4221")
	if err != nil {
		fmt.Println("Failed to bind to port 4221")
		os.Exit(1)
	}
	
	conn, err := l.Accept()
	if err != nil {
		fmt.Println("Error accepting connection: ", err.Error())
		os.Exit(1)
	}
	buf := make([]byte, 256)
		_,err =conn.Read(buf)
	if err!=nil{
		fmt.Print("Error while reading the data")
		os.Exit(1)
	}
	r :=strings.NewReader(string(buf))
	bufReader := bufio.NewReader(r)

	str,err := bufReader.ReadString('\n')
	if err!=nil{
		fmt.Print("Error while reading the data")
		os.Exit(1)
	}
	lines := strings.Fields(str)
	path := lines[1]
	res := "HTTP/1.1 404 Not Found\r\n\r\n"
	if(path=="/"){
		res = "HTTP/1.1 200 OK\r\n\r\n"
	}

	_,err = fmt.Fprintf(w,"%s",res)
	// _,err = conn.Write([]byte(res))
	if err!=nil{
		fmt.Print("Error Writing res header")
		os.Exit(1)
	}
	conn.Close()
}
