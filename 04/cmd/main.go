package main

import (
	"context"
	"flag"
	v1 "geek/04/api/v1"
	"geek/04/configs"
	"geek/04/internal/server"
	"geek/04/internal/service"
	"golang.org/x/sync/errgroup"
	"log"
	"os"
	"os/signal"
	"syscall"
)

type userApp struct {
	httpSrv     *server.HttpServer
	userService *service.UserService
}

func newUserApp(httpSrv *server.HttpServer,userService *service.UserService) *userApp {
	return &userApp{
		httpSrv:     httpSrv,
		userService: userService,
	}
}
func loadConfig() *configs.Config {
	var configFile string
	flag.StringVar(&configFile, "c", "config path", "full path of config file")
	flag.Parse()

	config := configs.NewConfig(configFile)
	log.Printf("config: %+v\n", config)

	return config
}

func main() {

	config := loadConfig()

	ctx, cancel := context.WithCancel(context.Background())
	g, errCtx := errgroup.WithContext(ctx)

	userApp, err := initApp(config)
	if err != nil {
		log.Println("http server start error: %v\n",err)
	}
	//创建一个http server
	/*mux := http.NewServeMux()
	mux.HandleFunc("/hello", helloServer)*/

	httpSrv := userApp.httpSrv
	httpSrv.AddRouters(v1.InitRouter)

	//rh := v1.InitRouter()
	//srv := &http.Server{
	//	Addr:    "0.0.0.0:8000",
	//	Handler: rh,
	//}

	/*g.Go(func() error {
		log.Println("http server start ok")
		return srv.ListenAndServe()
	})*/

	g.Go(func() error {
		<-errCtx.Done()
		log.Println("http server stop")
		return httpSrv.Shutdown(ctx)
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
