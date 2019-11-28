package main

import (
	"bgm38/pkg/server"
)

func main() {

	//resp, err := http.Get(url)
	//if err != nil {
	//	fmt.Println(err)
	//	return
	//}
	//fmt.Println(resp.Proto)
	//
	//fmt.Println(resp)
	//fmt.Println(resp.Header)
	server.Serve()
}
