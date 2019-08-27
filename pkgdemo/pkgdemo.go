/*
   author:Yekai
   company:Pdj
   filename:pkgdemo.go
*/
package pkgdemo

import (
	"fmt"

	_ "github.com/yekai1003/tutorial/mathdemo"
)

//外部可导出结构体
type ExternalPerson struct {
	Name string //大写，可导出
	Age  int
	sex  string //小写，不可导出
}

//内部不可导出结构体
type internalPerson struct {
	Name string
	Age  int
	sex  string
}

func init() {
	fmt.Println("pkgdemo'init is called")
}

//外部可导出函数
func NewPerson(name string, age int, sex string) *ExternalPerson {
	return &ExternalPerson{name, age, sex}
}

func NewInternalPerson(name string, age int, sex string) *internalPerson {
	return &internalPerson{name, age, sex}
}
