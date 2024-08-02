package main

import (
	"fmt"
	"reflect"
	"unsafe"
)

func main() {
	s := "hello"
	sHdrData := unsafe.Pointer((*reflect.StringHeader)(unsafe.Pointer(&s)).Data)
	sHdr := (*reflect.StringHeader)(unsafe.Pointer(&s))
	fmt.Println(sHdr.Len) // 5

	for i := 0; i < sHdr.Len; i++ {
		bp := *(*byte)(unsafe.Pointer(uintptr(sHdrData) + uintptr(i)))
		fmt.Println(string(bp))
	} // 문자열은 불변이므로 읽기만 하고 수정은 안됨

	// If we need to modify the string, we should create a new string
	newStr := []byte(s)
	newStr[2] = newStr[2] + 1 // Modify the third character
	modifiedStr := string(newStr)
	fmt.Println(modifiedStr)
}
