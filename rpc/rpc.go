package rpc

import (
	"crypto/rand"
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/x6r/rp/ipc"
)

// Login sends a handshake in the socket and returns an error or nil
func (c *Client) Login() error {
	if !c.Logged {
		payload, err := json.Marshal(Handshake{"1", c.ClientID})
		if err != nil {
			return err
		}

		i, err := ipc.NewIPC()
		if err != nil {
			return err
		}
		c.IPC = i

		// TODO: Response should be parsed
		c.IPC.Send(0, string(payload))
	}
	c.Logged = true

	return nil
}

func (c *Client) Logout() {
	c.Logged = false

	err := c.IPC.CloseSocket()
	if err != nil {
		panic(err)
	}
}

func (c *Client) SetActivity(activity Activity) error {
	if !c.Logged {
		return nil
	}

	nonce, err := c.getNonce()
	if err != nil {
		log.Println(err)
	}

	payload, err := json.Marshal(Frame{
		"SET_ACTIVITY",
		Args{
			os.Getpid(),
			mapActivity(&activity),
		},
		nonce,
	})

	if err != nil {
		return err
	}

	if _, err := c.IPC.Send(1, string(payload)); err != nil {
		return err
	}
	return nil
}

func (c *Client) ResetActivity() error {
	if !c.Logged {
		return nil
	}

	nonce, err := c.getNonce()
	if err != nil {
		log.Println(err)
	}

	payload, err := json.Marshal(Frame{
		"SET_ACTIVITY",
		Args{
			os.Getpid(),
			mapActivity(&Activity{}),
		},
		nonce,
	})

	if _, err := c.IPC.Send(1, string(payload)); err != nil {
		return err
	}
	return nil
}

func (c *Client) getNonce() (string, error) {
	buf := make([]byte, 16)
	_, err := rand.Read(buf)
	if err != nil {
		return "", err
	}

	buf[6] = (buf[6] & 0x0f) | 0x40

	return fmt.Sprintf("%x-%x-%x-%x-%x", buf[0:4], buf[4:6], buf[6:8], buf[8:10], buf[10:]), nil
}
