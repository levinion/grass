package client

import (
	"fmt"
	"net"
	"os"

	"github.com/levinion/grass/message"
)

func SendMsg(command int, args string) {
	conn, err := net.Dial("unix", message.SockAddr)
	if err != nil {
		fmt.Println("No server running...Quit now...")
		os.Exit(0)
	}
	conn.Write(message.BuildMsg(command, args).Marshal())
	info := recvFeedback(conn)
	fmt.Println(info)
	conn.Close()
}

func recvFeedback(conn net.Conn) string {
	buf := make([]byte, 1024)
	n, err := conn.Read(buf)
	if err != nil {
		panic(err)
	}
	data := buf[:n]
	return string(data)
}
