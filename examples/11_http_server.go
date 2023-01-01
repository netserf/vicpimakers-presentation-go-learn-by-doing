package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

const version = "0.1.0"
const defaultPort = "8080"

func main() {

	// You can create a basic HTTP server using the net/http package.
	// net/http servers use handlers to register server behaviour.
	// A handler is an object implementing the http.Handler interface.
	// You can use the http.HandlerFunc adapter on functions with the
	// appropriate signature.
	mux := http.NewServeMux()
	mux.HandleFunc("/", hello)
	mux.HandleFunc("/hello", hello)
	mux.HandleFunc("/headers", headers)

	// use HTTP_PORT environment variable, or default to 8080
	port := os.Getenv("HTTP_PORT")
	if port == "" {
		port = defaultPort
	} else {
		log.Printf("Override server port to %s", port)
	}

	// ListenAndServe() activates the server with the given port and a handler.
	log.Printf("Server listening on localhost port %s", port)
	log.Fatal(http.ListenAndServe(":"+port, mux))
}

// The handler functions take http.ResponseWriter and http.Request arguments.
// The response writer is used to fill in the HTTP response.
func hello(w http.ResponseWriter, r *http.Request) {
	log.Printf("Serving request: %s", r.URL.Path)
	host, _ := os.Hostname()
	fmt.Fprintf(w, "Hello, world!\n")
	fmt.Fprintf(w, "Version: %s\n", version)
	fmt.Fprintf(w, "Hostname: %s\n", host)
}

func headers(w http.ResponseWriter, r *http.Request) {
	log.Printf("Serving request: %s", r.URL.Path)
	for name, headers := range r.Header {
		for _, h := range headers {
			fmt.Fprintf(w, "%v: %v\n", name, h)
		}
	}
}
