package fileSrv

import (
	"net/http"
)

func Start() {
	http.ListenAndServe(":8080", http.FileServer(http.Dir("/Go/src")))
}
