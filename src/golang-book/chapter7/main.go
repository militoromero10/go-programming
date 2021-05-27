package main

import (
	"fmt"
)

func main() {
	//var scores = [5]float64{4.8, 3.2, 4.5, 1.2, 3.0}
	//fmt.Println(average(scores[:]))

	nextFunc := makeEvenGenerator()
	fmt.Println(nextFunc())
	fmt.Println(nextFunc())
	fmt.Println(nextFunc())

	anotherOne := makeOddGenerator()
	fmt.Println(anotherOne())
	fmt.Println(anotherOne())
	fmt.Println(anotherOne())

}

func makeOddGenerator() func() int {
	k := -1
	return func() int {
		k += 2
		return k

	}
}

func makeEvenGenerator() func() uint {
	i := uint(0)
	return func() (ret uint) {
		ret = i
		i += 2
		return ret
	}
}

func fib(n int) int {
	if n == 0 || n == 1 {
		return n
	}
	return fib(n-1) + fib(n-2)
}

func halves(x int) (int, bool) {
	if x%2 == 0 {
		return 0, false
	} else {
		return 1, true
	}
}

func average(values []float64) float64 {
	total := 0.0
	for _, ele := range values {
		total += ele
	}
	return total / float64(len(values))
}

func f1() {
	fmt.Println("funcion 1 ok")
}

func f2() {
	fmt.Println("function 2 ok")
}

func max(xs ...int) int {
	if len(xs) == 0 {
		return -1
	}

	max := xs[0]

	for _, val := range xs {
		if val > max {
			max = val
		}
	}
	return max
}
