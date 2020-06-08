package extend

import (
	"fmt"
	"net/http"
	"regexp"
	"strings"
)

type virtualservice struct {
	Host string
	Http []rule
}

type rule struct {
	Match match
	Route route
}

type match struct {
	CookieRex string
	UrlPrefix string
}

// for now only route to other containers in the same physical machine
// so port change is enough
// when mulit machine comes, registry center is required
type route struct {
	DestinationPort int
}

var config map[string]virtualservice

func init() {

	// load config

	// add extend
	AddHttpModifierExtend("virtualservice", vsModifyRequest)
}

// simulate istio virtualservice
// redirect by virtualservice config
// Host:
//    xxx.com
// Http:
//	 Match:
//		CookieRex || UrlPrefix
//   Route:
//		DestinationPort
func vsModifyRequest(r *http.Request) {

	if _, ok := config[r.Host]; !ok {
		return
	}

	for _, h := range config[r.Host].Http {

		if cookieRex(h.Match.CookieRex, r.Header.Get("cookie")) &&
			urlPrefix(h.Match.UrlPrefix, r.URL.Path) {

			// match rule, route to destination
			// change dst port only for now
			r.Host = fmt.Sprintf("%s:%d", r.Host, h.Route.DestinationPort)
			r.URL.Host = r.Host
		}
	}
}

func cookieRex(rule, cookie string) bool {

	if rule == "" {
		return true
	}
	r, _ := regexp.Compile(rule)
	return r.MatchString(cookie)
}

func urlPrefix(rule, path string) bool {
	if rule == "" {
		return true
	}
	return strings.HasPrefix(rule, path)
}
