package fileSrv

import (
	"fmt"
	"net"
	"net/http"
)

func Start() {
	addrs, _ := net.InterfaceAddrs()
	for _, address := range addrs {
		if ipnet, ok := address.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				fmt.Println(ipnet.IP.String() + ":8080")
			}
		}
	}
	http.ListenAndServe(":8080", http.FileServer(http.Dir("/Go/src")))
}
