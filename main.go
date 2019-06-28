package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

type Quiz struct {
	id       string
	question string
}
type Answer struct {
	score   string
	message string

	ints  string
	round int
}

func main() {
	access_token := "Bearer "
	access_token += "eyJ0eXAiOiJKV1QiLCJhbGciOiJSUzI1NiIsImp0aSI6IjgwMDNiOGJkMWU5ZmEyYTdjNjU5MDYxYTUwOTQ1ODQzNDU0YTJiYzMyZWU2YTI0OTcxNzg2ODRlNzk4NWE5NDU0MTUzNmI4MWZjNzExNGNmIn0.eyJhdWQiOiIxODIyMjQ5MzguMjNhNzJmNDA2NzI4M2I0OWY5NjZmOTMyMzViMTg2NDQzN2VjNWY2YTlmY2M5NjVlOGIzOTM5MGRmNWQ2YWE5NCIsImp0aSI6IjgwMDNiOGJkMWU5ZmEyYTdjNjU5MDYxYTUwOTQ1ODQzNDU0YTJiYzMyZWU2YTI0OTcxNzg2ODRlNzk4NWE5NDU0MTUzNmI4MWZjNzExNGNmIiwiaWF0IjoxNTYxNjE3OTMwLCJuYmYiOjE1NjE2MTc5MzAsImV4cCI6MTU3NzE2OTkzMCwic3ViIjoiOTg1ODQ5ODYxNjA5MTYwNzA0Iiwic2NvcGVzIjpbInJlYWQiXX0.HlyPPVNRzp9Zpz7hCPe0b1NAKAgk-T45819hMJYhHdXky72xbYCs373-wPH5hbAtuNIwc11h4yudQvxEuWL6cg-Pi1xHj38zkAseE_gK5wmnUWxpdvwZ-I5dhFaXuPmzVu4YGQxmVPMN07R7sZ6HY779ay90FpcErwjcISQNJtQKPMT3WRdvriSpCDJNpCCLRrKmXU7Ss1NL-_Nj_27ex8U9nwX6fldzSLnUX6auCRIirUie-mcjLMDiFqiGKBh7TSO7Ul5Hhv6uD8qqkXTGBk6M2Dkqe3sn1aI3FiEMQ3kbt6kSf4Pu42MuYgwdSDv7CCb8J56PXCZ7whDsCiGOxw"
	url := "https://apiv2.twitcasting.tv/internships/2019/games?level=2"

	h := http.Header{}
	h.Set("Authorization", access_token)

	response, err := http.Get(url)

	if err != nil {
		fmt.Print(err.Error())
		os.Exit(1)
	}

	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(responseData))
}
