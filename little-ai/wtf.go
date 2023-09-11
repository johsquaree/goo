package main

import (
	"fmt"
	"math/rand"
	"time"
)

var responses = []string{
	"pull yourself together.",
	"wtf.",
	"maybe.",
	"ask again later.",
	"without a doubt.",
	"very doubtful.",
	"i know but u cant understand.",
	"r u killer?",
	"im sure is that you are not going to be able to do that right now so shut up.",
	"u r very smart?",
	"im good enough.",

}

func getRandomResponse() string {
	rand.Seed(time.Now().UnixNano())
	randomIndex := rand.Intn(len(responses))
	return responses[randomIndex]
}

func main() {
	fmt.Println("ask me anything. if u want exit in me u should write 'exit'")

	for {
		fmt.Print("ask: ")
		var question string
		fmt.Scanln(&question)

		if question == "exit" {
			fmt.Println("see u later.")
			break
		}

		response := getRandomResponse()
		fmt.Printf("answer: %s\n", response)
	}
}
