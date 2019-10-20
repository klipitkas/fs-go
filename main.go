package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
)

var port int
var dir string

func main() {

	// Handle the flags that are provided.
	flag.IntVar(&port, "port", 8080, "The port that the server will listen to.")
	flag.StringVar(&dir, "dir", ".", "The root directory that will be served.")
	// Parse the flags.
	flag.Parse()

	// Create the fileserver.
	fs := http.FileServer(http.Dir(dir))

	// Print details to the console.
	log.Printf("FS - Port: %v | Dir: %v", port, dir)

	// Start the file server.
	if err := http.ListenAndServe(":"+fmt.Sprintf("%d", port), fs); err != nil {
		log.Fatalf("server listen on port %v: %v", port, err)
	}

}
