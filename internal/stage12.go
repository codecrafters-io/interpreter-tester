package internal

import (
	"strings"

	"github.com/codecrafters-io/interpreter-tester/internal/interpreter_executable"
	testcases "github.com/codecrafters-io/interpreter-tester/internal/test_cases"

	"github.com/codecrafters-io/tester-utils/random"
	"github.com/codecrafters-io/tester-utils/test_case_harness"
)

func testStrings(stageHarness *test_case_harness.TestCaseHarness) error {
	b := interpreter_executable.NewInterpreterExecutable(stageHarness)

	logger := stageHarness.Logger

	string1 := random.RandomElementFromArray(QUOTED_STRINGS) + " " + random.RandomElementFromArray(SINGLE_CHAR_OPERATORS) + " " + "\"unterminated"
	string2 := `"foo 	bar 123 // hello world!"`
	shuffledString2 := "(" + strings.Join(random.RandomElementsFromArray(QUOTED_STRINGS, 2), "+") + `) != "other_string"`

	tokenizeTestCases := testcases.MultiTestCase{
		TestCases: []testcases.TestCase{
			&testcases.TokenizeTestCase{FileContents: QUOTED_STRINGS[0], ExpectsError: false},
			&testcases.TokenizeTestCase{FileContents: string1, ExpectsError: true},
			&testcases.TokenizeTestCase{FileContents: string2, ExpectsError: false},
			&testcases.TokenizeTestCase{FileContents: shuffledString2, ExpectsError: false},
		},
	}
	return tokenizeTestCases.RunAll(b, logger)
}
