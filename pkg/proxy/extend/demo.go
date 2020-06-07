package extend

import (
	"log"
	"net/http"
)

// regist extend to proxy director
func init() {

	AddHttpModifierExtend("demo", demoModifyRequest)
}

func demoModifyRequest(r *http.Request) {

	log.Println("demo request", r.URL, r.URL.Port(), "---", r.URL.Scheme)
}
