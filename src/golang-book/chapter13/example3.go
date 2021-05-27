package main

import (
	"errors"
	"fmt"
	"os"
)

func main() {
	file, err := os.Create("out/test1.txt")

	if err != nil {
		error1 := errors.New("An error has ocurred 1")
		fmt.Println(error1)
		return
	}
	defer file.Close()
	file.WriteString("testing ... 1")

}
