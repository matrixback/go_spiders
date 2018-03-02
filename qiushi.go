package main

import (
    "fmt"
    "github.com/PuerkitoBio/goquery"
    "log"
)

func ExampleScrape() {
    doc, err := goquery.NewDocument("http://www.qiushibaike.com")
    if err != nil {
        log.Fatal(err)
    }
    doc.Find(".article").Each(func(i int, s *goquery.Selection) {
        if s.Find(".thumb").Nodes == nil && s.Find(".video_holder").Nodes == nil {
            content := s.Find(".content").Text()
            fmt.Printf("%s", content)
        }
    })
}

func main() {
    ExampleScrape()
}
