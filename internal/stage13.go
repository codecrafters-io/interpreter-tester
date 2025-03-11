package internal

import (
	"fmt"

	"github.com/codecrafters-io/interpreter-tester/internal/interpreter_executable"
	testcases "github.com/codecrafters-io/interpreter-tester/internal/test_cases"

	"github.com/codecrafters-io/tester-utils/test_case_harness"
)

func testNumbers(stageHarness *test_case_harness.TestCaseHarness) error {
	b := interpreter_executable.NewInterpreterExecutable(stageHarness)

	logger := stageHarness.Logger

	shuffledString1 := fmt.Sprintf(`(%d+%d) > %d != ("Success" != "Failure") != (%d >= %d)`, getRandInt(), getRandInt(), getRandInt(), getRandInt(), getRandInt())

	tokenizeTestCases := testcases.MultiTestCase{
		TestCases: []testcases.TestCase{
			&testcases.TokenizeTestCase{FileContents: getRandIntAsString(), ExpectsError: false},
			&testcases.TokenizeTestCase{FileContents: fmt.Sprintf("%d%d.%d%d", getRandInt(), getRandInt(), getRandInt(), getRandInt()), ExpectsError: false},
			&testcases.TokenizeTestCase{FileContents: fmt.Sprintf("%d.0000", getRandInt()), ExpectsError: false},
			&testcases.TokenizeTestCase{FileContents: shuffledString1, ExpectsError: false},
		},
	}
	return tokenizeTestCases.RunAll(b, logger)
}
