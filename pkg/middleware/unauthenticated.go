package middleware

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
)

func Unauthenticated(h httprouter.Handle, args MiddlewareArguments) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		h(w, r.WithContext(r.Context()), ps)
	}
}
