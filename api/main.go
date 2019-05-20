package main

import (
	"fmt"
	log "github.com/sirupsen/logrus"
	"net/http"
)

func main() {
	http.Handle("/api/v1/info", loggingMiddleware(http.HandlerFunc(handler)))
	http.ListenAndServe(":9527", nil)
}

func handler(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "this is go istio api : 9527 ")
	//log.Infof("istio api info : %s", req.RequestURI)
	//tr := http.Transport{DisableKeepAlives: true}
	//client := &http.Client{Transport: &tr}

	//resp, err := client.Get("http://127.0.0.1:9528/repo/v1/info")
	//
	//if resp != nil {
	//	defer resp.Body.Close()
	//}
	//checkError(err)
	//
	//body, err := ioutil.ReadAll(resp.Body)
	//checkError(err)
	//
	////fmt.Println(string(body))
	//
	//fmt.Fprintf(w, string(body))
}

func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		log.Infof("uri: %s", req.RequestURI)
		next.ServeHTTP(w, req)
	})
}

func checkError(err error) {
	if err != nil{
		log.Fatalln(err)
	}
}