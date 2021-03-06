package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main() {

	f, err := os.Open("chapter05/recipe03/temp/file.txt")
	if err != nil {
		panic(err)
	}

	c, err := ioutil.ReadAll(f)
	if err != nil {
		panic(err)
	}

	fmt.Printf("### File content ###\n%s\n", string(c))
	f.Close()

	f, err = os.OpenFile("chapter05/recipe03/temp/file.txt", os.O_CREATE|os.O_RDWR, 0777)
	if err != nil {
		panic(err)
	}
	io.WriteString(f, "Test string")
	f.Close()

}
