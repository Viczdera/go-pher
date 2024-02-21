package main

import (
	"fmt"
	"io"
	"net/http"
	"strings"
)

func main() {
	fmt.Println("Web stuff in golang")
	const baseUrl string = "http://localhost:8000"
	//performGetReq(route)
	performPostReq(baseUrl + "/post")
}

func performGetReq(url string) {
	res, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	defer res.Body.Close()

	//good res
	resString := parseResString(res.Body)
	fmt.Println("Response: ", resString)

}

func performPostReq(url string) {
	//dummy json
	jsonD := strings.NewReader(`{
		"item":"Getting started with flutter",
		"price":"8.99",
		"currency":"usd",
		"quantity":"1"
	}`)

	res, err := http.Post(url, "application/json", jsonD)
	if err != nil {
		panic(err)
	}
	defer res.Body.Close()

	//good res
	resString := parseResString(res.Body)
	fmt.Println("Response: ", resString)

}
func parseResString(resBody io.Reader) string {
	var resString strings.Builder
	var parsedResString string

	content, _ := io.ReadAll(resBody)
	resString.Write(content)
	parsedResString = resString.String()
	// return string(content)
	return parsedResString
}
