package main

import (
	"fmt"
	"strconv"

	"./util"
)

func main() {
	// quiz := Quiz{ID: "fd94a2ada12167b076a82ed2da97a8ad", Question: "7 ? 7 ? 8 ? 9 ? 5 = -53"}
	// quiz := Quiz{ID: "fd94a2ada12167b076a82ed2da97a8ad", Question: "2 ? 7 ? 4 ? 3 ? 6 ? 1 ? 2 = 55"}

	var (
		level = 3
		quiz  util.Quiz
	)
	quiz = util.CreateGame(strconv.Itoa(level))

	fmt.Printf("Quiz: %+v\n", quiz)
	strQuiz, answer := util.ParseQuestion(quiz.Question)

	util.EvalFormulas(strQuiz, answer, quiz.ID, level)

	// bufio.NewScanner(os.Stdin).Scan()

}
