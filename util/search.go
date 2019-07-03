package util

import (
	"fmt"
	"regexp"
	"strings"
	"sync"

	"github.com/Knetic/govaluate"
)

var symbols = []string{"+", "-", "*", "/"}

func Eval(formula, answer, quizID string) {
	// fmt.Printf("%+v\n", formula)
	expression, _ := govaluate.NewEvaluableExpression(formula)
	result, _ := expression.Evaluate(nil)
	resultStr := fmt.Sprintf("%v", result)
	if resultStr == answer {
		rep := regexp.MustCompile("1|2|3|4|5|6|7|8|9| ")
		symStr := rep.ReplaceAllString(formula, "")
		resultAnswer := postAnswear(quizID, symStr)
		fmt.Printf("%+v\n", resultAnswer)
	}
}

func EvalFormulas(strQuiz, answer, quizID string, level int) {
	wg := &sync.WaitGroup{} // WaitGroupの値を作る
	fmt.Printf("EvalFormulas\n")
	if level == 2 {
		for _, v := range symbols {
			wg.Add(1) // wgをインクリメント
			go func(replaced1 string) {
				for _, j := range symbols {
					wg.Add(1) // wgをインクリメント
					go func(replaced2 string) {
						for _, i := range symbols {
							wg.Add(1) // wgをインクリメント
							go func(replaced3 string) {
								for _, k := range symbols {
									wg.Add(1) // wgをインクリメント
									go func(formula string) {
										Eval(formula, answer, quizID)
										// fmt.Printf("formula: %+v\n", formula)
										wg.Done()
									}(strings.Replace(replaced3, "?", k, 1))
								}
								wg.Done()
							}(strings.Replace(replaced2, "?", i, 1))
						}
						wg.Done()
					}(strings.Replace(replaced1, "?", j, 1))
				}
				wg.Done()
			}(strings.Replace(strQuiz, "?", v, 1))
		}
	} else if level == 3 {
		for _, v := range symbols {
			wg.Add(1)
			go func(replaced1 string) {
				for _, j := range symbols {
					wg.Add(1)
					go func(replaced2 string) {
						for _, i := range symbols {
							wg.Add(1)
							go func(replaced3 string) {
								for _, k := range symbols {
									wg.Add(1)
									go func(replaced4 string) {
										for _, l := range symbols {
											wg.Add(1)
											go func(replaced5 string) {
												for _, m := range symbols {
													wg.Add(1) // wgをインクリメント
													go func(formula string) {
														Eval(formula, answer, quizID)
														wg.Done()
													}(strings.Replace(replaced5, "?", m, 1))
												}
												wg.Done()
											}(strings.Replace(replaced4, "?", l, 1))
										}
										wg.Done()
									}(strings.Replace(replaced3, "?", k, 1))
								}
								wg.Done()
							}(strings.Replace(replaced2, "?", i, 1))
						}
						wg.Done()
					}(strings.Replace(replaced1, "?", j, 1))
				}
				wg.Done()
			}(strings.Replace(strQuiz, "?", v, 1))
		}
	}
	wg.Wait()
	fmt.Printf("EvalFormulas Done.\n")
}
