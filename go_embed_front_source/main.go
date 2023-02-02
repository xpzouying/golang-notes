package main

import (
	"embed"
	"io/fs"
	"log"
	"net/http"
)

//go:embed all:dist
var distFS embed.FS

func main() {

	distFS, err := fs.Sub(distFS, "dist")
	panicError(err)

	http.Handle("/", http.FileServer(http.FS(distFS)))

	log.Printf("starting HTTP server at :9090")
	log.Fatal(http.ListenAndServe(":9090", nil))
}

func panicError(err error) {
	if err != nil {
		panic(err)
	}
}
