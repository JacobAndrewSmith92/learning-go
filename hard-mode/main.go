package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	// Udemy Solution
	// f, err := os.Open(os.Args[1])
	// if err != nil {
	// 	fmt.Println("File reading error", err)
	// 	return
	// }
	// io.Copy(os.Stdout, f)

	// My Solution
	f, err := ioutil.ReadFile(os.Args[1])
	if err != nil {
		fmt.Println("File reading error", err)
		return
	}
	fmt.Println(string(f))
}
