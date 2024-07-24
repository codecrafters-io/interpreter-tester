package internal

import (
	"fmt"
	"strings"

	"github.com/codecrafters-io/interpreter-tester/internal/interpreter_executable"
	testcases "github.com/codecrafters-io/interpreter-tester/internal/test_cases"

	"github.com/codecrafters-io/tester-utils/random"
	"github.com/codecrafters-io/tester-utils/test_case_harness"
)

func testEvaluateLiterals(stageHarness *test_case_harness.TestCaseHarness) error {
	b := interpreter_executable.NewInterpreterExecutable(stageHarness)

	logger := stageHarness.Logger

	numberLiteral1 := getRandIntAsString()
	numberLiteral2 := fmt.Sprintf("%d.%d", getRandInt(), getRandInt())
	stringLiteral1 := "\"" + strings.Join(random.RandomElementsFromArray(STRINGS, 2), " ") + "\""
	stringLiteral2 := "\"" + getRandIntAsString() + "\""

	evaluateTestCases := testcases.MultiEvaluateTestCase{
		FileContents: []string{numberLiteral1, numberLiteral2, stringLiteral1, stringLiteral2},
	}
	return evaluateTestCases.RunAll(b, logger)
}
