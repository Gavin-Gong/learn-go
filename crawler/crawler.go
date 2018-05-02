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

		doc, _ := goquery.NewDocument(url)
		doc.Find("#content figure > a").Each(func(i int, s *goquery.Selection) {
			url, _ := s.Attr("href")
			title, _ := s.Attr("title")
			doc, _ := goquery.NewDocument(url)
			str := doc.Find("#content .prev-next-page").Text()
			size := parsePageSize(str)

			Pages = append(Pages, Page{entry: url, title: title, size: size})
			fmt.Println("获取入口页面:", title)
		})
		fmt.Println("开始 爬取图集")
		for idx, page := range Pages {
			recv := make(chan string, page.size)
			go func() {
				for i := 0; i < page.size; i++ {
					var url string
					if i == 0 {
						url = page.entry
					} else {
						url = page.entry + "/" + strconv.Itoa(i+1)
					}
					// fmt.Println(url)
					doc, _ := goquery.NewDocument(url)
					image, _ := doc.Find("#content figure img").Attr("src")
					recv <- image
				}
			}()
			for i := 0; i < page.size; i++ {
				str := <-recv
				// fmt.Println(page.images, str)
				Pages[idx].images = append(Pages[idx].images, str)
			}
		}
	}
	fmt.Println("获取所有页面数据完成")
	for _, page := range Pages {
		fmt.Println("开始下载", page.title)
		dir := page.title
		wd, _ := os.Getwd()
		path := wd + "/temp/mzitu/" + dir
		if _, err := os.Stat(path); err == nil {
			continue
		}
		os.MkdirAll(path, 774)
		fileName := 0
		for _, url := range page.images {
			fileName++
			sufix := "/" + strconv.Itoa(fileName) + ".jpg"
			err := downloader.DownLoad(path+sufix, url)
			if err == nil {
				fmt.Println("下载", url, "成功")
			} else {
				fmt.Println("下载", url, "失败")
			}
		}
	}
}
func getDetailPages(url string) {
	doc, _ := goquery.NewDocument(url)
	str := doc.Find("#content .prev-next-page").Text()
	size := parsePageSize(str)
	fmt.Println(size)
}

// 解析页数
func parsePageSize(s string) int {
	arr := strings.Replace(s, "页", "", -1)
	size := strings.Split(arr, "/")
	sz, _ := strconv.Atoi(size[1])
	return sz
}
