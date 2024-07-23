package internal

import (
	"fmt"
	"strings"

	"github.com/codecrafters-io/interpreter-tester/internal/interpreter_executable"
	testcases "github.com/codecrafters-io/interpreter-tester/internal/test_cases"

	"github.com/codecrafters-io/tester-utils/random"
	"github.com/codecrafters-io/tester-utils/test_case_harness"
)

func testEvaluateParens(stageHarness *test_case_harness.TestCaseHarness) error {
	b := interpreter_executable.NewInterpreterExecutable(stageHarness)

	logger := stageHarness.Logger

	parens1 := "(true)"
	parens2 := fmt.Sprintf("(%d)", getRandInt())
	parens3 := fmt.Sprintf("(\"%s\")", strings.Join(random.RandomElementsFromArray(STRINGS, 2), " "))
	parens4 := fmt.Sprintf("((%s))", random.RandomElementFromArray(BOOLEANS))

	evaluateTestCases := testcases.MultiEvaluateTestCase{
		FileContents: []string{parens1, parens2, parens3, parens4},
	}
	return evaluateTestCases.RunAll(b, logger)
}
