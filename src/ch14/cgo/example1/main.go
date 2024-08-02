package main

/*
	#cgo LDFLAGS:-L. -lm -lmylib
	#include <stdio.h>
	#include <math.h>
	#include "mylib.h"

	int add(int a, int b) {
		int sum = a + b;
		printf("a: %d, b: %d, sum: %d\n", a, b, sum);
		return sum;
	}
*/
import "C" //  C는 자동으로 생성된 패키지로 식별자는 주석에 포함된 C코드에서 가져온다.
import "fmt"

// 위는 주석 블록 (C)

func main() {
	sum := C.add(3, 2) // 위의 주석에 정의
	fmt.Println(sum)
	fmt.Println(C.sqrt(100))        // math.c에 정의
	fmt.Println(C.multiply(10, 20)) // mylib.dll에 정의
}

// C psuedo-package는 내장 C타입을 나타내는 C.int, C.char 같은 타입과
// Go 문자열을 c 문자열로 변환하는 C.CString같은 함수를 정의한다.
