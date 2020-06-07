package main

import (
	"log"
	"net/http"

	"github.com/adamzhoul/tproxy/pkg/proxy"
)

func main() {

	proxy := proxy.NewTproxy()
	log.Fatal(http.ListenAndServe(":9090", proxy))
}
