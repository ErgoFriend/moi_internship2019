package main

import (
	"fmt"
)

func main() {
	// quiz := Quiz{ID: "fd94a2ada12167b076a82ed2da97a8ad", Question: "7 ? 7 ? 8 ? 9 ? 5 = -53"}

	quiz := createGame()
	strQuiz, ans := ParseQuestion(quiz.Question)

	fmt.Printf("%+v\n", quiz)
	fmt.Printf("%+v\n", strQuiz)
	fmt.Printf("%+v\n", ans)

	result := postAnswear(quiz.ID, "*-*+")
	fmt.Printf("%+v\n", result)

	_ = deleteGame()
}
