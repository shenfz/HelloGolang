#include <iostream>

extern "C" {
  #include "hello.h"
}

void SayHelloByCMore(const char* s) {
     std::cout << s;
}