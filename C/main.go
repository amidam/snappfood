package main

import (
	"fmt"
	"net/http"
	"net/http/httputil"
	"sync"
)

func main() {
	urls := []string{"https://rogerwelin.github.io/",
		"https://snappfood.ir/",
		"https://www.isna.ir/",
		"https://www.asriran.com/",
		"https://www.google.com/"}

	m := fetch(urls)
	for k, v := range m {
		fmt.Printf("%s -> %s\n", k, v)
	}
}

func fetch(urls []string) map[string]string {
	var results = make(map[string]string)
	var wg sync.WaitGroup

	for _, url := range urls {
		wg.Add(1)
		go func(url string) error {
			defer wg.Done()
			resp, err := http.Get(url)
			if err != nil {
				return err
			}
			dump, err := httputil.DumpResponse(resp, true)
			if err != nil {
				return err
			}
			results[url] = string(dump)
			return nil
		}(url)
	}
	wg.Wait()
	return results
}
