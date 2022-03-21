package ipc

import (
	"bytes"
	"encoding/binary"
	"net"
	"os"
)

type IPC struct {
	Socket net.Conn
}

func NewIPC() (*IPC, error) {
	ipc := IPC{}

	if err := ipc.OpenSocket(); err != nil {
		return nil, err
	}

	return &ipc, nil
}

// Choose the right directory to the ipc socket and return it
func GetIpcPath() string {
	variablesnames := []string{"XDG_RUNTIME_DIR", "TMPDIR", "TMP", "TEMP"}

	for _, variablename := range variablesnames {
		path, exists := os.LookupEnv(variablename)

		if exists {
			return path
		}
	}

	return "/tmp"
}

func (ipc *IPC) CloseSocket() error {
	if ipc.Socket != nil {
		ipc.Socket.Close()
		ipc.Socket = nil
	}
	return nil
}

// Read the socket response
func (ipc *IPC) Read() string {
	buf := make([]byte, 512)
	payloadlength, _ := ipc.Socket.Read(buf)

	return string(buf[8:payloadlength])
}

// Send opcode and payload to the unix socket
func (ipc *IPC) Send(opcode int, payload string) (string, error) {
	buf := new(bytes.Buffer)

	err := binary.Write(buf, binary.LittleEndian, int32(opcode))
	if err != nil {
		return "", err
	}

	err = binary.Write(buf, binary.LittleEndian, int32(len(payload)))
	if err != nil {
		return "", err
	}

	buf.Write([]byte(payload))
	_, err = ipc.Socket.Write(buf.Bytes())
	if err != nil {
		return "", err
	}

	return ipc.Read(), nil
}
