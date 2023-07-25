package message

import (
	"bytes"
	"encoding/gob"
	"os"
	"path/filepath"
)

const (
	Add = iota
	Remove
	Stop
	Start
	Show
)

var SockAddr = filepath.Join(os.TempDir(), "github.com/levinion/grass.sock")

type Msg struct {
	Command int
	Args    string
}

func BuildMsg(command int, args string) *Msg {
	return &Msg{
		Command: command,
		Args:    args,
	}
}

func (m *Msg) Marshal() []byte {
	buf := new(bytes.Buffer)
	gob.NewEncoder(buf).Encode(m)
	return buf.Bytes()
}

func UnMarshal(data []byte) *Msg {
	buffer := bytes.NewBuffer(data)
	msg := new(Msg)
	gob.NewDecoder(buffer).Decode(msg)
	return msg
}
