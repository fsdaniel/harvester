package ui

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rancher/apiserver/pkg/middleware"
)

func RegisterAPIUI() http.Handler {
	r := mux.NewRouter()
	r.UseEncodedPath()

	//API UI
	uiContent := middleware.NewMiddlewareChain(middleware.Gzip, middleware.DenyFrameOptions,
		middleware.CacheMiddleware("json", "js", "css")).Handler(Content())

	r.PathPrefix("/api-ui").Handler(uiContent)
	r.NotFoundHandler = http.NotFoundHandler()
	return r
}
