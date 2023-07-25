package server

import (
	"net"
	"os"

	"github.com/levinion/grass/app"
	"github.com/levinion/grass/message"
)

func Serve(applist *app.AppList) {
	os.Remove(message.SockAddr)
	listener, err := net.Listen("unix", message.SockAddr)
	if err != nil {
		panic(err)
	}
	for {
		conn, err := listener.Accept()
		if err != nil {
			panic(err)
		}
		msg := readMsg(conn)
		handleMsg(applist, msg, conn)
	}
}

func readMsg(conn net.Conn) *message.Msg {
	buf := make([]byte, 1024)
	n, err := conn.Read(buf)
	if err != nil {
		panic(err)
	}
	data := buf[:n]
	return message.UnMarshal(data)
}

func handleMsg(applist *app.AppList, msg *message.Msg, conn net.Conn) {
	var feedback string
	switch msg.Command {
	case message.Add:
		feedback = applist.AddThenStart(msg.Args)
	case message.Remove:
		feedback = applist.Remove(msg.Args)
	case message.Stop:
		feedback = applist.Stop(msg.Args)
	case message.Start:
		feedback = applist.Start(msg.Args)
	case message.Show:
		feedback = applist.Show()
	case message.Reload:
		feedback = applist.Reload(msg.Args)
	}
	sendFeedback(conn, feedback)
}

func sendFeedback(conn net.Conn, message string) {
	conn.Write([]byte(message))
}
