package internal

import (
	"strings"

	"github.com/codecrafters-io/interpreter-tester/internal/interpreter_executable"
	testcases "github.com/codecrafters-io/interpreter-tester/internal/test_cases"

	"github.com/codecrafters-io/tester-utils/random"
	"github.com/codecrafters-io/tester-utils/test_case_harness"
)

func testParen(stageHarness *test_case_harness.TestCaseHarness) error {
	b := interpreter_executable.NewInterpreterExecutable(stageHarness)

	logger := stageHarness.Logger

	shuffledString1 := strings.Join(random.RandomElementsFromArray(PARENS, 5), "")
	shuffledString2 := strings.Join(random.RandomElementsFromArray(PARENS, 7), "")
	tokenizeTestCases := testcases.MultiTestCase{
		TestCases: []testcases.TestCase{
			&testcases.TokenizeTestCase{FileContents: "(", ExpectsError: false},
			&testcases.TokenizeTestCase{FileContents: "))", ExpectsError: false},
			&testcases.TokenizeTestCase{FileContents: shuffledString1, ExpectsError: false},
			&testcases.TokenizeTestCase{FileContents: shuffledString2, ExpectsError: false},
		},
	}
	return tokenizeTestCases.RunAll(b, logger)
}
