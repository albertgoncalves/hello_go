package main

import (
    "fmt"
    "github.com/PuerkitoBio/goquery"
    "os"
)

func panicIf(err error) {
    if err != nil {
        panic(err)
    }
}

func printHref(i int, s *goquery.Selection) {
    href, exists := s.Attr("href")
    if exists {
        fmt.Printf("%d %s\n", i, href)
    }
}

func main() {
    const fn string = "index.html"

    f, err := os.Open(fn)
    panicIf(err)
    defer f.Close()

    html, err := goquery.NewDocumentFromReader(f)
    panicIf(err)

    html.Find("a").Each(printHref)
}
