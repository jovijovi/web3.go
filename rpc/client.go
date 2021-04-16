package rpc

import (
	"github.com/jovijovi/web3.go/rpc/transport"
)

type Client struct {
	transport transport.Transport
}

func NewClient(addr string) (*Client, error) {
	c := &Client{}

	t, err := transport.NewTransport(addr)
	if err != nil {
		return nil, err
	}
	c.transport = t
	return c, nil
}

func (c *Client) Close() error {
	return c.transport.Close()
}

func (c *Client) Call(method string, out interface{}, params ...interface{}) error {
	return c.transport.Call(method, out, params...)
}
