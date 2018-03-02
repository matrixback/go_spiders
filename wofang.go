package main

import (
    "fmt"
    "strconv"
    "time"
    "github.com/PuerkitoBio/goquery"
    "log"
    "bytes"
    "encoding/csv"
    "os"
)

func search() {
    a := 0
    fileName := "wofang.csv"
    buf := new(bytes.Buffer)
    r2 := csv.NewWriter(buf)
    for i := 1; i <= 10; i++ {
        fmt.Println("正在抓取第" + strconv.Itoa(i) + "页......")
        url := "http://www.wofang.com/building/p" + strconv.Itoa(i) + "/"
        if i == 1 {
            url = "http://www.wofang.com/building/"
        }

        doc, err := goquery.NewDocument(url)
        if err != nil {
            log.Fatal(err)
            continue
        }

        doc.Find(".m ul li").Each(func(i int, s *goquery.Selection) {
            name := s.Find(".title a").Text()
            location := s.Find(".time").Text()
            price := s.Find(".sale-price font").Text()
            if price != "" {
                a++
                s := make([]string, 3)
                s[0] = name
                s[1] = price
                s[2] = location
                r2.Write(s)
                r2.Flush()
                fmt.Printf("%s,%s,%s\n", name, price, location)
            }
        })
        time.Sleep(2 * time.Second)
    }
    fout, err := os.Create(fileName)
    defer fout.Close()
    if err != nil {
        fmt.Println(fileName, err)
        return
    }
    fout.WriteString(buf.String())
    fmt.Print(a)
}

func main() {
    t1 := time.Now()
    search()
    elapsed := time.Since(t1)
    fmt.Println("")
    fmt.Println("爬虫结束，总共耗时：", elapsed)
}


