package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"sync"
)

var level = 3

// var size = int(math.Pow(4, float64(level)))

func main() {
	// quiz := Quiz{ID: "fd94a2ada12167b076a82ed2da97a8ad", Question: "7 ? 7 ? 8 ? 9 ? 5 = -53"}
	// quiz := Quiz{ID: "fd94a2ada12167b076a82ed2da97a8ad", Question: "2 ? 7 ? 4 ? 3 ? 6 ? 1 ? 2 = 55"}

	quiz := createGame(strconv.Itoa(level))
	fmt.Printf("Quiz: %+v\n", quiz)
	if len(quiz.ID) < 1 {
		_ = deleteGame()
	}
	strQuiz, answer := ParseQuestion(quiz.Question)

	// fmt.Printf("%+v\n", quiz)
	// fmt.Printf("After: %+v=%+v\n", strQuiz, ans)

	// mapChan := make(chan map[string]string, 1)
	// Search(&answerFormula, strQuiz, answer)
	formulas := Forms(strQuiz)
	// for _, v := range formulas {
	// 	fmt.Printf("%+v\n", v)
	// }
	wg := &sync.WaitGroup{} // WaitGroupの値を作る

	var notAnswered = true

	for _, formula := range formulas {
		answer := answer
		wg.Add(1) // wgをインクリメント
		if notAnswered {
			go func(formula string, answer string) {
				result := Eval(formula, answer)
				// fmt.Printf("Eval: %s = %s\n", formula, result)
				// rr, _ := strconv.Atoi(result)
				// aa, _ := strconv.Atoi(answer)
				if result == answer {
					// if rr == aa {
					rep := regexp.MustCompile("1|2|3|4|5|6|7|8|9| ")
					symStr := rep.ReplaceAllString(formula, "")
					// fmt.Printf("Quiz: %+v\n", quiz)
					// fmt.Printf("Eval: %s = %s\n", formula, result)

					resultAnswer := postAnswear(quiz.ID, symStr)
					fmt.Printf("%+v\n", resultAnswer)

					notAnswered = false
				}
				wg.Done()
			}(formula, answer)
		}
	}

	wg.Wait()
	// mapChan <- formulasMap
	// value := <-mapChan
	// fmt.Printf("%s\n", value)

	// fmt.Printf("%+v\n", resultAnswer)

	bufio.NewScanner(os.Stdin).Scan()

}
