package main

import (
	"errors"
	"fmt"
	"io/ioutil"
	//	"os"
)

func main() {
	bs, err := ioutil.ReadFile("test.txt")
	if err != nil {
		err := errors.New("An error has ocurred 1")
		fmt.Println(err)
		return
	}
	str := string(bs)
	fmt.Println(str)

}

/*func main() {
	file, error := os.Open("test.txt")
	if error != nil {
		fmt.Println("An error has ocurred 1")
		return
	}

	defer file.Close()

	//get the file size
	stat, err := file.Stat()
	if err != nil {
		fmt.Println("An error has ocurred 2")
		return
	}

	bs := make([]byte, stat.Size())
	_, err1 := file.Read(bs)
	if err1 != nil {
		fmt.Println("An error has ocurred 3")
		return
	}

	str := string(bs)
	fmt.Println(str)
}*/
