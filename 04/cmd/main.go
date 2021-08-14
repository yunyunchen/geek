package main

import (
	"context"
	v1 "geek/04/api/v1"
	"golang.org/x/sync/errgroup"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

func main() {

	//InitializeAllInstance()
	ctx, cancel := context.WithCancel(context.Background())
	g, errCtx := errgroup.WithContext(ctx)

	//创建一个http server
	/*mux := http.NewServeMux()
	mux.HandleFunc("/hello", helloServer)*/

	rh := v1.InitRouter()
	srv := &http.Server{
		Addr:    "0.0.0.0:8000",
		Handler: rh,
	}

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
