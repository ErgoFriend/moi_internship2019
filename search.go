package main

import (
	"fmt"
	"strings"
	"sync"

	"github.com/novalagung/golpal"
)

var symbols = []string{"+", "-", "*", "/"}

func Eval(formula, answer string) (bool, string) {
	output, err := golpal.New().ExecuteSimple(formula)
	fmt.Printf("Eval: %s = %s\n", formula, output)
	if err != nil {
		return false, ""
	}
	// fmt.Printf("%s=%s\n", output, answer)
	if output == answer {
		return true, formula
	}
	return false, ""
}

func Eval2(formula, answer string) {
	output, _ := golpal.New().ExecuteSimple(formula)
	fmt.Printf("Eval: %s = %s\n", formula, output)
}

func Search(strQuiz, answer string) (bool, string) {
	var wg sync.WaitGroup
	for _, v := range symbols {
		v := v
		replaced1 := strings.Replace(strQuiz, "?", v, 1)
		wg.Add(1)
		go func() {
			for _, j := range symbols {
				j := j
				replaced2 := strings.Replace(replaced1, "?", j, 1)
				wg.Add(1)
				go func() {
					for _, i := range symbols {
						i := i
						replaced3 := strings.Replace(replaced2, "?", i, 1)
						wg.Add(1)
						go func() {
							for _, k := range symbols {
								k := k
								replaced4 := strings.Replace(replaced3, "?", k, 1)
								// _, _ = Eval(replaced4, answer)
								defer wg.Done()
								Eval2(replaced4, answer)
								// if status {
								// 	return status, result
								// }
							}
						}()
					}
				}()
			}
		}()
	}

	wg.Wait()
	return false, ""
}
