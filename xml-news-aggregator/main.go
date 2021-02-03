package main

import (
	"fmt"
	"net/http"

	"github.com/URAmiRBin/go-learning/sentdex/handler"
	"github.com/URAmiRBin/go-learning/sentdex/parser"
)

func main() {
	n, err := parser.GetNews("https://www.washingtonpost.com/news-sitemaps/index.xml")
	if err != nil {
		fmt.Println(err)
		return
	}

	np := handler.NewsPage{Title: "NEWS AGGREGATOR", News: n}

	// First arg is the path to handle
	// Second is the function to run when a request is sent to this path
	http.HandleFunc("/", handler.Homehandler)
	http.HandleFunc("/news/", np.Newshandler)
	// second argument is about host, nil is local
	http.ListenAndServe(":8000", nil)
	// Use ListenAndServeTLS for https
}
