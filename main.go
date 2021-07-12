package main

import (
	"errors"
	"fmt"
	"net/http"
)

//
var errRequestedFailed = errors.New("Requested Failed")

func main() {

	var results = make(map[string]string)

	urls := []string{
		"https://www.airbnb.com/",
		"https://www.google.com/",
		"https://facebook.com/",
	}

	for _, url := range urls {
		result := "OK"
		err := hitURL(url)

		if err != nil {
			result = "FAILED"
		}
		results[url] = result
	}

	for url, result := range results {
		fmt.Println(url, result)

	}

	// fmt.Println(results)
}

func hitURL(url string) error {
	fmt.Println("checking: ", url)
	resp, err := http.Get(url)

	if err != nil || resp.StatusCode >= 400 {
		return errRequestedFailed
	}
	return nil

}
