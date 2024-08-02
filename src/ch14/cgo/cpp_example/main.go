package main

/*
#cgo CXXFLAGS: -std=c++11
#cgo LDFLAGS: -L. -lhello

#include <stdlib.h>

extern void SayHello(const char* name);
*/
import "C"
import "unsafe"

func main() {
	name := C.CString("World")
	defer C.free(unsafe.Pointer(name))

	C.SayHello(name)
}
