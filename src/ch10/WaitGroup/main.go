package main

import (
	"fmt"
	"sync"
)

func main() { // main 고루틴
	var wg sync.WaitGroup
	wg.Add(3)
	go func() {
		defer wg.Done()
		fmt.Println("first function")
	}()
	go func() {
		defer wg.Done()
		fmt.Println("second function")
	}()
	go func() {
		defer wg.Done()
		fmt.Println("third function")
	}()
	wg.Wait()
}
