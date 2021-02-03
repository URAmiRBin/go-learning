package handler

import (
	"fmt"
	"net/http"
)

func Homehandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, `<h1>Welcome</h1>
<p> Hello my friend </p>
<p> Add your variables %s </p>
	`, "<strong> here </strong>")
	fmt.Println(*r)
}
