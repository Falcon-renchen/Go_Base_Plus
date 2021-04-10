package String_Type

import (
	"fmt"
)

type String string

func (this String) Len() int {
	return len(this) //输出长度
}

func NewString(str string) String {
	//return str  //直接写会报错，因为str是string格式，不是String格式
	return String(str)
}

func From(str string) String {
	//return str  //直接写会报错，因为str是string格式，不是String格式
	return String(str)
}

func FromInt(n int) String {
	//return String(n)  //会输出{  直接输出{  因为ascii的123的 {
	//return String(strconv.Itoa(n))
	return String(fmt.Sprintf("%d",n)) //一样
}

func (this String) ForEach(f func(item string)) {
	for i:=0; i<len(this); i++ {
		//fmt.Println(i," ",fmt.Sprintf("%c",str[i]))//打印单个字符
		f(fmt.Sprintf("%c",this[i]))

	}
}

func (this String) ForEach2(f func(s string)) {
	for index,c := range this {
		fmt.Println(index," ", fmt.Sprintf("%c",c))
	}
}