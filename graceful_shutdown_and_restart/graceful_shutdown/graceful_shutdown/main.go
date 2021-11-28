package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
)

func main() {
	// 定义需要捕获的信号，并且通过channel进行传递。
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt) // os.Interrupt = syscall.SIGINT

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(3 * time.Second)
		fmt.Fprintf(w, "hello world")
	})
	server := &http.Server{
		Addr:    ":8080",
		Handler: nil,
	}
	go func() {
		log.Println(server.ListenAndServe())
	}()

	// 等待接收信号
	s := <-c
	log.Printf("receive signal: %v", s)

	log.Println("http server shutdown: ", server.Shutdown(context.Background()))
}
