package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"sync"
)

func main() {
	// quiz := Quiz{ID: "fd94a2ada12167b076a82ed2da97a8ad", Question: "7 ? 7 ? 8 ? 9 ? 5 = -53"}
	// quiz := Quiz{ID: "fd94a2ada12167b076a82ed2da97a8ad", Question: "2 ? 7 ? 4 ? 3 ? 6 ? 1 ? 2 = 55"}

	var (
		level = 3
		quiz  Quiz
	)
	quiz = createGame(strconv.Itoa(level))

	if len(quiz.ID) < 1 {
		status := deleteGame()
		if !status {
			os.Exit(3)
		}
		quiz = createGame(strconv.Itoa(level))
	}
	fmt.Printf("Quiz: %+v\n", quiz)
	strQuiz, answer := ParseQuestion(quiz.Question)
	fmt.Printf("%v = %+v\n", strQuiz, answer)

	// wg := &sync.WaitGroup{} // WaitGroupの値を作る
	// wg.Add(1)
	// go EvalFormulas(strQuiz, answer, quiz.ID, level)
	// wg.Done()
	// wg.Wait()

	wg := &sync.WaitGroup{} // WaitGroupの値を作る
	for _, v := range symbols {
		fmt.Printf("v\n")
		go func(replaced1 string) {
			for _, j := range symbols {
				fmt.Printf("j\n")
				go func(replaced2 string) {
					for _, i := range symbols {
						go func(replaced3 string) {
							for _, k := range symbols {
								go func(replaced4 string) {
									for _, l := range symbols {
										go func(replaced5 string) {
											for _, m := range symbols {
												wg.Add(1) // wgをインクリメント
												go func(formula string) {
													Eval(formula, answer, quiz.ID)
													// fmt.Printf("formula: %+v\n", formula)
													wg.Done()
												}(strings.Replace(replaced5, "?", m, 1))
											}
										}(strings.Replace(replaced4, "?", l, 1))
									}
								}(strings.Replace(replaced3, "?", k, 1))
							}
						}(strings.Replace(replaced2, "?", i, 1))
					}
				}(strings.Replace(replaced1, "?", j, 1))
			}
		}(strings.Replace(strQuiz, "?", v, 1))
	}

	wg.Wait()

	// fmt.Printf("%+v\n", resultAnswer)

}
