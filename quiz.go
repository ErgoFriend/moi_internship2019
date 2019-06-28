package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"
)

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
	score   string
	message string

	hints string
	round int
}

func parseQuestion(question string) (string, int) {

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
	accessToken := "Bearer eyJ0eXAiOiJKV1QiLCJhbGciOiJSUzI1NiIsImp0aSI6IjgwMDNiOGJkMWU5ZmEyYTdjNjU5MDYxYTUwOTQ1ODQzNDU0YTJiYzMyZWU2YTI0OTcxNzg2ODRlNzk4NWE5NDU0MTUzNmI4MWZjNzExNGNmIn0.eyJhdWQiOiIxODIyMjQ5MzguMjNhNzJmNDA2NzI4M2I0OWY5NjZmOTMyMzViMTg2NDQzN2VjNWY2YTlmY2M5NjVlOGIzOTM5MGRmNWQ2YWE5NCIsImp0aSI6IjgwMDNiOGJkMWU5ZmEyYTdjNjU5MDYxYTUwOTQ1ODQzNDU0YTJiYzMyZWU2YTI0OTcxNzg2ODRlNzk4NWE5NDU0MTUzNmI4MWZjNzExNGNmIiwiaWF0IjoxNTYxNjE3OTMwLCJuYmYiOjE1NjE2MTc5MzAsImV4cCI6MTU3NzE2OTkzMCwic3ViIjoiOTg1ODQ5ODYxNjA5MTYwNzA0Iiwic2NvcGVzIjpbInJlYWQiXX0.HlyPPVNRzp9Zpz7hCPe0b1NAKAgk-T45819hMJYhHdXky72xbYCs373-wPH5hbAtuNIwc11h4yudQvxEuWL6cg-Pi1xHj38zkAseE_gK5wmnUWxpdvwZ-I5dhFaXuPmzVu4YGQxmVPMN07R7sZ6HY779ay90FpcErwjcISQNJtQKPMT3WRdvriSpCDJNpCCLRrKmXU7Ss1NL-_Nj_27ex8U9nwX6fldzSLnUX6auCRIirUie-mcjLMDiFqiGKBh7TSO7Ul5Hhv6uD8qqkXTGBk6M2Dkqe3sn1aI3FiEMQ3kbt6kSf4Pu42MuYgwdSDv7CCb8J56PXCZ7whDsCiGOxw"
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

func deleteGame() bool {
	accessToken := "Bearer eyJ0eXAiOiJKV1QiLCJhbGciOiJSUzI1NiIsImp0aSI6IjgwMDNiOGJkMWU5ZmEyYTdjNjU5MDYxYTUwOTQ1ODQzNDU0YTJiYzMyZWU2YTI0OTcxNzg2ODRlNzk4NWE5NDU0MTUzNmI4MWZjNzExNGNmIn0.eyJhdWQiOiIxODIyMjQ5MzguMjNhNzJmNDA2NzI4M2I0OWY5NjZmOTMyMzViMTg2NDQzN2VjNWY2YTlmY2M5NjVlOGIzOTM5MGRmNWQ2YWE5NCIsImp0aSI6IjgwMDNiOGJkMWU5ZmEyYTdjNjU5MDYxYTUwOTQ1ODQzNDU0YTJiYzMyZWU2YTI0OTcxNzg2ODRlNzk4NWE5NDU0MTUzNmI4MWZjNzExNGNmIiwiaWF0IjoxNTYxNjE3OTMwLCJuYmYiOjE1NjE2MTc5MzAsImV4cCI6MTU3NzE2OTkzMCwic3ViIjoiOTg1ODQ5ODYxNjA5MTYwNzA0Iiwic2NvcGVzIjpbInJlYWQiXX0.HlyPPVNRzp9Zpz7hCPe0b1NAKAgk-T45819hMJYhHdXky72xbYCs373-wPH5hbAtuNIwc11h4yudQvxEuWL6cg-Pi1xHj38zkAseE_gK5wmnUWxpdvwZ-I5dhFaXuPmzVu4YGQxmVPMN07R7sZ6HY779ay90FpcErwjcISQNJtQKPMT3WRdvriSpCDJNpCCLRrKmXU7Ss1NL-_Nj_27ex8U9nwX6fldzSLnUX6auCRIirUie-mcjLMDiFqiGKBh7TSO7Ul5Hhv6uD8qqkXTGBk6M2Dkqe3sn1aI3FiEMQ3kbt6kSf4Pu42MuYgwdSDv7CCb8J56PXCZ7whDsCiGOxw"
	url := "	https://apiv2.twitcasting.tv/internships/2019/games/"

	client := &http.Client{}
	req, _ := http.NewRequest("DELETE", url, nil)
	req.Header.Set("Authorization", accessToken)
	res, _ := client.Do(req)
	defer res.Body.Close()

	if res.StatusCode == 200 {
		return true
	}
	return false
}
