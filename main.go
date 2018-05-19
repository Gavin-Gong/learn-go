package main

// import _ "github.com/Gavin-Gong/learn-go/example"
import (
	_ "github.com/Gavin-Gong/learn-go/crawler"
	middleware "github.com/Gavin-Gong/learn-go/web/04-middleware"
)

// import _ "github.com/Gavin-Gong/learn-go/practice/downloader"
// import "github.com/Gavin-Gong/learn-go/practice/spotlight-pic"
// import "github.com/Gavin-Gong/learn-go/practice/downloader"

func main() {
	// crawler.Start()
	// downloader.Start()
	// basic.Start()
	// fileSrv.Start()
	// router.Start()
	middleware.Start()
}
