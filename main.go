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

func (u User) GetName() string {
	return u.Name
}
func (u User) PrintName() {
	fmt.Println(u.Name)
}

type User2 struct {
	User
	id2 int
}

func main() {
	u1 := User{11, "oleg", 18}
	fmt.Println(u1.GetName())
	var ss User2
	fmt.Println(ss.User.GetName())
	// еще раз посмотреть на использование структуры внудри другой структуры.....
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
