package handler

import (
	"html/template"
	"net/http"

	"github.com/URAmiRBin/go-learning/sentdex/model"
)

type NewsPage struct {
	Title string
	News  []model.News
}

func (np NewsPage) Newshandler(w http.ResponseWriter, r *http.Request) {
	p := NewsPage{Title: np.Title, News: np.News}
	t, _ := template.ParseFiles("news.html")
	t.Execute(w, p)
}
