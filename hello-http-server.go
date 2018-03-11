package main

import (
	"flag"
	"fmt"
	"log"
	"net"
	"net/http"
	"time"
)

var (
	portFlag = flag.Int("port", 8080, "HTTP port to bind to")
	version  = "n/a"
)

func main() {
	flag.Parse()
	fmt.Printf("HTTP Hello Server v%s\n", version)
	fmt.Printf("Starting HTTP server on :%d\n", *portFlag)
	initServer()
}

func initServer() {
	multiplexer := http.NewServeMux()
	multiplexer.HandleFunc("/", handleIndex)

	listen, err := net.Listen("tcp4", ":8080")
	if err != nil {
		log.Fatal(err)
	}
	log.Fatal(http.Serve(listen, createAccessLoggingInterceptor(multiplexer)))

	// log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", *portFlag), ))
}

func handleIndex(w http.ResponseWriter, req *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Hello, world\n\nHTTP Hello Server v%s", version)
}

func createAccessLoggingInterceptor(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(r http.ResponseWriter, req *http.Request) {
		fmt.Printf("%s %s %s %s\n", time.Now(), req.RemoteAddr, req.Method, req.URL)
		handler.ServeHTTP(r, req)
	})
}
