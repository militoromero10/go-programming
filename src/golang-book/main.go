package main

import (
	"fmt"
)

/*
"io"
"strings"*/

func main() {
	/*r := strings.NewReader("Hello, Reader!")

	b := make([]byte, 8)
	for {
		n, err := r.Read(b)
		fmt.Printf("n = %v err = %v b = %v\n", n, err, b)
		fmt.Printf("b[:n] = %q\n", b[:n])
		if err == io.EOF {
			break
		}
	}*/

	fmt.Println('A', 'Z', 'a', 'z', ' ')

	for j := 'a'; j <= 'z'; j++ {
		if j+13 > 122 {
			fmt.Printf("%v -> %v\n", string(j), string((65 + ((j + 13) % 91))))
		} else {
			fmt.Printf("%v -> %v\n", string(j), string(j+13))
		}
	}

}
