package main

import (
    "fmt"
)

type Fetcher interface {
    Fetch (url string) (title string, body string, urls []string, err error)
}

func Crawl(url string, depth int, fetcher Fetcher) {
    if depth <= 0 {
        return
    }

    title, body, urls, err := fetcher.Fetch(url)
    if err != nil {
        fmt.Println(err)
        return
    }
    save_html(title, body)
    for _, u := range urls {
        Crawl(u, depth - 1, fetcher)
    }
    return
}

func main() {
    Crawl("http://golang.org/", 4, fetcher)
}

func save_html(title, content string) error {
    fout, err := os.Create(title)
    if err != nil {
        fmt.Printf("Create file error, file name: %s \n", title)
        return err
    }

    fout.WriteString(content)
    return nil
}

type MyFetcher map[string]string;

func (m MyFetcher) Fetch (url string) (title string, body string, urls []string, err error) {

}

