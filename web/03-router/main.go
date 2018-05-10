package router

import (
	"fmt"
	"net/http"
)

func handleFunc(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, r.URL.Path)
}

type urlHandler struct {
	method string
}
type Route struct {
	handler    http.Handler
	handlerMap map[string]interface{}
}

func (r *Route) Get(url string, handler func(w http.ResponseWriter, r *http.Request)) {
	r.handlerMap["method"] = url
	r.handlerMap["do"] = handler
}
func (r *Route) Post(url string, handler func(w http.ResponseWriter, r *http.Request)) {
	r.handlerMap["method"] = "GET"
}
func (this *Route) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	fmt.Println(this.handlerMap)
	// for _, v := range this.handlerMap {

	// }
}

func handler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("xx"))
}
func NewRouter() Route {
	r := Route{
		handlerMap: make(map[string]interface{}),
	}
	// r.handlerMap = make(map[string]interface{})
	return r
}

func Start() {

	r := NewRouter()
	r.Get("/hello/{:id}", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("router herer")
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
