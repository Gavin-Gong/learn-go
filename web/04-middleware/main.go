package middleware

import (
	"fmt"
	"net/http"
)

func loggerMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("start logger")
		next.ServeHTTP(w, r)
		fmt.Println("end logger")
	})
}
func authMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		next.ServeHTTP(w, r)

		fmt.Println("auth")
	})
}

func handlerMsg(w http.ResponseWriter, r *http.Request) {
	fmt.Println("start handle msg")
}

func Start() {
	http.Handle("/message", authMiddleware(loggerMiddleware(http.HandlerFunc(handlerMsg))))
	http.ListenAndServe(":8080", nil)
}
