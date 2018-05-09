package router

import (
	"fmt"
	"net/http"
)

func handleFunc(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, r.URL.Path)
}

type Route struct {
	handler    http.Handler
	handlerMap map[string]string
}

func (r *Route) Get(url string, handler func(w http.ResponseWriter, r *http.Request)) {

}
func (r *Route) Post(url string, handler func(w http.ResponseWriter, r *http.Request)) {

}
func (this *Route) ServeHTTP(w http.ResponseWriter, req *http.Request) {

}

func handler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("xx"))
}
func NewRouter() Route {
	r := Route{}
	return r
}

func Start() {

	r := NewRouter()
	r.Get("/hello/{:id}", func(w http.ResponseWriter, r *http.Request) {

	})
	r.Get("/url", func(w http.ResponseWriter, r *http.Request) {

	})
	r.Post("/url", func(w http.ResponseWriter, r *http.Request) {

	})
	server := http.Server{
		Addr:    ":8080",
		Handler: &r,
	}
	server.ListenAndServe()
}
