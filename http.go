package iter

import (
	"net/http"
)

type httpIt struct {
	ch chan Responder[*http.Request, http.ResponseWriter]
}

func HTTP() Iter[Responder[*http.Request, http.ResponseWriter]] {
	return httpIt{ch: make(chan Responder[*http.Request, http.ResponseWriter])}
}

func (h httpIt) ForEach(f func(Responder[*http.Request, http.ResponseWriter]) (stop bool)) {
	for resp := range h.ch {
		if f(resp) {
			return
		}
	}
}

func (h httpIt) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	h.ch <- Responder[*http.Request, http.ResponseWriter]{req, w}
}
