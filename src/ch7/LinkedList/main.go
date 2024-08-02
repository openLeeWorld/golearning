package main

import (
	"LinkedList/LinkedList"
	"fmt"
)

func main() {
	var ll *LinkedList.LinkedList
	ll = ll.Insert(0, 4)
	ll = ll.Insert(1, 5)
	ll = ll.Insert(2, 9)
	fmt.Println(ll.Value)
	fmt.Println(ll.Next.Value)
	fmt.Println(ll.Next.Next.Value)
}
