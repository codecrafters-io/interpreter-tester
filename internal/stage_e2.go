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

	evaluateTestCases := testcases.MultiTestCase{
		TestCases: []testcases.TestCase{
			&testcases.EvaluateTestCase{FileContents: numberLiteral1, ExpectsError: false},
			&testcases.EvaluateTestCase{FileContents: numberLiteral2, ExpectsError: false},
			&testcases.EvaluateTestCase{FileContents: stringLiteral1, ExpectsError: false},
			&testcases.EvaluateTestCase{FileContents: stringLiteral2, ExpectsError: false},
		},
	}
	return evaluateTestCases.RunAll(b, logger)
}
