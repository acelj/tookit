#include <gtest/gtest.h>
#include "fun.h"

TEST(fun, add) {
    EXPECT_EQ(4, add(3, 4));   // 断言
    EXPECT_EQ(8, add(4, 3));
}

int main(int argc, char* argv[])
{
    testing::InitGoogleTest(&argc, argv);  //初始化，所有测试从这启动
    return RUN_ALL_TESTS(); // 运行所有测试用例
}
