// hello.cpp
#include <iostream>

extern "C" { // 이름 맹글링을 막고, C링킹 규칙을 강제해서  c라이브러리 및 코드와의 호환성을 보장
    void SayHello(const char* name) {
        std::cout << "Hello, " << name << "!" << std::endl;
    }
}
