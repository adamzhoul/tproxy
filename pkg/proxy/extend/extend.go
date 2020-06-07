package extend

import (
	"log"
	"net/http"
	"sync"
)

var (
	httpModifier map[string]func(r *http.Request)
	mutex        sync.Mutex
)

func AddHttpModifierExtend(k string, f func(r *http.Request)) {
	mutex.Lock()
	defer mutex.Unlock()

	if httpModifier == nil {
		httpModifier = make(map[string]func(*http.Request))
	}

	if _, ok := httpModifier[k]; ok {
		log.Println("modifier exists, pass ", k)
		return
	}
	httpModifier[k] = f
}

func GetDirector() func(*http.Request) {
	return func(r *http.Request) {

		for k, f := range httpModifier {
			log.Println("run :modifier", k)
			f(r)
		}
	}
}
