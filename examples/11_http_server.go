package main

import (
	"fmt"
	"net/http"
)

const port = ":8090"

func hello(w http.ResponseWriter, req *http.Request) {
	fmt.Println("Received request on /hello")
	fmt.Fprintf(w, "hello\n")
}

func headers(w http.ResponseWriter, req *http.Request) {
	fmt.Println("Received request on /headers")
	for name, headers := range req.Header {
		for _, h := range headers {
			fmt.Fprintf(w, "%v: %v\n", name, h)
		}
	}
}

// You can create a basic HTTP server using the net/http package.
func main() {
	// net/http servers use handlers to register server behaviour.
	// A handler is an object implementing the http.Handler interface. A common
	// way to write a handler is by using the http.HandlerFunc adapter on
	// functions with the appropriate signature.
	// Functions serving as handlers take an http.ResponseWriter and an
	// http.Request as arguments. The response writer is used to fill in the
	// HTTP response.

	// A simple response on the /hello endpoint.
	http.HandleFunc("/hello", hello)

	// Reads in all the request headers and outputs them in the response.
	http.HandleFunc("/headers", headers)

	fmt.Println("Listening on localhost" + port)
	// ListenAndServe() activates the server with the port and a handler. nil
	// tells it to use the default router with the handlers configured above.
	http.ListenAndServe(port, nil)
}
