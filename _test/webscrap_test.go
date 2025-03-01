package test

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/gocolly/colly/v2"
)

const SUrl = "https://dle.rae.es/usar"

func TestWS(t *testing.T) {
	t.Log("Ok?")
	c := colly.NewCollector(colly.AllowedDomains("dle.rae.es"))
	c.OnHTML("article", func(h *colly.HTMLElement) {
		t.Log("Ok?2")
		t.Log(h.Text)
	})
	c.OnRequest(func(r *colly.Request) {
		t.Log("Visiting", r.URL)	
	})
	c.Visit(SUrl)
}

func TestWS2(t *testing.T) {
	
	c := colly.NewCollector(
		// Visit only domains: hackerspaces.org, wiki.hackerspaces.org
		colly.AllowedDomains("hackerspaces.org", "wiki.hackerspaces.org"),
	)

	// On every a element which has href attribute call callback
	c.OnHTML("a[href]", func(e *colly.HTMLElement) {
		link := e.Attr("href")
		// Print link
		t.Logf("Link found: %q -> %s\n", e.Text, link)
		// Visit link found on page
		// Only those links are visited which are in AllowedDomains
		c.Visit(e.Request.AbsoluteURL(link))
	})

	// Before making a request print "Visiting ..."
	c.OnRequest(func(r *colly.Request) {
		t.Log("Visiting", r.URL.String())
	})

	// Start scraping on https://hackerspaces.org
	c.Visit("https://hackerspaces.org/")

}

func TestWS3(t *testing.T) {
	c := colly.NewCollector()
	c.AllowURLRevisit = false
//	extensions.RandomUserAgent(c)
	c.UserAgent = "Mozilla/5.0 (Windows; U; Windows NT 5.1; en-US; rv:1.9.0.7) Gecko/2009021910 Firefox/3.0.7";

	c.OnHTML("h1.c-page-header__title", func(elem *colly.HTMLElement) {
		fmt.Println(elem.Text)
	})

	fmt.Println(c.Visit(SUrl))
}

func TestWS4(t *testing.T) {
	rq := http.NewRequest(http.MethodGet, SUrl, nil)
	rq.
}
