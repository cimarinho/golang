package main

type I3 interface {
	M()
}
type T3 struct {}

func (T3) M() {}
type T4 struct {}
func (T4) M() {}
func main() {
	var i I = T3{}
	i = T2{}
	_ = i
}
