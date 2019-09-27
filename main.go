package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"time"
)

func main() {
	port := flag.String("port", "8080", "Port to listen to")
	flag.Parse()
	log.Printf("Echo server is listening on port %s", *port)

	http.HandleFunc("/", HelloServer)
	if err := http.ListenAndServe(":"+*port, nil); err != nil && err != http.ErrServerClosed {
		log.Fatalf("Could not listen on %s: %v", *port, err)
	}
	log.Println("Server stopped")

}

func HelloServer(w http.ResponseWriter, r *http.Request) {
	start := time.Now()
	log.Printf(
		"%s %s %s %v",
		r.Method,
		r.RequestURI,
		r.RemoteAddr,
		time.Since(start),
	)

	value := r.URL.Path[1:]
	if value == "" {
		value = "there"
	}

	fmt.Fprintf(w, "Hello, %s!", value)
}
