package rp

import (
	"github.com/x6r/rp/ipc"
	"github.com/x6r/rp/rpc"
)

func NewClient(ID string) (*rpc.Client, error) {
	c := &rpc.Client{ClientID: ID}

	i, err := ipc.NewIPC()
	if err != nil {
		return nil, err
	}

	c.IPC = i

	if err := c.Login(); err != nil {
		return nil, err
	}

	return c, nil
}
