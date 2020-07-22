package main

import (
	"fmt"
	"reflect"
)

type I interface {
	M()
}
type T1 struct{}

func (T1) M() {
	fmt.Println("t1")
}

type T2 struct{}

func (T2) M() {
	fmt.Println("t2")
}
func main() {
	var i I = T1{}
	i = T2{}

	j := T1{}
	fmt.Println("J=", j)


	//_ = i

	i.M()

	fmt.Println(reflect.TypeOf(i).PkgPath(), reflect.TypeOf(i).Name())
	fmt.Println(reflect.TypeOf(i).String())

	fmt.Printf("%T\n", i)
}
