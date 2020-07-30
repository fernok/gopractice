package runRandom

import (
	"crypto/rand"
	"fmt"
	"log"
	"math/big"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/PuerkitoBio/goquery"
)

type extractedProblemLinks struct {
	tagTitle      string
	problemNumber int
	title         string
	rank          string
}

// RunRandom generates a string of a random BaekJoon problem
func RunRandom(user string) (string, string) {
	startTime := time.Now()

	var baseURL string = "https://www.acmicpc.net"
	var problemURL string = "/problem/tags"
	var userURL string = "/user/"
	var userID string = user
	var newRand int

	resultList := getPage(baseURL, problemURL)

	solvedList := getUserSolvedProblemInfo(baseURL, userURL, userID)

	for true {
		newRand = getRandomIndex(len(resultList))
		if resultList[newRand].problemNumber == 0 {
			continue
		}
		if !contains(solvedList, resultList[newRand].problemNumber) {
			break
		}
	}
	fmt.Println(baseURL + "/problem/" + strconv.Itoa(resultList[newRand].problemNumber))

	elapsedTime := time.Now().Sub(startTime)
	fmt.Println("Time elapsed: ", elapsedTime)

	return resultList[newRand].title, baseURL + "/problem/" + strconv.Itoa(resultList[newRand].problemNumber)
}

func getPage(url, subURL string) []extractedProblemLinks {
	var resultProblems []extractedProblemLinks
	var targetURL string = url + subURL
	c := make(chan []extractedProblemLinks)
	numberOfTags := 0

	res, err := http.Get(targetURL)
	checkErr(err)
	checkCode(res)

	defer res.Body.Close()

	doc, err := goquery.NewDocumentFromReader(res.Body)
	checkErr(err)

	searchCards := doc.Find("tr")

	searchCards.Each(func(i int, card *goquery.Selection) {
		if i == 0 {
			return
		}
		var address string
		var tagTitle string
		var numberOfProblems int
		a := card.Find("a")
		href, ok := a.Attr("href")
		if ok {
			address = url + href
		}
		tdList := card.Find("td")
		tdList.Each(func(j int, td *goquery.Selection) {
			if j == 0 {
				tagTitle = td.Text()
			} else {
				numberOfProblems = atoi(td.Text())
			}
		})
		if numberOfProblems < 5 {
			return
		}
		numberOfTags++
		go extraactSeperateProblems(address, tagTitle, c)
	})

	for i := 0; i < numberOfTags; i++ {
		problems := <-c
		resultProblems = append(resultProblems, problems...)
	}

	return resultProblems
}

func extraactSeperateProblems(url, title string, c chan<- []extractedProblemLinks) {
	var resultList []extractedProblemLinks

	res, err := http.Get(url)
	checkErr(err)
	checkCode(res)

	defer res.Body.Close()

	doc, err := goquery.NewDocumentFromReader(res.Body)
	checkErr(err)

	trList := doc.Find("tr")

	trList.Each(func(i int, tr *goquery.Selection) {
		var problem extractedProblemLinks

		problem.tagTitle = title
		problem.problemNumber = atoi(tr.Find(".list_problem_id").Text())
		problem.rank = tr.Find(".level_hidden").Text()

		tr.Find("td").Each(func(j int, td *goquery.Selection) {
			if j != 1 {
				return
			}
			problem.title = td.Find("a").Text()
		})

		resultList = append(resultList, problem)
	})

	c <- resultList
}

func getUserSolvedProblemInfo(baseURL, userURL, userID string) []int {
	targetURL := baseURL + userURL + userID
	var resultList []int

	res, err := http.Get(targetURL)
	checkErr(err)
	checkCode(res)

	defer res.Body.Close()

	doc, err := goquery.NewDocumentFromReader(res.Body)
	checkErr(err)

	doc.Find(".panel-body").First().Find(".problem_number").Each(func(i int, span *goquery.Selection) {
		resultList = append(resultList, atoi(span.Text()))
	})

	return resultList
}

func checkErr(err error) {
	if err != nil {
		log.Fatalln(err)
	}
}

func checkCode(res *http.Response) {
	if res.StatusCode != 200 {
		log.Fatalf("status code error: %d %s", res.StatusCode, res.Status)
	}
}

// CleanString cleans the given string
func CleanString(str string) string {
	return strings.Join(strings.Fields(strings.TrimSpace(str)), " ")
}

func atoi(str string) int {
	if len(str) == 0 {
		return 0
	}
	n, err := strconv.Atoi(str)
	checkErr(err)
	return n
}

func contains(arr []int, number int) bool {
	for _, a := range arr {
		if a == number {
			return true
		}
	}
	return false
}

func getRandomIndex(number int) int {
	n, err := rand.Int(rand.Reader, big.NewInt(int64(number)))
	checkErr(err)
	return int(n.Int64())
}
