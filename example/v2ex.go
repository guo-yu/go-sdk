package main

import "fmt"
import "github.com/turingou/go-sdk"

func main() {
	// Init a SDK instance
	v2ex := &SDK{
		"hotTopics": {
			"method": "get",
			"uri":    "topics/hot.json",
		},
		"latestTopics": {
			"method": "get",
			"uri":    "topics/latest.json",
		},
		"member": {
			"method": "get",
			"uri":    "members/show.json",
		},
	}

	body, res, err := v2ex.member(map[string]string{"username": "turing"})

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(body)
	fmt.Println(body["bio"])
}
