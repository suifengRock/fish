package main

import (
	"fmt"
	"log"

	"github.com/PuerkitoBio/goquery"
)

func ExampleScrape() {
	doc, err := goquery.NewDocument("http://0.0.0.0:8800/admin/login")
	if err != nil {
		log.Fatal(err)
	}

	val, _ := doc.Find("input[name=_xsrf]").Attr("value")
	fmt.Println(val)
}
func main() {
	ExampleScrape()
}
