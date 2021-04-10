package Test

import (
	"Go_Base/Utils"
	"strconv"
	"testing"
)



func Test_str(t *testing.T)  {
	t.Log("abc")
}

func Test_join(t *testing.T)  {
	str := Utils.Join("abc","bcd")
	t.Log(str)
}


//表驱动测试
func TestJoin(t *testing.T) {
	type args struct {
		strs []string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
		{"t1",args{[]string{"a","b"}},"ab"},
		{"t2",args{[]string{"a","b","c"}},"abc"},
		{"t3",args{[]string{"a","b","c","d"}},"abcd"},

	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Utils.Join(tt.args.strs...); got != tt.want {
				t.Errorf("Join() = %v, want %v", got, tt.want)
			}
		})
	}
}

//基准测试
func BenchmarkJoin(b *testing.B)  {
	for i:=0; i<b.N; i++ {
		Utils.Join(strconv.Itoa(i))
	}
}

var strs = []string{"a","b","c","d","e"}

func BenchmarkJoin2(b *testing.B)  {
	for i:=0; i<b.N; i++ {
		Utils.Join2(strs...)//可变参数
	}
}

func BenchmarkJoin3(b *testing.B)  {
	for i:=0; i<b.N; i++ {
		Utils.Join3(strs...)//可变参数
	}
}
