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

func handler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("xx"))
}
func NewRouter() Route {
	r := Route{}
	r.handler = http.HandleFunc("/", handler)
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

	if err := http.ListenAndServe(":8080", r.handler); err != nil {
		fmt.Println(err)
	}
}
