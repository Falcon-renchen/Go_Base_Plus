package Map

import (
	"fmt"
	"sort"
)

type User2 struct {
	User map[string]string
	Id string
	Name string
	Age int
}

type User map[string]string

func NewUser() User {
	return make(map[string]string)
}

func (this User) With(k string, v string) {
	this[k] = v
}

func (this User) Fields() []string {
	keys := []string{}
	for k, _ := range this {
		keys = append(keys,k)
	}
	sort.Sort(sort.StringSlice(keys))//正排
	sort.Sort(sort.Reverse(sort.StringSlice(keys)))//倒排
	return keys
}

//不是每一次都是按顺序来的,map不能按顺序来的
func (this User) String() string {
	str := ""
	for index,k := range this.Fields() {
		str += fmt.Sprintf("%d、%v->%v",index+1,k,this[k])
	}
	return str
}