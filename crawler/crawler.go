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
	page := 147
	var url string
	for i := 0; i < page; i++ {
		if i == 0 {
			url = URL
		} else {
			url = URL + "page/" + strconv.Itoa(i)
		}

		doc, _ := goquery.NewDocument(url)
		doc.Find("#content figure > a").Each(func(i int, s *goquery.Selection) {
			href, _ := s.Attr("href")
			fmt.Println("获取url页面: %s", href)
			P1 = append(P1, href)
		})
	}
}
func getDetailPages(url string) {
	doc, _ := goquery.NewDocument(url)
	str := doc.Find("#content .prev-next-page").Text()
	size := parsePageSize(str)
	fmt.Println(size)
}

func parsePageSize(s string) int {
	arr := strings.Replace(s, "页", "", -1)
	size := strings.Split(arr, "/")
	sz, _ := strconv.Atoi(size[1])
	return sz
}
