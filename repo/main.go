package main

import (
	"fmt"
	log "github.com/sirupsen/logrus"
	"net/http"
)

func main() {
	http.Handle("/repo/v1/info", loggingMiddleware(http.HandlerFunc(handler)))
	http.ListenAndServe(":9528", nil)
}

func handler(w http.ResponseWriter, req *http.Request) {
	//fmt.Println(string(body))
	fmt.Fprintf(w, "this is go istio repo : 9528 version : v2 ")
}

func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		log.Infof("uri: %s", req.RequestURI)
		next.ServeHTTP(w, req)
	})
}
