#pragma once

class TestCase
{
public:
    TestCase(const char* case_name) : testcase_name(case_name){}

    // 执行测试案例的方法
    virtual void Run() = 0;

    int nTestResult; // 测试案例的执行结果 
    const char* testcase_name; // 测试案例名称
};
