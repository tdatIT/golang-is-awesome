package oop

import "fmt"

type User struct {
	Id   int
	Name string
	Age  int
	BMI  float32
}

func (user *User) ShowInfo() {
	fmt.Printf("Show info user: {\n"+
		"\tid:%v,\n"+
		"\tname:%v,\n"+
		"\tage:%v,\n"+
		"\tbmi:%v \n}",
		user.Id, user.Name, user.Age, user.BMI)
}

func (user *User) SetId(id int) {
	user.Id = id
}

func (user *User) Run() {
	println("Chạy nhanh vậy anh trai")
}
