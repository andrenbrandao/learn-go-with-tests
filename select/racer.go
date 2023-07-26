package main

import (
	"net/http"
)

func Racer(url1, url2 string) string {
	ch := make(chan string)

	getUrl := func(url string) {
		_, err := http.Get(url)

		if err == nil {
			ch <- url
		}
	}

	go getUrl(url1)
	go getUrl(url2)

	res := <-ch
	return res
}
