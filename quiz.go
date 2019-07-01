package main

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"
)

var token string = "eyJ0eXAiOiJKV1QiLCJhbGciOiJSUzI1NiIsImp0aSI6IjYxNDEyNWIzMTA3Nzk5NjMyNDcwOTRkMTAwZWFiODgzNGNmNTNhMTJmZmZjNTg4ZWE1MGVjMzQ1MDMzZDEyNDJmM2VkOThjMmMyMTkyZjcwIn0.eyJhdWQiOiIxODIyMjQ5MzguMjNhNzJmNDA2NzI4M2I0OWY5NjZmOTMyMzViMTg2NDQzN2VjNWY2YTlmY2M5NjVlOGIzOTM5MGRmNWQ2YWE5NCIsImp0aSI6IjYxNDEyNWIzMTA3Nzk5NjMyNDcwOTRkMTAwZWFiODgzNGNmNTNhMTJmZmZjNTg4ZWE1MGVjMzQ1MDMzZDEyNDJmM2VkOThjMmMyMTkyZjcwIiwiaWF0IjoxNTYxOTY2MTcyLCJuYmYiOjE1NjE5NjYxNzIsImV4cCI6MTU3NzUxODE3Miwic3ViIjoiOTg1ODQ5ODYxNjA5MTYwNzA0Iiwic2NvcGVzIjpbInJlYWQiXX0.DiPD3wKAtQBp9LSR-bRQ7jbiyU3DllruQXmjpila25JevjzCabL4vpHdlbMtmsdJEXZu3iBYyjL-6mbXCBW4GHHzawvnG5P8Vixogao2E9vbzUiZrauvBn1ysRiDNHenVykeSPls8BpnMIMlxysJ7B121vU1fJ3H665icP7I2FmJinybC-KWtcz-msIFw2d5TIN1pf3xxTa-UfzTXIaLf7vzjzIGnJnRFM35RjFwu2dx-uYf8x4GEQDtrb1hTqHf9zzWv68cw7uE0paLlCfeyo-oKuEXJqBdglh06eCzfgfCj0aaB7V5LGDV93-mTNlzp1RudMYnW-m6ORTUeJ5t0Q"

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

func ParseQuestion(question string) (string, int) {

	equalPosition := equalPosition(question)
	ques := question[:equalPosition]
	answer := question[equalPosition:]

	ans, _ := strconv.Atoi(answer)
	return ques, ans
}

func equalPosition(question string) int {
	for i := 0; i < len(question); i++ {
		if question[i] == '=' {
			return i
		}
	}
	return 0
}

func createGame() Quiz {
	accessToken := "Bearer " + token
	url := "https://apiv2.twitcasting.tv/internships/2019/games?level=2"

	client := &http.Client{}
	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Set("Authorization", accessToken)
	res, _ := client.Do(req)
	body, _ := ioutil.ReadAll(res.Body)
	defer res.Body.Close()

	var quiz Quiz
	json.Unmarshal(body, &quiz)

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
	url := "https://apiv2.twitcasting.tv/internships/2019/games/"

	client := &http.Client{}
	req, _ := http.NewRequest("DELETE", url, nil)
	req.Header.Set("Authorization", accessToken)
	res, _ := client.Do(req)
	_, _ = ioutil.ReadAll(res.Body)
	defer res.Body.Close()

	if res.StatusCode == 200 {
		return true
	}
	return false
}
