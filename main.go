package main

import (
	"log"
	"net/http"

	"github.com/adamzhoul/tproxy/pkg/proxy"
)

func main() {

	proxy := proxy.NewTproxy()
	log.Println("proxy server on :9090")
	log.Fatal(http.ListenAndServe(":9090", proxy))
}
