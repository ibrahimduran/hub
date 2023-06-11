package hub

import (
	"context"
	"errors"
	"fmt"
	"net"
	"net/netip"

	"github.com/ibrahimduran/hub/internal/protocol"
)

var (
	ErrAcceptFailed = errors.New("client accept failed")
)

type Server interface {
	Listen(port uint16)
	Use(provider protocol.Provider)
}

type server struct {
	sock      *net.TCPListener
	providers []protocol.Provider
}

func NewServer() Server {
	s := server{}

	return &s
}

func (s *server) Use(provider protocol.Provider) {
	s.providers = append(s.providers, provider)
}

func (s *server) Listen(port uint16) {
	addr := netip.AddrPortFrom(netip.IPv4Unspecified(), port)
	tcpAddr := net.TCPAddrFromAddrPort(addr)
	sock, err := net.ListenTCP("tcp", tcpAddr)

	if err != nil {
		panic(err)
	}

	s.sock = sock

	for {
		err := s.accept()

		if err != nil {
			fmt.Println(errors.Join(ErrAcceptFailed, err))
		}
	}
}

func (s *server) accept() error {
	conn, err := s.sock.AcceptTCP()

	if err != nil {
		return err
	}

	proto := protocol.Protocol{
		Providers:         s.providers,
		Stream:            conn,
		BufferSize:        200,
		BufferBacklogSize: 1000,
	}

	go proto.Run(context.Background())

	return nil
}
