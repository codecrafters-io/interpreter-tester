package internal

import (
	"fmt"
	"math/rand"
	"strconv"
	"strings"

	"github.com/codecrafters-io/interpreter-tester/internal/assertions"
	"github.com/codecrafters-io/interpreter-tester/internal/interpreter_executable"
	loxapi "github.com/codecrafters-io/interpreter-tester/internal/lox/api"
	testcases "github.com/codecrafters-io/interpreter-tester/internal/test_cases"
	"github.com/codecrafters-io/tester-utils/test_case_harness"
)

var (
	firstTestCaseFileContents = `
print clock() + <<RANDOM_INTEGER>>;
`
	secondTestCaseFileContents = `
print clock() / 1000;
`
)

func testClock(stageHarness *test_case_harness.TestCaseHarness) error {
	b := interpreter_executable.NewInterpreterExecutable(stageHarness)
	logger := stageHarness.Logger

	firstTestCaseFileContents = strings.ReplaceAll(firstTestCaseFileContents, "<<RANDOM_INTEGER>>", fmt.Sprintf("%d", rand.Intn(1000)))
	firstTestCase := buildSingleLineClockOutputTestCase(firstTestCaseFileContents)
	secondTestCase := buildSingleLineClockOutputTestCase(secondTestCaseFileContents)
	otherTestCases := GetTestCasesForCurrentStageWithRandomValues("f1")

	testCases := []testcases.TestCase{}
	testCases = append(testCases, &firstTestCase)
	testCases = append(testCases, &secondTestCase)
	testCases = append(testCases, otherTestCases...)

	runTestCase := testcases.MultiTestCase{TestCases: testCases}
	return runTestCase.RunAll(b, logger)
}

func buildSingleLineClockOutputTestCase(fileContents string) testcases.RunTestCase {
	ourLoxStdout, ourLoxExitCode, _ := loxapi.Run(fileContents)

	if ourLoxExitCode != 0 {
		panic(fmt.Sprintf("CodeCrafters internal error: expected exit code 0 from golox, got %d", ourLoxExitCode))
	}

	assertion, err := buildIntegerAssertion(ourLoxStdout)
	if err != nil {
		panic(err)
	}

	return testcases.RunTestCase{
		FileContents:    fileContents,
		OutputAssertion: assertion,
	}
}

func buildIntegerAssertion(ourLoxStdout string) (assertions.Assertion, error) {
	expectedStdoutLines := strings.Split(ourLoxStdout, "\n")
	if len(expectedStdoutLines) != 1 {
		return nil, fmt.Errorf("CodeCrafters internal error: expected a single line of output from golox, got %d lines", len(expectedStdoutLines))
	}

	value, err := strconv.ParseFloat(expectedStdoutLines[0], 64)
	if err != nil {
		return nil, fmt.Errorf("CodeCrafters internal error: failed to parse expected output (%q) as float: %v", expectedStdoutLines[0], err)
	}

	tolerance := 5
	return assertions.NewIntegerAssertion(int(value)-tolerance, int(value)+tolerance), nil
}
