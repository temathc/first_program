package main

import (
	"encoding/json"
	"fmt"
)

type User struct {
	Id   int
	Name string
	Age  int
}

func (u User) isA() bool {
	return u.Age >= 18
}

func main() {
	u, n := User{11, "oleg", 18}, User{14, "dima", 11}
	fmt.Println(u.isA(), n.isA())
}

func Pars() {
	data := []byte(`
	{
		"ok": true,
		"total": 3,
		"documents": [
			{"id":11, "title": "c++"},
			{"id":2, "title": "suddenly perl"},
			{"id":5, "title": "go"}
		]
	}
	`)
	var resp struct {
		DocumentsQuantity []struct {
			Id    int    `json:"id"`
			Title string `json:"title"`
		} `json:"documents"`
	}
	err := json.Unmarshal(data, &resp)
	if err != nil {
		panic(err)
	}
	fmt.Println(resp.DocumentsQuantity)
}
