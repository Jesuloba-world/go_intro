package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	// file, err := os.Create("a.txt")

	// if err != nil {
	// 	log.Fatal(err)
	// }

	// file.Close()

	// var fileInfo os.FileInfo
	// fileInfo, err = os.Stat("a.txt")
	// p := fmt.Println

	// p(fileInfo.Name())
	// p(fileInfo.Size())
	// p(fileInfo.Mode())
	// p(fileInfo.ModTime())

	// file, err := os.OpenFile(
	// 	"a.txt",
	// 	os.O_WRONLY|os.O_TRUNC|os.O_CREATE,
	// 	0644,
	// )

	// if err != nil {
	// 	log.Fatal(err)
	// }

	// defer file.Close()

	// byteSlice := []byte("I love Golang")
	// byteWrtten, err := file.Write(byteSlice)

	// if err != nil {
	// 	log.Fatal(err)
	// }

	// log.Printf("Byte Written %d\n", byteWrtten)

	scanner := bufio.NewScanner(os.Stdin)

	scanner.Scan()

	text := scanner.Text()

	fmt.Println(strings.Repeat("#", 10))

	fmt.Println("input text", text)
}
