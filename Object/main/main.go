package main

import (
	"Go_Base/Object"
	"fmt"
)

//改变值的两种方法
func change(u *Object.User)  {
	u.Id = 202
}
//func change(u Object.User) Object.User {
//	u.Id = 202
//	return u
//}

//接口
func SaveModel(service Object.IService) Object.IService {
	///service.Save()
	return service
}

func main() {
	u := Object.NewUser(nil, nil)
	u.Id = 101
	change(u)
	//u= change(u)
	fmt.Printf("%+v",u)

	var u1 Object.User
	var u2 *Object.User
	fmt.Println(u1,u2)


	var service Object.IService = Object.NewUserService()
	SaveModel(service)


	//IService 代言了UserService 和 ProdService
	SaveModel(Object.NewProdService()).List()

	//Object.NewProdService().Save(data).List() //链式调用 返回IService
	user := Object.NewUser(
		Object.WithUserName("wyp"),
		Object.WithUserSex(1),
		)
	Object.NewUserService().Save(user)
}
