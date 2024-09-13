package socket

import "net"

type Server struct {
	listener *net.Listener
}

func NewDefaultServer() (*Server, error) {
	// default list on 10090
	l, err := net.Listen("tcp", ":12345")
	if err != nil {
		return nil, err
	}
	server := &Server{
		listener: &l,
	}
	return server, nil
}

func NewServer(host string) (*Server, error) {
	l, err := net.Listen("tcp", host)
	if err != nil {
		return nil, err
	}
	server := &Server{
		listener: &l,
	}
	return server, nil
}

func (s *Server) Accept() (net.Conn, error) {
	return (*s.listener).Accept()
}

func (s *Server) Close() error {
	return (*s.listener).Close()
}
