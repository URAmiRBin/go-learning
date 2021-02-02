package parser

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/URAmiRBin/go-learning/sentdex/model"
)

type SitemapIndex struct {
	Locations []string `xml:"sitemap>loc"`
}

type NewsList struct {
	Titles    []string `xml:"url>news>title"`
	Keywords  []string `xml:"url>news>keywords"`
	Locations []string `xml:"url>loc"`
}

func GetNews(url string) ([]model.News, error) {
	fmt.Println("Getting sitemap indices")
	s, err := getSitemapIndex(url)
	if err != nil {
		fmt.Println("Error in getting sitemap")
		return []model.News{}, err
	}

	fmt.Println("Crawling news...")
	n, err := crawlNews(s)
	if err != nil {
		fmt.Println("Error in crawling news")
		return []model.News{}, err
	}

	fmt.Println("Converting")
	return convertNews(n), nil
}

func convertNews(n NewsList) []model.News {
	fmt.Println(len(n.Titles))
	fmt.Println(len(n.Keywords))
	fmt.Println(len(n.Locations))
	modelNews := make([]model.News, 0)
	for idx := range n.Titles {
		modelNews = append(modelNews, model.News{Title: n.Titles[idx], Loc: n.Locations[idx]})
	}
	return modelNews
}

func crawlNews(s SitemapIndex) (NewsList, error) {
	var n NewsList
	for idx, Location := range s.Locations {
		if idx > 1 {
			break
		}
		fmt.Printf("Crawling %s\n", Location)
		resp, _ := http.Get(strings.TrimSpace(Location))
		bytes, _ := ioutil.ReadAll(resp.Body)
		xml.Unmarshal(bytes, &n)
	}
	return n, nil
}

func getSitemapIndex(url string) (SitemapIndex, error) {
	fmt.Println("started fetching")
	resp, err := http.Get(url)
	if err != nil {
		return SitemapIndex{}, err
	}

	bytes, _ := ioutil.ReadAll(resp.Body)
	resp.Body.Close()
	fmt.Println("finished fetching")

	var s SitemapIndex
	xml.Unmarshal(bytes, &s)
	return s, nil

}
