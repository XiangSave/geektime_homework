package server

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"golang.org/x/sync/errgroup"
)

type Server struct {
	ctx    context.Context
	cancel context.CancelFunc
	s      http.Server
}

func New(addr string, mux *http.ServeMux) *Server {
	var s Server
	s.s = http.Server{
		Addr:    addr,
		Handler: mux,
	}

	return &s
}

func (s *Server) Run() error {
	s.ctx, s.cancel = context.WithCancel(context.Background())
	eg, ctx := errgroup.WithContext(s.ctx)
	eg.Go(func() error {
		defer fmt.Println("Listen defer")
		return s.s.ListenAndServe()
	})

	eg.Go(func() error {
		<-ctx.Done()
		return s.s.Shutdown(s.ctx)
	})

	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT)

	eg.Go(func() error {
		for {
			select {
			case <-ctx.Done():
				return ctx.Err()
			case <-c:
				log.Println("run s.Stop()")
				// s.Stop()
				s.s.Shutdown(s.ctx)
			}
		}
	})

	if err := eg.Wait(); err != nil && !errors.Is(err, context.Canceled) {
		// if err := eg.Wait(); err != nil {
		return err
	}
	return nil
}

func (s *Server) Stop() error {
	if s.cancel != nil {
		s.cancel()
	}
	return nil
}
