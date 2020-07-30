package main

import (
	"net/http"
	"text/template"

	"github.com/fernok/randomBaekJoon/runRandom"
)

// RandomResult is the result of random searching
type RandomResult struct {
	Address string
	Title   string
	Account string
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	userID := r.FormValue("account")
	var p RandomResult
	if len(userID) == 0 {
		p = RandomResult{Address: "enter your name", Title: "", Account: "guest"}
	} else {
		title, url := runRandom.RunRandom(userID)
		p = RandomResult{Address: url, Title: title + " : ", Account: userID}
	}
	t, _ := template.ParseFiles("index.html")
	t.Execute(w, p)
}

func main() {
	http.HandleFunc("/", indexHandler)
	http.ListenAndServe(":8000", nil)
}
