package Object

import "Go_Base/Struct"

type User struct {
	Id, Sex int
	Name string
}

func NewUser(Struct.UserAttrFunc, Struct.UserAttrFunc) *User {
	return &User{}
}

func WithUserID(id int) Struct.UserAttrFunc {
	return func(user *Struct.User) {
		user.Id = id
	}
}

func WithUserName(name string) Struct.UserAttrFunc {
	return func(user *Struct.User) {
		user.Name = name
	}
}

func WithUserSex(sex byte) Struct.UserAttrFunc {
	return func(user *Struct.User) {
		user.Sex = int(sex)
	}
}