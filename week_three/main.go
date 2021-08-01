package main

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"golang.org/x/sync/errgroup"
)

type Server struct {
	ctx    context.Context
	cancel context.CancelFunc
	s      http.Server
}

func helloHandler() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "This is the hello handler.")
		time.Sleep(1 * time.Second)
		fmt.Fprintln(w, "hello handler end")
	})
}

func New(addr string, handler http.Handler) *Server {
	var s Server
	s.s = http.Server{
		Addr:    addr,
		Handler: handler,
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
				s.cancel()
			}
		}
	})

	if err := eg.Wait(); err != nil && !errors.Is(err, context.Canceled) {
		return err
	}
	return nil
}

func main() {
	fmt.Println("week three")
	s := New("0.0.0.0:8080", helloHandler())
	err := s.Run()
	if err != nil {
		log.Println(err)
	}
}
