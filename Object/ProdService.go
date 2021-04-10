package Object

import "log"

type ProdService struct {

}

func NewProdService() *ProdService {
	return &ProdService{}
}

//保存入库
func (this *ProdService) Save(data interface{}) IService {
	log.Println("商品保存入库成功")
	return this	//链式调用
}

func (this *ProdService) List() IService {
	log.Println("商品列表调用")
	return this	//链式调用
}