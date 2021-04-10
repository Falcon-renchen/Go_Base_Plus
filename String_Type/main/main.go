package main

import (
	"Go_Base/String_Type"
	"fmt"
)

func main() {
	//s := String_Type.NewString("adc")
	//s := String_Type.From("adc")

	s := String_Type.FromInt(123)	//
	fmt.Println(s,s.Len())

	str := String_Type.From("abcde")

	//回调
	str.ForEach(func(item string) {
		fmt.Println(item)
	})

	str2 := String_Type.From("我abc")
	for i:=0; i<len(str2); i++ {
		fmt.Println(i," ", fmt.Sprintf("%c",str2[i]))
	}
	//回调函数
	str2.ForEach2(func(s string) {
		fmt.Println(s)
	})

	for i:=0; i<len(str2); i++ {
		fmt.Println(i," ",str2[i],fmt.Sprintf("%T",str2[i]))
	}
	fmt.Println([]byte(str2)) //uint8
	for _,c := range str {
		fmt.Println(c)
	}
	fmt.Println([]rune(str2)) //int32

}
