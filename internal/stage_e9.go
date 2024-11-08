package internal

import (
	"fmt"

	"github.com/codecrafters-io/interpreter-tester/internal/interpreter_executable"
	testcases "github.com/codecrafters-io/interpreter-tester/internal/test_cases"

	"github.com/codecrafters-io/tester-utils/random"
	"github.com/codecrafters-io/tester-utils/test_case_harness"
)

func testEvaluateEquality(stageHarness *test_case_harness.TestCaseHarness) error {
	b := interpreter_executable.NewInterpreterExecutable(stageHarness)

	logger := stageHarness.Logger

	strings := random.RandomElementsFromArray(STRINGS, 2)
	s1, s2 := strings[0], strings[1]
	n1, n2, n3 := getRandInt(), getRandInt(), getRandInt()
	equality1 := fmt.Sprintf("\"%s\" != \"%s\"", s1, s2)
	equality2 := fmt.Sprintf("\"%s\" == \"%s\"", s1, s1)
	equality3 := fmt.Sprintf("%d == \"%d\"", n1, n1)
	equality4 := fmt.Sprintf("%d == (%d + %d)", n2+n3, n2, n3)

	evaluateTestCases := testcases.MultiTestCase{
		TestCases: []testcases.TestCase{
			&testcases.EvaluateTestCase{FileContents: equality1, ExpectsError: false},
			&testcases.EvaluateTestCase{FileContents: equality2, ExpectsError: false},
			&testcases.EvaluateTestCase{FileContents: equality3, ExpectsError: false},
			&testcases.EvaluateTestCase{FileContents: equality4, ExpectsError: false},
		},
	}
	return evaluateTestCases.RunAll(b, logger)
}
