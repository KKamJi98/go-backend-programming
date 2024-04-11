package main

import "fmt"

func compositeLiteral() {
	xi := []int{2, 4, 7, 9, 42}
	fmt.Println(xi)

	m := map[string]int{
		"James":           32,
		"Miss Moneypenny": 27,
	}
	fmt.Println(m)
}
