package main

import (
	"context"
	"fmt"
	"golang.org/x/sync/errgroup"
	"io"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

func helloServer(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "hello, world!\n")
	fmt.Println("hello ")
}

func serveStart(srv *http.Server) error {
	http.HandleFunc("/hello", helloServer)
	return srv.ListenAndServe()
}

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	g, errCtx := errgroup.WithContext(ctx)

	//创建一个http server
	srv := &http.Server{Addr: "0.0.0.0:8000"}

	g.Go(func() error {
		log.Println("http server start ok")
		return srv.ListenAndServe()
	})

	g.Go(func() error {
		<-errCtx.Done()
		log.Println("http server stop")
		return srv.Shutdown(ctx)
	})

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	g.Go(func() error {
		for {
			select {
			case <-errCtx.Done():
				log.Println("errgroup exit ...")
				return errCtx.Err()
			case <-quit:
				cancel()
				log.Println("signal stop quit")
			}
		}
		return nil
	})

	if err := g.Wait(); err != nil {
		log.Println("group error: ", err)
	}
}
