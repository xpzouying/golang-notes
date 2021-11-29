package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func main() {

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(3 * time.Second)

		fmt.Fprintf(w, "hello world")
	})

	log.Fatal(http.ListenAndServe(":8080", nil))
}
