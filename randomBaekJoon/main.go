package main

import (
	"fmt"
	"net/http"
	"text/template"

	"github.com/fernok/randomBaekJoon/runRandom"
)

// RandomResult is the result of random searching
type RandomResult struct {
	Address string
	Message string
	Title   string
	Account string
}

var problemsToSolve []runRandom.ExtractedProblemLinks

func indexHandler(w http.ResponseWriter, r *http.Request) {
	userID := r.FormValue("account")
	var p RandomResult
	if len(userID) == 0 {
		p = RandomResult{Address: "", Message: "Enter your name!", Title: "", Account: "guest"}
	} else {
		var userSolvedProblems = runRandom.GetUserSolvedProblemInfo(userID)
		if userSolvedProblems[0] == -1 {
			p = RandomResult{Address: "", Message: "The user does not exist!", Title: "", Account: "guest"}
		} else {
			title, url := runRandom.RunRandom(problemsToSolve, userSolvedProblems)
			p = RandomResult{Address: url, Message: "", Title: title + " : ", Account: userID}
		}
	}
	t, _ := template.ParseFiles("index.html")
	t.Execute(w, p)
}

func main() {
	problemsToSolve = runRandom.GetPage()

	fmt.Println("Server Running on PORT 8000")

	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	http.HandleFunc("/", indexHandler)
	http.ListenAndServe(":8000", nil)
}
