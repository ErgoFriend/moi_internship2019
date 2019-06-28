package main

import (
	"fmt"
)

func main() {
	var (
		strQuiz string
		ans     int
	)
	quiz := Quiz{ID: "fd94a2ada12167b076a82ed2da97a8ad", Question: "7 ? 7 ? 8 ? 9 ? 5 = -53"}

	// quiz = createGame()
	strQuiz, ans = parseQuestion(quiz.Question)

	fmt.Printf("%+v\n", quiz)
}
