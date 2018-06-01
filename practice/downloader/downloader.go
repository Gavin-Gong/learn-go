package downloader

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

func Start() {
	err := DownLoad("./temp/xx/xx.jpg", "http://i.meizitu.net/2018/04/25c01.jpg")
	if err != nil {
		println(err.Error())
	}
}
func ensureDir(dir string) {
	fmt.Println(dir)
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		dirErr := os.MkdirAll(dir, os.ModePerm)
		if dirErr != nil {
			fmt.Println(dirErr.Error())
		}
	}
}

func DownLoad(path string, url string) error {
	pathArr := strings.Split(path, "/")
	dir := strings.Join(pathArr[0:len(pathArr)-1], "/")
	ensureDir(dir)

	out, err := os.Create(path)
	if err != nil {
		return err
	}
	defer out.Close()

	// get date by url
	client := &http.Client{}
	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Set("Referer", url)

	res, err := client.Do(req)
	defer res.Body.Close()
	if err != nil {
		return err
	}

	// 写入文件
	_, err = io.Copy(out, res.Body)
	if err != nil {
		return err
	}
	return nil
}
