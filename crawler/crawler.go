package crawler

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/Gavin-Gong/learn-go/practice/downloader"
	"github.com/PuerkitoBio/goquery"
)

func Start() {
	URL = "http://m.mzitu.com/"

	// 查询所有一级页面链接
	// fmt.Println(P1)

	getEnterPages()
}

func getEnterPages() {
	page := 176
	var url string
	for i := 0; i < page; i++ {
		if i == 0 {
			url = URL
		} else {
			url = URL + "page/" + strconv.Itoa(i)
		}

		// 处理一个入口列表页面
		doc, _ := goquery.NewDocument(url)
		doc.Find("#content figure > a").Each(func(i int, s *goquery.Selection) {
			entry, _ := s.Attr("href")
			title, _ := s.Attr("title")
			doc, _ := goquery.NewDocument(entry)
			str := doc.Find("#content .prev-next-page").Text()
			fmt.Println(str)
			size := parsePageSize(str)

			// 爬取列表页的一个图集
			fmt.Println("开始 爬取图集:", title)
			dir := title
			wd, _ := os.Getwd()
			path := wd + "/temp/mzitu/" + dir
			if _, err := os.Stat(path); err == nil {
				return
			}
			os.MkdirAll(path, 774)

			fileName := 0
			for i := 0; i < size; i++ {
				var url string
				if i == 0 {
					url = entry
				} else {
					url = entry + "/" + strconv.Itoa(i+1)
				}
				doc, _ := goquery.NewDocument(url)
				image, _ := doc.Find("#content figure img").Attr("src")
				fmt.Println("下载", image)
				fileName++
				sufix := "/" + strconv.Itoa(fileName) + ".jpg"
				err := downloader.DownLoad(path+sufix, image)
				if err == nil {
					fmt.Println("下载", url, "成功")
				} else {
					fmt.Println("下载", url, "失败")
				}
			}
		})
	}
}

// 解析页数
func parsePageSize(s string) int {
	fmt.Println(s)
	arr := strings.Replace(s, "页", "", -1)
	size := strings.Split(arr, "/")
	sz, _ := strconv.Atoi(size[1])
	return sz
}
