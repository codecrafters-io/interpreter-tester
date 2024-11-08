package internal

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/codecrafters-io/interpreter-tester/internal/assertions"
	"github.com/codecrafters-io/interpreter-tester/internal/interpreter_executable"
	loxapi "github.com/codecrafters-io/interpreter-tester/internal/lox/api"
	testcases "github.com/codecrafters-io/interpreter-tester/internal/test_cases"
	"github.com/codecrafters-io/tester-utils/test_case_harness"
)

func testClock(stageHarness *test_case_harness.TestCaseHarness) error {
	b := interpreter_executable.NewInterpreterExecutable(stageHarness)

	logger := stageHarness.Logger

	parsedTestCases := GetTestCasesForCurrentStageWithRandomValues("f1") // This parses the files in f1 and returns the fileContents and parsed frontMatter

	firstTestCase := parsedTestCases[0]
	firstTestExecutionResult := executeWithGolox(firstTestCase.FileContents)
	firstAssertion, err := getIntegerAssertion(firstTestExecutionResult)
	if err != nil {
		return err
	}

	runTestCases := testcases.MultiTestCase{
		TestCases: []testcases.TestCase{
			&testcases.CustomAssertionRunTestCase{FileContents: firstTestCase.FileContents, FrontMatter: firstTestCase.FrontMatter, Assertion: firstAssertion, ExpectedResult: firstTestExecutionResult},
		},
	}

	return runTestCases.RunAll(b, logger)
}

func executeWithGolox(fileContents string) testcases.GoloxExecutionResult {
	expectedStdout, expectedExitCode, expectedStderr := loxapi.Run(fileContents)
	return testcases.GoloxExecutionResult{
		Stdout:   expectedStdout,
		Stderr:   expectedStderr,
		Exitcode: expectedExitCode,
	}
}

func getIntegerAssertion(testExecutionResult testcases.GoloxExecutionResult) (assertions.Assertion, error) {
	expectedStdoutLines := strings.Split(testExecutionResult.Stdout, "\n")
	if len(expectedStdoutLines) != 1 {
		return nil, fmt.Errorf("CodeCrafters internal error: expected a single line of output from golox, got %d lines", len(expectedStdoutLines))
	}

	tolerance := 5
	value, err := strconv.ParseFloat(expectedStdoutLines[0], 64)
	if err != nil {
		return nil, fmt.Errorf("CodeCrafters internal error: failed to parse expected output as float: %v", err)
	}
	assertion := assertions.NewIntegerAssertion(int(value)-tolerance, int(value)+tolerance)
	return assertion, nil
}
