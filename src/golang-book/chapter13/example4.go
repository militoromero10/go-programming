package main

import (
	"errors"
	"fmt"
	"os"
)

func main() {
	dir, err := os.Open(".")
	if err != nil {
		e := errors.New("An error has ocurred 1")
		fmt.Println(e)
		return
	}
	defer dir.Close()

	filesInfo, err := dir.ReadDir(-1)
	if err != nil {
		e := errors.New("An error has ocurred 2")
		fmt.Println(e)
		return
	}

	for _, fi := range filesInfo {
		fmt.Println(fi.Name())
	}

}
