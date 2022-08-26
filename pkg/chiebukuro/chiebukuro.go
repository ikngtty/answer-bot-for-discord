package chiebukuro

import (
	"fmt"
	"math/rand"
	"net/http"
	"time"

	"github.com/PuerkitoBio/goquery"
)

type Question struct {
	BestAnswer string
	URL        string
}

func Sample() (question *Question, err error) {
	questionURL, err := sampleQuestionURL()
	if err != nil {
		return
	}
	question, err = fetchQuestion(questionURL)
	return
}

func sampleQuestionURL() (questionURL string, err error) {
	rand.Seed(time.Now().UnixNano())

	page := rand.Intn(100) + 1
	listURL := fmt.Sprintf("https://chiebukuro.yahoo.co.jp/question/list?flg=1&page=%d", page)
	resp, err := http.Get(listURL)
	if err != nil {
		return
	}
	defer resp.Body.Close()

	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		return
	}

	questionIndex := rand.Intn(40)
	questionNode := doc.Find("#qa_lst").Children().Get(questionIndex)
	questionURL, _ = goquery.NewDocumentFromNode(questionNode).Find("a").Attr("href")
	return
}

func fetchQuestion(url string) (question *Question, err error) {
	resp, err := http.Get(url)
	if err != nil {
		return
	}
	defer resp.Body.Close()

	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		return
	}

	bestAnswerNode := doc.Find("#ba").Find("h2").Get(1)
	bestAnswer := goquery.NewDocumentFromNode(bestAnswerNode).Text()
	question = &Question{BestAnswer: bestAnswer, URL: url}
	return
}
