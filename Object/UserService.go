package Object

import "log"

type UserService struct {

}

func NewUserService() *UserService {
	return &UserService{}
}

//断言 针对interface进行的类型判断
//保存入库
func (this *UserService) Save(data interface{}) IService {
	//如果user实体，则成功，否则则显示参数错误
	if user,ok:=data.(*User);ok {
		log.Println("%v",user.Name)
		log.Println("用户保存入库成功")
	} else {
		log.Fatal("用户参数错误")
	}

	return this	//链式调用
}

func (this *UserService) List() IService {
	log.Println("用户商品列表调用")
	return this  //链式调用
}