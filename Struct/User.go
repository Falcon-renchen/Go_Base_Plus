package Struct

type UserAttrFunc func(*User) //设置User属性的函数类型
type UserAttrFuncs []UserAttrFunc

//对每一个属性赋值
func (this UserAttrFuncs) apply(u *User) {
	for _, f := range this {
		f(u)
	}
}

type User struct {
	Id, Sex int
	Name string
}

//有选择性的对ID进行赋值
func setid(u *User)  {
	u.Id = 132
}

func NewUser(fs ...UserAttrFunc) *User {
	u := &User{}
	UserAttrFuncs(fs).apply(u)
	for _,f :=range fs{
		f(u)
	}
	return u
}

func WithUserID(id int) func(u *User) {
	return func(u *User) {
		u.Id = id
	}
}

func WithUserName(name string) func(u *User) {
	return func(u *User) {
		u.Name = name
	}
}