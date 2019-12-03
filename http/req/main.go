package main

import (
	"fmt"
	"github.com/imroc/req"
)

func main() {
	r, err = req.Post(
		"http://127.0.0.1:8702/msg/list",
		req.Header{
			"Accept": "application/json",
		})
	if err != nil {
		fmt.Println(err)
	}
	r.ToJSON(&foo)       // response => struct/map
	fmt.Printf("%+v", r) // print info (try it, you may surprise)
}
