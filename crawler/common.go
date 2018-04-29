package crawler

type Page struct {
	entry  string
	size   int
	title  string
	images []string
}

var (
	URL   string
	P1    []string // 一级页面url
	Pages []Page
)
