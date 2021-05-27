package main

import "fmt"

func main() {

	/* 2
	x := new(int)
	zero(x)
	fmt.Println(*x)*/

	x := 5
	zero(&x)
	//zero(x) //1
	fmt.Println(x)

	a := 1
	b := 2
	fmt.Println(a, b)
	swap(&a, &b)
	fmt.Println(a, b)

	val := 1.5
	square(&val)
	fmt.Println(val)

}

func zero(value *int) {
	*value = 0
}

/* 1
func zero(value int) {
	value = 0
}
*/

func swap(x, y *int) {
	aux := *x
	*x = *y
	*y = aux
}

func square(x *float64) {
	*x = *x * *x
}
