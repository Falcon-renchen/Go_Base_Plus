package Utils

import (
	"bytes"
	"strings"
)

//累加字符串
func Join(strs ...string) string {
	ret := ""
	for _, s := range strs {
		ret+=s
	}
	return ret
}

func abc()  {
	//if UserLogin("wyp","123")=="OK" {
	//	ok
	//}
	//if UserLogin("wyp","")=="error" {
	//	!ok
	//}
}

func Join2(strs ...string) string {
	return strings.Join(strs,"")
}

func Join3(strs ...string) string {
	var bf bytes.Buffer
	for _, s := range strs {
		bf.WriteString(s)
	}
	return bf.String()
}