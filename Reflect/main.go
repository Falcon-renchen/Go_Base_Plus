package main

import (
	"fmt"
	"reflect"
)

type User struct {
	UserId int	`name:"uid" bcd:"345"`
	UserName string
}

func Map2Struct(m map[string]interface{},u interface{})  {
	v := reflect.ValueOf(u)

	if v.Kind()==reflect.Ptr {
		v = v.Elem()
		if v.Kind() != reflect.Struct {
			panic("must struct")
		}
		findFromMap := func(key string,nameTag string) interface{} {
			for k, v := range m {
				if k == key || k == nameTag {
					return v
				}
			}
			return nil
		}
		for i:=0; i<v.NumField(); i++ {
			get_value := findFromMap(v.Type().Field(i).Name,v.Type().Field(i).Tag.Get("name"))
			if get_value != nil && reflect.ValueOf(get_value).Kind()==v.Field(i).Kind() {
				v.Field(i).Set(reflect.ValueOf(get_value))
			}
		}
	} else {
		panic("must ptr")
	}
}

func main() {
	u := &User{101,"wyp"}
	m := map[string]interface{} {
		"id":123,
		"uid":101,
		"UserName":"wyp",
		"age":19,
	}

	Map2Struct(m,u)
	fmt.Println(u)


	t := reflect.ValueOf(u)

//	fmt.Println(t.Name(),t.NumField())
	if t.Kind() == reflect.Ptr {
		t = t.Elem()//把t变成指针指向的变量/内容
	}

	for i:=0; i<t.NumField(); i++ {
		if t.Field(i).Kind()==reflect.Int {
			fmt.Println(t.Field(i).Int())
		}
	}

	for i:=0; i<t.NumField(); i++ {
		if t.Field(i).Kind()==reflect.String {
			fmt.Println(t.Field(i).String())
		}
	}

	for i:=0; i<t.NumField(); i++ {
		fmt.Println(t.Field(i).Interface())
		if t.Field(i).Kind() == reflect.Int {
			t.Field(i).SetInt(12)
			t.Field(i).Set(reflect.ValueOf(13))
		}
		if t.Field(i).Kind() == reflect.String {
			t.Field(i).SetString("wyp")
		}

	}
	fmt.Println(u)



	values := []interface{}{12,"wyp"}
	for i:=0; i<t.NumField(); i++ {
		if t.Field(i).Kind() == reflect.ValueOf(values[i]).Kind() {
			t.Field(i).Set(reflect.ValueOf(values[i]))
		}
	}
	fmt.Println(u)
}
