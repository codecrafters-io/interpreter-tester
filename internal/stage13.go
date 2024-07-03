package internal

import (
	"github.com/codecrafters-io/interpreter-tester/internal/interpreter_executable"
	testcases "github.com/codecrafters-io/interpreter-tester/internal/test_cases"

	"github.com/codecrafters-io/tester-utils/test_case_harness"
)

func testNumbers(stageHarness *test_case_harness.TestCaseHarness) error {
	b := interpreter_executable.NewInterpreterExecutable(stageHarness)

	logger := stageHarness.Logger

	shuffledString1 := `"Hello" = "Hello" && 42 == 42`
	shuffledString2 := `(5+3) > 7 ; "Success" != "Failure" & 10 >= 5`

	commandTestCases := testcases.MultiTokenizeTestCase{
		FileContents: []string{"1234.1234", "1234.1234.1234.", shuffledString1, shuffledString2},
	}
	if err := commandTestCases.RunAll(b, logger); err != nil {
		return err
	}

	return nil
}
