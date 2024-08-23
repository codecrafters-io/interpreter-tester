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

	shuffledString1 := fmt.Sprintf(`"%s" = "%s" && %d == %d`, getRandString(), getRandString(), getRandInt(), getRandInt())
	shuffledString2 := fmt.Sprintf(`(%d+%d) > %d ; "Success" != "Failure" & %d >= %d`, getRandInt(), getRandInt(), getRandInt(), getRandInt(), getRandInt())

	tokenizeTestCases := testcases.MultiTokenizeTestCase{
		FileContents: []string{getRandIntAsString(), fmt.Sprintf("%d.%d", getRandInt(), getRandInt()), shuffledString1, shuffledString2},
	}
	return tokenizeTestCases.RunAll(b, logger)
}
