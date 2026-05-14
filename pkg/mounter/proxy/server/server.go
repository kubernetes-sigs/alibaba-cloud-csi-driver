package server

import (
	"errors"
	"net"
	"sync"
	"sync/atomic"
	"time"

	"k8s.io/klog/v2"
)

// Server accepts connections on a Unix socket and dispatches each to Handle.
type Server struct {
	listener *net.UnixListener
	timeout  time.Duration
	connSeq  atomic.Int64

	wg   sync.WaitGroup
	done chan struct{}
}

// NewServer creates a Server that will accept connections on the given listener.
func NewServer(listener *net.UnixListener, timeout time.Duration) *Server {
	return &Server{
		listener: listener,
		timeout:  timeout,
		done:     make(chan struct{}),
	}
}

// Serve accepts connections in a loop. Each connection is handled in a
// separate goroutine. Serve blocks until the listener is closed.
func (s *Server) Serve() {
	defer close(s.done)
	for {
		conn, err := s.listener.AcceptUnix()
		if err != nil {
			if errors.Is(err, net.ErrClosed) {
				break
			}
			klog.ErrorS(err, "Failed to accept connection")
			continue
		}
		s.wg.Go(func() {
			s.handle(conn)
		})
	}
}

func (s *Server) handle(conn *net.UnixConn) {
	seq := s.connSeq.Add(1)
	defer func() {
		if err := conn.Close(); err != nil {
			klog.ErrorS(err, "Failed to close", "seq", seq)
		}
	}()
	if err := Handle(conn, s.timeout, seq); err != nil {
		klog.ErrorS(err, "Failed to handle", "seq", seq)
	}
}

// Close stops the server by closing the listener. It blocks until Serve has
// exited and all in-flight connections have been handled.
func (s *Server) Close() error {
	err := s.listener.Close()
	<-s.done
	s.wg.Wait()
	return err
}
