package util

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"os"
)

var token = "eyJ0eXAiOiJKV1QiLCJhbGciOiJSUzI1NiIsImp0aSI6Ijg3NzEyYjUzMTY2NjhhODdkNzM2NDU3ZTdhMjFhNDFkNTYyMzViMWE2ZWI2M2ZkYzY1NTlkNDM4ZTU3MmEwNGQyYzU0M2U3YWYzZTIwNzYwIn0.eyJhdWQiOiIxODIyMjQ5MzguMjNhNzJmNDA2NzI4M2I0OWY5NjZmOTMyMzViMTg2NDQzN2VjNWY2YTlmY2M5NjVlOGIzOTM5MGRmNWQ2YWE5NCIsImp0aSI6Ijg3NzEyYjUzMTY2NjhhODdkNzM2NDU3ZTdhMjFhNDFkNTYyMzViMWE2ZWI2M2ZkYzY1NTlkNDM4ZTU3MmEwNGQyYzU0M2U3YWYzZTIwNzYwIiwiaWF0IjoxNTYyMTU0NzIyLCJuYmYiOjE1NjIxNTQ3MjIsImV4cCI6MTU3NzcwNjcyMiwic3ViIjoiOTg1ODQ5ODYxNjA5MTYwNzA0Iiwic2NvcGVzIjpbInJlYWQiXX0.TdADyi_whM-FpH4hTBjvcCME17YtUftMm-jUZLzRo2G1n2qoANcrqNw2geXZvii6P93NWnBalAN2mWuG2CfRdQj3jix89Bnkxkxb5usNkRPrMtq8I_gMDcBYmz2xH9jUGiIq9z622zFcV9Z8UD-t6LB0WraYmg_QsGCU3PKatgt-M26u86de-KW-eGfY9a_pE-o63ggSb_lN70jBDgzuu7tohSviads-9aFwwoiw7Suu2ojz1X6NhMz5lfFcV8EDpXhh7v7H1IRX24oQqTLjkXoKobSnD7dRj2kFOorqfyQtnaIrlJKu1JnT3tz-8MAn2pZi0a6jf8sbIwLkEH9wpA"

type QuizError struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}

type Quiz struct {
	ID       string    `json:"id"`
	Question string    `json:"question"`
	Error    QuizError `json:"error"`
}
type Answer struct {
	Score   string `json:"score"`
	Message string `json:"message"`
	Hints   string `json:"hints"`
	Round   int    `json:"round"`
}

// 6 ? 9 ? 8 ? 7 ? 1 ? 6 ? 4 = 20
// 6 ? 9 ? 8 ? 7 ? 1 = 20
func ParseQuestion(question string) (string, string) {
	equalPosition := equalPosition(question)
	// fmt.Printf("len: %+v\n", len(question))
	// fmt.Printf("ParseQuestion: %+v\n", equalPosition)
	ques := question[:equalPosition-1]
	answer := question[equalPosition+2:]
	// fmt.Printf("ques: %+v\n", ques)
	// fmt.Printf("answer: %+v\n", answer)
	return ques, answer
}

func equalPosition(question string) int {
	for i := 0; i < len(question); i++ {
		if question[i] == '=' {
			return i
		}
	}
	return 0
}

func CreateGame(level string) Quiz {
	accessToken := "Bearer " + token
	url := "https://apiv2.twitcasting.tv/internships/2019/games?level=" + level

	client := &http.Client{}
	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Set("Authorization", accessToken)
	res, _ := client.Do(req)
	body, _ := ioutil.ReadAll(res.Body)
	defer res.Body.Close()

	var quiz Quiz
	json.Unmarshal(body, &quiz)

	if len(quiz.ID) < 1 {
		status := deleteGame()
		if !status {
			os.Exit(3)
		}
		quiz = CreateGame(level)
	}

	return quiz
}

func postAnswear(gameID, answer string) Answer {
	accessToken := "Bearer " + token
	url := "https://apiv2.twitcasting.tv/internships/2019/games/" + gameID
	client := &http.Client{}
	jsonStr := `{ "answer": "` + answer + `"}`
	req, _ := http.NewRequest(
		"POST",
		url,
		bytes.NewBuffer([]byte(jsonStr)),
	)
	req.Header.Set("Authorization", accessToken)
	req.Header.Set("Content-Type", "application/json")
	res, _ := client.Do(req)
	body, _ := ioutil.ReadAll(res.Body)
	defer res.Body.Close()

	var ans Answer
	json.Unmarshal(body, &ans)

	if ans.Round == 1 {
		_ = deleteGame()
	}

	return ans
}

func deleteGame() bool {
	accessToken := "Bearer " + token
	url := "https://apiv2.twitcasting.tv/internships/2019/games"

	client := &http.Client{}
	req, _ := http.NewRequest("DELETE", url, nil)
	req.Header.Set("Authorization", accessToken)
	res, _ := client.Do(req)
	// body, _ := ioutil.ReadAll(res.Body)
	// defer res.Body.Close()

	if res.StatusCode == 200 {
		return true
	}
	return false
}
