package downloader

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func init() {
	DownLoad("./xx.jpg", "http://i.meizitu.net/2018/04/25c01.jpg")
}
func DownLoad(path string, url string) error {
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
	fmt.Println("下载成功")
	return nil
}
