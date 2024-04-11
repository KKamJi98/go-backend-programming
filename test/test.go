package main

import "fmt"

var y int // zero value => 0

func test() {
	y = 7
	fmt.Println(y)
	fmt.Printf("%T", y)
}
