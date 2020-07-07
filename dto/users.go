package dto

import "net/http"

//Users 实体类用户
type Users struct {
	Name string
	Sex  string
}

//UnmarshalerHTTP 实现方法
func (u *Users) UnmarshalerHTTP(*http.Request) error {
	u.Name = "zxs"
	u.Sex = "man"
	return nil
}

//UsersA 实体类用户
type UsersA struct {
	Age  int
	Size float64
}

//UnmarshalerHTTP 实现方法
func (u *UsersA) UnmarshalerHTTP(*http.Request) error {
	u.Age = 28
	u.Size = 15
	return nil
}
