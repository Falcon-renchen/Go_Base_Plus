package main

import (
	"Go_Base/Struct"
	"fmt"
)

func main() {
	u := Struct.NewUser(
		Struct.WithUserID(120),
		Struct.WithUserName("wyp"),
		)
	fmt.Println(u)
}