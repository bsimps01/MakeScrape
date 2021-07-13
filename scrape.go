package main

import (
	"fmt"

	"github.com/gocolly/colly"
)

type Article struct {
	ID       string `json:"id"`
	HEADLINE string `json:"headline"`
	TITLE    string `json:"title"`
	LINK     string `json:"link"`
}

// main() contains code adapted from example found in Colly's docs:
// http://go-colly.org/docs/examples/basic/
func main() {

	//allIDs := make([]Article, 0)
	// Instantiate default collector
	c := colly.NewCollector()
	stories := []Article{}

	// On every a element which has href attribute call callback
	c.OnHTML(".l-grid__content > a", func(e *colly.HTMLElement) {
		//link := e.Attr("href")

		story := Article{}
		story.HEADLINE = e.ChildText(".l-grid__content-title u-text-h5")
		story.TITLE = e.ChildText(".p-headlkine-stack__link")
		story.LINK = e.ChildText(".p-headline-stack__link")
		stories = append(stories, story)

		// Print link
		fmt.Printf("Headline: %q\n, Title: %q\n, Link: %q\n, ", story.HEADLINE, story.TITLE, story.LINK)
	})

	// Before making a request print "Visiting ..."
	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL.String())
	})

	// Start scraping on https://hackerspaces.org
	c.Visit("http://www.mlb.com/")
}
