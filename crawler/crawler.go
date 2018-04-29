package crawler

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

func Start() {
	URL = "http://m.mzitu.com/"

	// 查询所有一级页面链接
	// fmt.Println(P1)

	getEnterPages()
}

//
func getEnterPages() {
	// name :=
	page := 1
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
		fmt.Println("开始爬取图集")
		for _, page := range Pages {
			recv := make(chan string, page.size)
			go func() {
				for i := 0; i < page.size; i++ {
					var url string
					if i == 0 {
						url = page.entry
					} else {
						url = page.entry + strconv.Itoa(i+1)
					}
					doc, _ := goquery.NewDocument(url)
					image, _ := doc.Find("#content figure img").Attr("src")
					fmt.Println(image)
					recv <- image
				}
			}()
			str := <-recv
			fmt.Println("chan", str)
			page.images = append(page.images, str)
		}
		// fmt.Println(Pages)
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
