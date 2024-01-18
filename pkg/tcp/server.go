package tcp

import "net"

type TCPHandler func(conn net.Conn)

type TCPServer struct {
	Handler TCPHandler
}

func NewTCPServer(handler TCPHandler) *TCPServer {
	return &TCPServer{
		Handler: handler,
	}
}

func (s *TCPServer) ListenAndServe(addr string) error {
	l, err := net.Listen("tcp", addr)
	if err != nil {
		return err
	}
	defer l.Close()
	for {
		conn, err := l.Accept()
		if err != nil {
			return err
		}
		go s.Handler(conn)
	}
}
