package main

import (
	"fmt"
	"strconv"
)

type counter []int

func countToTen() counter {
	c10 := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	return c10
}

func (c counter) print() {
	for _, num := range c {
		fmt.Println(num)
	}
}

func (c counter) evenOrOdd() {
	for _, count := range c {
		if count%2 == 0 {
			cs := strconv.Itoa(count)
			fmt.Println(cs + " is even")
		} else {
			cs := strconv.Itoa(count)
			fmt.Println(cs + " is odd")
		}
	}
}
