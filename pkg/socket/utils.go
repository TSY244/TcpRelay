package socket

import "net"

func Write(conn net.Conn, data []byte) (int, error) {
	return conn.Write(data)
}

func Read(conn net.Conn, data []byte) (int, error) {
	return conn.Read(data)
}
