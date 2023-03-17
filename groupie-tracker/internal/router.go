package internal

import (
	"log"
	"net/http"
)

type Router struct{
	Mux *http.ServeMux
}
func NewHandler()*Router{
	router := *http.NewServeMux()
	serveMux := Router{&router}
	serveMux.Register(serveMux.Mux)
	return &serveMux
}
func middleware(h http.HandlerFunc) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf(
			"%s %s\n",
			r.Method,
			r.URL,
		)
		h.ServeHTTP(w, r)
	})
}
func (h *Router) Register(router *http.ServeMux){
	files := http.FileServer(http.Dir("style"))
	router.Handle(root, middleware(h.Home))
	router.Handle(trackerURL, middleware(h.trackerHandler))
	router.Handle(searchUrl, middleware(h.Search))
	router.Handle(style, middleware(http.StripPrefix("/style/", files).ServeHTTP))
}