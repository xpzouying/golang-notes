package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"golang.org/x/sync/errgroup"
)

func handleHello(w http.ResponseWriter, r *http.Request) {
	time.Sleep(3 * time.Second)
	fmt.Fprintf(w, "hello world")
}

func main() {
	var (
		port = ":8080"
	)

	done := make(chan struct{}, 1)
	g := new(errgroup.Group)

	g.Go(func() error {
		// 定义需要捕获的信号，并且通过channel进行传递。
		c := make(chan os.Signal, 1)
		signal.Notify(c, os.Interrupt) // os.Interrupt = syscall.SIGINT

		for {
			<-c

			log.Println("quit by signal")
			close(done)
			return nil
		}
	})

	g.Go(func() error {
		http.HandleFunc("/", handleHello)

		server := &http.Server{
			Addr:    port,
			Handler: nil,
		}

		log.Printf("start http server: %s", port)

		g.Go(func() error {
			<-done

			log.Println("will shutdown http server...")

			log.Println("http server shutdown: ", server.Shutdown(context.Background()))
			return nil
		})

		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("http server exit: %v", err)
		} else {
			log.Printf("http server exit succ: %v", err)
		}

		return nil
	})

	err := g.Wait()
	log.Printf("main quite: %v", err)
}
