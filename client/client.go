package client

import (
	"fmt"
	"net"
	"os"
	"time"

	"github.com/levinion/grass/message"
)

func SendMsg(command int, args string) {
	conn := retry(3)
	conn.Write(message.BuildMsg(command, args).Marshal())
	info := recvFeedback(conn)
	fmt.Println(info)
	conn.Close()
}

func retry(count int) net.Conn {
	if count == 0 {
		fmt.Println("No server running...Quit now...")
		os.Exit(0)
	}
	conn, err := net.Dial("unix", message.SockAddr)
	if err != nil {
		// try every 0.1s
		time.Sleep(time.Millisecond * 100)
		return retry(count - 1)
	}
	return conn
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
