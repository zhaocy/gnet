package gnet

import (
	"fmt"
	"testing"
)

type A struct {
	Id string
	Name string
	Status int
}

type B struct {
	Id string
	Num int
	Name string
}
func TestCopy(t *testing.T) {
	var a = A{Id:"1", Name:"a",Status:1}
	var b = &B{}
	toMap := StructToMap(a)
	Copy(a,b,"Name")
	fmt.Printf("%v %v %v\n",b.Id,b.Name,b.Num,)
	fmt.Println(toMap["Name"])

}
