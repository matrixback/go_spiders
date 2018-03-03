package main

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"strings"
	"net/http"
	"os"
	"runtime"
	"sync"
)

var URL string = "http://www.yinwang.org"

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	fmt.Println(URL)
	doc, err := goquery.NewDocument(URL)
	if err != nil {
		fmt.Println("没有获取到文件")
		fmt.Println(err)
		return
	}
	links := make([]string, 0)
	doc.Find("a").Each(func(i int, s *goquery.Selection) {
		link, exists := s.Attr("href")
		if exists == true {
			if !strings.HasPrefix(link, "http") {
				link = URL + link
			}
			links = append(links, link)
		}
	})
	fmt.Println(links)
	wg := sync.WaitGroup{}
	linksLength := len(links)
	linksLength = 5
	wg.Add(linksLength)
	for i := 0; i < linksLength; i++ {
		go SaveHtml(&wg, links[i])
	}
	wg.Wait()


}

func SaveHtml(wg *sync.WaitGroup, url string) {
	defer wg.Done()
	if !strings.HasPrefix(url, "http") {
		fmt.Println("url error")
		fmt.Println(url)
		return
	}

	resp, err := http.Get(url)
	if err != nil {
		fmt.Println(err)
		return
	}

	doc, err := goquery.NewDocumentFromResponse(resp)
	if err != nil {
		fmt.Println(err)
		return
	}
	title := doc.Find("title").Text()
	title += ".html"
	body := doc.Find("body").Text()
	fmt.Println(title)
	f, err := os.Create(title)
	if err != nil {
		fmt.Println("创建文件失败")
	}
	defer f.Close()
	f.WriteString(body)

}
