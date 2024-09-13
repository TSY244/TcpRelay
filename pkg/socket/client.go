package socket

import "net"

type Client struct {
	Conn *net.Conn
}

func NewDefaultClient() (*Client, error) {
	conn, err := net.Dial("tcp", "192.168.79.137:12345")
	if err != nil {
		return nil, err
	}
	client := &Client{
		Conn: &conn,
	}
	return client, nil
}

func NewClient(ip, port string) (*Client, error) {
	conn, err := net.Dial("tcp", ip+":"+port)
	if err != nil {
		return nil, err
	}
	client := &Client{
		Conn: &conn,
	}
	return client, nil
}

func (c *Client) Close() error {
	return (*c.Conn).Close()
}
