package router

import (
	"fmt"
	"net/http"
	"regexp"
)

type registerHandle func(w http.ResponseWriter, r *http.Request)

type urlHandler struct {
	method    string
	url       string
	regexpUrl *regexp.Regexp
	handle    registerHandle
	params    map[string]string
}
type Route struct {
	handler    http.Handler
	handlerMap map[string]urlHandler
}

func (r *Route) Get(url string, handler func(w http.ResponseWriter, r *http.Request)) {
	r.handlerMap["GET:"+url] = urlHandler{
		method:    "GET",
		url:       url,
		handle:    handler,
		params:    make(map[string]string),
		regexpUrl: regexp.MustCompile(url),
	}
	// hello/{:id}/{:uid}
	// hello/ss/xx
	// hello/xx/xx/xx
	reg := regexp.MustCompile(`{:([a-zA-Z0-9]+)}`)
	paramsArr := reg.FindStringSubmatch(url)
	for _, v := range paramsArr {
		r.handlerMap["GET"+url].params[v] = ""
	}
}
func (r *Route) Post(url string, handler func(w http.ResponseWriter, r *http.Request)) {
	r.handlerMap["POST:"+url] = urlHandler{
		method: "POST",
		url:    url,
		handle: handler,
	}
}
func (this *Route) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	for _, v := range this.handlerMap {
		method := v.method
		// match correct router
		if req.Method == method && v.regexpUrl.MatchString(req.URL.Path) {
			v.handle(w, req)
		}
	}
}

func handler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("xx"))
}
func NewRouter() Route {
	r := Route{
		handlerMap: make(map[string]urlHandler),
	}
	return r
}

func Start() {

	r := NewRouter()
	r.Get("/hello/{:id}", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("route: hello/{:id}")
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
