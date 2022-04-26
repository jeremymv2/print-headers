package main

import (
	"fmt"
	"log"
	"net/http"
	"net/http/httputil"
)

const PORT = 8080

func main() {
	startServer(handler)
}

func startServer(handler func(http.ResponseWriter, *http.Request)) {
	http.HandleFunc("/", handler)
	log.Printf("Starting server. Listening on %d..", PORT)
	http.ListenAndServe(fmt.Sprintf(":%d", PORT), nil)
}

func handler(w http.ResponseWriter, r *http.Request) {
	reqDump, err := httputil.DumpRequest(r, true)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("REQUEST\n%s", string(reqDump))
	resp := fmt.Sprintf("Hello %s", r.Header.Get("User-Agent"))
	_, err = w.Write([]byte(resp))
	if err != nil {
		log.Panicf("not able to write http output: %s", err)
	}
}
