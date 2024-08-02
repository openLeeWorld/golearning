package main

import (
	"BST/tree"
	"fmt"
)

func main() {
	var it *tree.IntTree // tree패키지의 IntTree type 포인터
	it = it.Insert(5)
	it = it.Insert(3)
	it = it.Insert(10)
	it = it.Insert(2)
	fmt.Println(it.Contains(2))  // true
	fmt.Println(it.Contains(12)) // false
}
