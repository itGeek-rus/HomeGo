package main

import "fmt"

//Данный код выведет значние "false"
//Так вышло, потому что "A()" возвращает полностью нулевой интерфейс, в то время как "B()" возвращает интерфейс с типом "*impl", а значение "nil"
//В "Go" сравнения равны только если  только если интерфейсы имеют одинаковый тип и одинаковое значение
//"a" тип "nil", значение "nil", "b" тип "*impl", значение "nil"
type impl struct {
}

type I interface {
	C()
}

func (*impl) C() {

}

func A() I {
	return nil
}

func B() I {
	var ret *impl
	return ret
}

func main() {
	a := A()
	b := B()
	fmt.Println(a == b)
}
