package main

import "fmt"

type I interface {
	f1(name string)
	f2(name string) (error, float32)
	f3() int64
}

type T string

func (T) f1(name string) {
	fmt.Println(name)
}
func (T) f2(name string) (error, float32) {
	return nil, 10.2
}
func (T) f3() int64 {
	return 10
}

func main() {
	var i  T = "rerere"
	i.f1("f1")
	fmt.Println(i.f2("f2"))
	fmt.Println(i.f3())

}
