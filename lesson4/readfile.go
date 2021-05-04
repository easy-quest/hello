package main

import "os"
import "bufio"

import "bytes"
import "fmt"
import "io/ioutil"

func main() {

	fmt.Println("### Read as reader ###")
	f, err := os.Open("temp/file.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	// Чтение файла с ридером
	wr := bytes.Buffer{}
	sc := bufio.NewScanner(f)
	for sc.Scan() {
		wr.WriteString(sc.Text())
	}
	fmt.Println(wr.String())

	fmt.Println("### ReadFile ###")
	// для более мелких файлов
	fContent, err := ioutil.ReadFile("temp/file.txt")
	if err != nil {
		panic(err)
	}
	fmt.Println(string(fContent))

}
