package main

import (
	"fmt"
	"os"

	"github.com/ikngtty/answer-bot-for-discord/pkg/chiebukuro"
)

func main() {
	question, err := chiebukuro.Sample()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println(question.BestAnswer)
	fmt.Println(question.URL)
}
