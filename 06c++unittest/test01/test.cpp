#include "unittest.h"

int Foo(int a, int b)
{
    return a + b;
}

NTEST(FooTest_PassDemo)
{
    EXPECT_EQ(3, Foo(1, 2));
    EXPECT_EQ(2, Foo(1, 1));
}

NTEST(FooTest_FailDemo)
{
    EXPECT_EQ(4, Foo(1, 2));
    EXPECT_EQ(2, Foo(1, 2));
}

int main(int argc, const char* argv[])
{
    return RUN_ALL_TESTS();
}


/**
 ======================================
Run TestCase:FooTest_PassDemo
End TestCase:FooTest_PassDemo
======================================
Run TestCase:FooTest_FailDemo
Failed
Expect:4
Actual:3
Failed
Expect:2
Actual:3
End TestCase:FooTest_FailDemo
======================================
Total TestCase : 2
Passed : 0
Failed : 2
 * 
 */