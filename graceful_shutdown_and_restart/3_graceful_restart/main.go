package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	reuseport "github.com/libp2p/go-reuseport"
)

func main() {
	// 定义需要捕获的信号，并且通过channel进行传递。
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt) // os.Interrupt = syscall.SIGINT

	var server *http.Server
	{
		http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
			pid := os.Getpid()
			log.Printf("receive request: PID=%d", pid)

			time.Sleep(3 * time.Second)
			fmt.Fprintf(w, "hello world, from pid=%d", pid)
		})

		lis, err := reuseport.Listen("tcp", ":8080")
		if err != nil {
			panic(err)
		}

		server = &http.Server{
			Handler: nil,
		}
		go func() {
			if err := server.Serve(lis); err != nil && err != http.ErrServerClosed {
				log.Fatalf("http server exit: %v", err)
			}
		}()
	}

	// 等待接收信号
	s := <-c
	log.Printf("receive signal: %v", s)

	log.Println("http server shutdown: ", server.Shutdown(context.Background()))
}
