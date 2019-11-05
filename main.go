package main

import (
	"flag"
	"log"
	"net/http"
)

var (
	addr = flag.String("addr", ":8080", "The address that the server will listen to.")
	dir  = flag.String("dir", ".", "The root directory that will be served.")
)

func main() {
	// Parse the flags.
	flag.Parse()

	// Create the fileserver.
	fs := http.FileServer(http.Dir(*dir))

	// Print details to the console.
	log.Printf("fs-go - address: %s | directory: %s", *addr, *dir)

	// Start the file server.
	log.Fatal(http.ListenAndServe(*addr, fs))
}
