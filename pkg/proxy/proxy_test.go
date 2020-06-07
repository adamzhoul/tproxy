package proxy

import (
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"testing"
)

func TestProxy(t *testing.T) {

	proxy, _ := url.Parse("http://127.0.0.1:9090")
	tr := &http.Transport{
		Proxy: http.ProxyURL(proxy),
	}

	req, _ := http.NewRequest("GET", "http://www.zl.com", nil)
	client := &http.Client{}
	client.Transport = tr
	res, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()

	body, _ := ioutil.ReadAll(res.Body)
	log.Println(string(body))
}
