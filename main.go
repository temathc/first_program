package main

import (
	"encoding/json"
	"fmt"
	"strings"
)

// type User struct {
// 	Id   int
// 	Name string
// 	Age  int
// }

// func (u User) GetName() string {
// 	return u.Name
// }
// func (u User) PrintName() {
// 	fmt.Println(u.Name)
// }

// type User2 struct {
// 	User
// 	id2 int
// }

func main() {
	// u1 := User{11, "oleg", 18}
	// fmt.Println(u1.GetName())
	//var ss User2
	//fmt.Println(ss.User.GetName())
	// еще раз посмотреть на использование структуры внудри другой структуры.....

	d := "cat and dog, one dog,two cats and one man cat and dog, one dog,two cats and one man"
	// sss := strings.Fields(d)
	// sort.Slice(sss, func(i, j int) bool { return sss[i] < sss[j] })
	// // for i := range sss {

	// // }
	// s2 := make(map[string]int)

	// for _, v := range sss {
	// 	_, vv := s2[v]
	// 	if vv {
	// 		s2[v] += 1
	// 	} else {
	// 		s2[v] = 1
	// 	}

	// }
	// //fmt.Println(s2)

	// fmt.Println(regexp.MustCompile(sss[2]))

	tempStr := strings.Fields(d)
	fmt.Println(tempStr)
	m := make(map[string]int)
	for _, word := range tempStr {
		m[word]++
		fmt.Println(m, "\n---")
	}
	//fmt.Println(m)

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
