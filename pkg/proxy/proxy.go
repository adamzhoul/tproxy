package proxy

import (
	"github.com/adamzhoul/tproxy/pkg/proxy/extend"
	"net/http"
	"net/http/httputil"
)

type Tproxy struct {
	rp httputil.ReverseProxy
}

func NewTproxy() *Tproxy {

	t := Tproxy{
		rp: httputil.ReverseProxy{
			Director: extend.GetDirector(),
		},
	}
	return &t
}

func (t *Tproxy) ServeHTTP(rw http.ResponseWriter, req *http.Request) {

	t.rp.ServeHTTP(rw, req)
}
