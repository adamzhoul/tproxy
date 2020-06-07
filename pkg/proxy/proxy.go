package proxy

import (
	"github.com/adamzhoul/tproxy/pkg/proxy/extend"
	"net/http"
	"net/http/httputil"
)

type Tproxy struct {
	httputil.ReverseProxy
}

func NewTproxy() *Tproxy {

	t := Tproxy{}
	t.Director = extend.GetDirector()
	return &t
}

func NewTProxyWithModifierExtend(k string, f func(*http.Request)) *Tproxy {

	t := Tproxy{}
	extend.AddHttpModifierExtend(k, f)
	t.Director = extend.GetDirector()

	return &t
}
