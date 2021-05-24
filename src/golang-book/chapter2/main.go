package main

import "fmt"

// this is a comment

func main() {
	fmt.Println("F° -> C°", "ingrese el valor: ")
	var F float64
	fmt.Scanf("%f", &F)
	C := (F - 32) * 5 / 9
	fmt.Println(C)
}
