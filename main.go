package main

// import _ "github.com/Gavin-Gong/learn-go/example"
import (
	"fmt"
	"os"

	_ "github.com/Gavin-Gong/learn-go/crawler"
	"github.com/Gavin-Gong/learn-go/web/01-basic"
)

// import _ "github.com/Gavin-Gong/learn-go/practice/downloader"
// import "github.com/Gavin-Gong/learn-go/practice/spotlight-pic"
// import "github.com/Gavin-Gong/learn-go/practice/downloader"

func main() {
	// crawler.Start()
	// downloader.Start()
	basic.Start()
}

func inits() {
	// path := "./test/main.go"
	wd, _ := os.Getwd()
	os.MkdirAll(wd+"/test", 774)
	file, err := os.Create(wd + "/test" + "/main.go")
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()
}
