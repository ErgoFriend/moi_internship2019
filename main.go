package main

import (
	"fmt"
)

func main() {
	// quiz := Quiz{ID: "fd94a2ada12167b076a82ed2da97a8ad", Question: "7 ? 7 ? 8 ? 9 ? 5 = -53"}

	quiz := createGame()
	strQuiz, ans := ParseQuestion(quiz.Question)

	fmt.Printf("%+v\n", quiz)
	// fmt.Printf("%+v = %+v\n", strQuiz, ans)

	_, formula := Search(strQuiz, ans)
	fmt.Printf("%+v\n", formula)

	// rep := regexp.MustCompile("1|2|3|4|5|6|7|8|9| ")
	// symStr := rep.ReplaceAllString(formula, "")
	// fmt.Printf("%+v\n", a)

	// result := postAnswear(quiz.ID, symStr)
	// fmt.Printf("%+v\n", result)

	// _ = deleteGame()
}
