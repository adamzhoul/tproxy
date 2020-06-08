package extend

import (
	"log"
	"net/http"
	"sync"
)

var (
	httpFilter           map[string]func(r *http.Request)
	httpModifier         map[string]func(r *http.Request)
	httpResponseModifier map[string]func()
	mutex                sync.Mutex
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

		// execute sequence is very important
		// TODO: for now, only 1 extend , do it later
		for k, f := range httpModifier {
			log.Println("run :modifier", k)
			f(r)
		}
	}
}
