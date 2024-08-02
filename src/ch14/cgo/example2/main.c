//#include "_cgo_export.h"
#include "libgo.h"

int add(int a, int b) {
    int doubleA = doubler(a);
    int sum = doubleA + b;
    return sum;
}
// go코드를 빌드하면 //export 지시어에 따라 c헤더파일을 자동으로 생성