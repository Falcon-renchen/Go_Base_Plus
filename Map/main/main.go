package main

import (
	"Go_Base/Map"
	"fmt"
)
//传 map，channel，slice不需要写指针，，
//按址传递
func change(u Map.User)  {
	u["id"] = "404"
}

func main() {
	// k值必须是可比较的值
	// v 可以写成func
	u := Map.NewUser()
	u["id"] = "101"
	u["name"] = "wyp"

	change(u)
	fmt.Println(u)


	u.With("id","101")
	u.With("name","wyp")
	fmt.Println(u)
	u.String()

	//sort.SliceStable() //保证顺序
	//sort.Slice(u, func(i, j int) bool {
	//	//id1 := u[i]["id"].(int) //断言功能,就是断定是这个类型 专针对interface
	//	//id2 := u[j]["id"].(int)
	//	//return id1 > id2
	//
	//})
}