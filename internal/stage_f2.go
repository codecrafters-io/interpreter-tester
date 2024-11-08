package internal

import (
	"strings"

	"github.com/codecrafters-io/interpreter-tester/internal/assertions"
	"github.com/codecrafters-io/interpreter-tester/internal/interpreter_executable"
	testcases "github.com/codecrafters-io/interpreter-tester/internal/test_cases"
	"github.com/codecrafters-io/tester-utils/test_case_harness"
)

func testFunctionsNoArgs(stageHarness *test_case_harness.TestCaseHarness) error {
	b := interpreter_executable.NewInterpreterExecutable(stageHarness)

	logger := stageHarness.Logger

	parsedTestCases := GetTestCasesForCurrentStageWithRandomValues("f2") // This parses the files in f2 and returns the fileContents and parsed frontMatter

	firstTestCase := parsedTestCases[0]
	firstTestExecutionResult := executeWithGolox(firstTestCase.FileContents)

	secondTestCase := parsedTestCases[1]
	secondTestExecutionResult := executeWithGolox(secondTestCase.FileContents)

	thirdTestCase := parsedTestCases[2]
	thirdTestExecutionResult := executeWithGolox(thirdTestCase.FileContents)

	fourthTestCase := parsedTestCases[3]
	fourthTestExecutionResult := executeWithGolox(fourthTestCase.FileContents)

	runTestCases := testcases.MultiTestCase{
		TestCases: []testcases.TestCase{
			&testcases.CustomAssertionRunTestCase{FileContents: firstTestCase.FileContents, FrontMatter: firstTestCase.FrontMatter, Assertion: GetStdoutAssertion(firstTestExecutionResult), ExpectedResult: firstTestExecutionResult},
			&testcases.CustomAssertionRunTestCase{FileContents: secondTestCase.FileContents, FrontMatter: secondTestCase.FrontMatter, Assertion: GetStdoutAssertion(secondTestExecutionResult), ExpectedResult: secondTestExecutionResult},
			&testcases.CustomAssertionRunTestCase{FileContents: thirdTestCase.FileContents, FrontMatter: thirdTestCase.FrontMatter, Assertion: GetStdoutAssertion(thirdTestExecutionResult), ExpectedResult: thirdTestExecutionResult},
			&testcases.CustomAssertionRunTestCase{FileContents: fourthTestCase.FileContents, FrontMatter: fourthTestCase.FrontMatter, Assertion: GetStdoutAssertion(fourthTestExecutionResult), ExpectedResult: fourthTestExecutionResult},
		},
	}

	return runTestCases.RunAll(b, logger)
}

func GetStdoutAssertion(testExecutionResult testcases.GoloxExecutionResult) assertions.Assertion {
	expectedStdoutLines := strings.Split(testExecutionResult.Stdout, "\n")
	return assertions.NewStdoutAssertion(expectedStdoutLines)
}
