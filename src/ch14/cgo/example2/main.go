package main

import "fmt"

/*
	#cgo LDFLAGS:-L. -llibc
	extern int add(int a, int b);
*/
import "C"

//export doubler
func doubler(i int) int {
	return i * 2
} // export 주석으로 c코드로 go함수 노출

func main() {
	sum := C.add(3, 2)
	fmt.Println(sum)
}
