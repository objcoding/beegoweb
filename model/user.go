package model

import "fmt"

type User struct {
	Name string `json:"name"`
	Age int `json:"age"`
}

func (u *User) play() {
	fmt.Println("玩")
}

func (u *User) eat()  {
	fmt.Println("吃")
}