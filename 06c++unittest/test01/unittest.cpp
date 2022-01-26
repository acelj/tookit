#include "unittest.h"
#include <iostream>

UnitTest* UnitTest::GetInstance()
{
    static UnitTest instance;
    return &instance;
}

TestCase* UnitTest::RegisterTestCase(TestCase* testcase)
{
    testcases_.push_back(testcase);
    return testcase;
}

int UnitTest::Run()
{
    nTestResult = 1;
    for (std::vector<TestCase*>::iterator it = testcases_.begin();
        it != testcases_.end(); ++it)
    {
        TestCase* testcase = *it;
        CurrentTestCase = testcase;
        std::cout << "======================================" << std::endl;
        std::cout << "Run TestCase:" << testcase->testcase_name << std::endl;
        testcase->Run();
        std::cout << "End TestCase:" << testcase->testcase_name << std::endl;
        if (testcase->nTestResult)
        {
            nPassed++;
        }
        else
        {
            nFailed++;
            nTestResult = 0;
        }
    }

    std::cout << "======================================" << std::endl;
    std::cout << "Total TestCase : " << nPassed + nFailed << std::endl;
    std::cout << "Passed : " << nPassed << std::endl;
    std::cout << "Failed : " << nFailed << std::endl;
    return nTestResult;
}