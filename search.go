package main

import (
	"fmt"
	"strings"
	"sync"

	"github.com/Knetic/govaluate"
)

var symbols = []string{"+", "-", "*", "/"}

func Eval(formula, answer string) string {
	expression, _ := govaluate.NewEvaluableExpression(formula)
	result, _ := expression.Evaluate(nil)

	str := fmt.Sprintf("%v", result)
	// fmt.Printf("Eval: %s = %v\n", formula, str)
	return str
}

func Forms(strQuiz string) [4 * 4 * 4 * 4 * 4 * 4]string {
	var formulas [4 * 4 * 4 * 4 * 4 * 4]string

	if len(strQuiz) == 17 {
		loc := 0
		for _, v := range symbols {
			replaced1 := strings.Replace(strQuiz, "?", v, 1)
			for _, j := range symbols {
				replaced2 := strings.Replace(replaced1, "?", j, 1)
				for _, i := range symbols {
					replaced3 := strings.Replace(replaced2, "?", i, 1)
					for _, k := range symbols {
						formulas[loc] = strings.Replace(replaced3, "?", k, 1)
						loc++
					}
				}
			}
		}
	} else if len(strQuiz) == 25 {
		loc := 0
		for _, v := range symbols {
			replaced1 := strings.Replace(strQuiz, "?", v, 1)
			for _, j := range symbols {
				replaced2 := strings.Replace(replaced1, "?", j, 1)
				for _, i := range symbols {
					replaced3 := strings.Replace(replaced2, "?", i, 1)
					for _, k := range symbols {
						replaced4 := strings.Replace(replaced3, "?", k, 1)
						for _, l := range symbols {
							replaced5 := strings.Replace(replaced4, "?", l, 1)
							for _, m := range symbols {
								formulas[loc] = strings.Replace(replaced5, "?", m, 1)
								loc++
							}
						}
					}
				}
			}
		}
	}

	return formulas
}

func Search(answerFormula *string, strQuiz, answer string) {
	formulas := Forms(strQuiz)
	// for _, v := range formulas {
	// 	fmt.Printf("%+v\n", v)
	// }
	wg := &sync.WaitGroup{} // WaitGroupの値を作る

	for _, formula := range formulas {
		wg.Add(1) // wgをインクリメント
		go func(formula string) {
			result := Eval(formula, answer)
			// fmt.Printf("Eval: %s = %s\n", formula, result)
			if result == answer {
				*answerFormula = formula
			}
			wg.Done()
		}(formula)
	}

	wg.Wait()

}
