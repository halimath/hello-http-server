package main

import (
	"flag"
	"fmt"
	"log"
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
	fmt.Printf("Starting HTTP server on 0.0.0.0:%d\n", *portFlag)
	initServer()
}

func initServer() {
	multiplexer := http.NewServeMux()
	multiplexer.HandleFunc("/", handleIndex)

	log.Fatal(http.ListenAndServe(fmt.Sprintf("0.0.0.0:%d", *portFlag), createAccessLoggingInterceptor(multiplexer)))
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
