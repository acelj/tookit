#include "test.h"
#include <stdio.h>
#include <unistd.h>
#include <string>
#include <iostream>

GoString buildGoString(const char* p, size_t n)
{
    return {p, static_cast<ptrdiff_t>(n)};
}

// 不需要重新定义Foo_return
// g++ main.cpp ./test.so
int main()
{
    printf("demo...\n");
    int re = Add(1,2);
    printf("re = %d\n", re);

    std::string str = "liudehua";
    struct Foo_return res;
    res = Foo(buildGoString(str.c_str(), str.size()), 111);
    printf("re = %lld,%lld\n",res.r0, res.r1); 
    return 0;
}
