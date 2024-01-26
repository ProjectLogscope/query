package server

import (
	"context"
	"errors"
	"fmt"
	"net"
	"net/http"
	"net/netip"

	"github.com/hardeepnarang10/query/service/router"
)

type server struct {
	router router.Router
}

func New(r router.Router) (Server, error) {
	return &server{
		router: r,
	}, nil
}

func (s *server) Start(port uint16) error {
	if s == nil {
		return errors.New("server: start attempted on uninitialized server instance")
	}

	err := s.router.Listen(net.TCPAddrFromAddrPort(netip.AddrPortFrom(netip.IPv4Unspecified(), port)).String())
	if err != nil && !errors.Is(err, http.ErrServerClosed) {
		return fmt.Errorf("unable to listen and serve without tls: %w", err)
	}
	return nil
}

func (s *server) Shutdown(ctx context.Context) error {
	if s == nil {
		return errors.New("server: shutdown attempted on uninitialized server instance")
	}

	if err := s.router.Shutdown(ctx); err != nil {
		return fmt.Errorf("unable to gracefully shutdown server: %w", err)
	}
	return nil
}
