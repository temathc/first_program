package main

import (
	"fmt"
)

type User struct {
	name string
	age  int
}

func (u *User) SetName(name string) int {
	u.name = name
	return 333
}

func main() {
	u := User{"Oleg", 18}
	fmt.Println(u)
	fmt.Println(u.SetName("Vasya"))
	fmt.Println(u)
}
