package main

import (
	"fmt"
	"log"
	"os"

	"github.com/gocolly/colly"
)

var link string
var fName string

func getDomain() {

}

func getUrls(link string) {

	c := colly.NewCollector()
	// search for all href links in 'a' tag [url/paths]
	c.OnHTML("a[href]", func(e *colly.HTMLElement) {
		// fmt.Println(e.Attr("href"))
		link := e.Attr("href")
		getUrls(link)
	})
	// print informational message
	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting: ", r.URL.String())
	})
	// visit the address/website
	c.Visit("https://www.kitaa.me")
	fmt.Println("Completed!")
}

func saveUrls(fName string) {

	urls_ := fName
	file, err := os.OpenFile(urls_, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatalf("Could not create file, err :%q", err)
		return
	}
	defer file.Close()
	// write to file
	if _, err := file.WriteString(link + "\n"); err != nil {
		log.Fatalf("Could not write to file, err :%q", err)
		return
	}
}

func printUrls() {

}
