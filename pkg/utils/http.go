package utils

import (
	"net/http"
)

func Get(url string) (*http.Response, error) {
	client := &http.Client{}
	request, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return &http.Response{}, err
	}
	//request.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/74.0.3729.169 Safari/537.36")
	//for k, v := range request.Header {
	//	fmt.Println(k, v)
	//}
	//fmt.Println(url)
	resp, err := client.Do(request)
	return resp, err
}
