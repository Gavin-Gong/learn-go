package main

// import _ "github.com/Gavin-Gong/learn-go/example"
import (
	"fmt"
	"os"

	"github.com/Gavin-Gong/learn-go/crawler"
)

// import _ "github.com/Gavin-Gong/learn-go/practice/downloader"
// import "github.com/Gavin-Gong/learn-go/practice/spotlight-pic"
// import "github.com/Gavin-Gong/learn-go/practice/downloader"

func main() {
	crawler.Start()
	// downloader.Start()
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
