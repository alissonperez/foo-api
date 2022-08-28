package v1

import (
	"net/http"

	"github.com/alissonperez/api-foo/net/v1/handler"
)

func WrapDependenciesHandler(handlerFunc func(w http.ResponseWriter, r *http.Request, deps *handler.HandlerDependencies), deps *handler.HandlerDependencies) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		handlerFunc(w, r, deps)
	}
}
