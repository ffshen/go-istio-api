package main

import (
	"fmt"
	"net/http"
	log "github.com/sirupsen/logrus"
)

func main() {
	http.Handle("/api/v1/info", loggingMiddleware(http.HandlerFunc(handler)))
	http.ListenAndServe(":9527", nil)
}

func handler(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "this is go istio api : 9527 ")
}

func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		log.Infof("uri: %s", req.RequestURI)
		next.ServeHTTP(w, req)
	})
}