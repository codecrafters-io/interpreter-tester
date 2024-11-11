package internal

import (
	"strings"

	"github.com/codecrafters-io/interpreter-tester/internal/interpreter_executable"
	testcases "github.com/codecrafters-io/interpreter-tester/internal/test_cases"

	"github.com/codecrafters-io/tester-utils/random"
	"github.com/codecrafters-io/tester-utils/test_case_harness"
)

func testIdentifier(stageHarness *test_case_harness.TestCaseHarness) error {
	b := interpreter_executable.NewInterpreterExecutable(stageHarness)

	logger := stageHarness.Logger

	identifier1 := strings.Join(random.RandomElementsFromArray(IDENTIFIERS[6:], 2), " ")
	identifier2 := "_123" + strings.Join(random.RandomElementsFromArray(IDENTIFIERS, 5), " ")
	identifier3 := `message = "Hello, World!"
number = 123`
	identifier4 := `{
// This is a complex test case
str1 = "Test"
str2 = "Case"
num1 = 100
num2 = 200.00
result = (str1 == str2) != ((num1 + num2) >= 300)
}`

	tokenizeTestCases := testcases.MultiTestCase{
		TestCases: []testcases.TestCase{
			&testcases.TokenizeTestCase{FileContents: identifier1, ExpectsError: false},
			&testcases.TokenizeTestCase{FileContents: identifier2, ExpectsError: false},
			&testcases.TokenizeTestCase{FileContents: identifier3, ExpectsError: false},
			&testcases.TokenizeTestCase{FileContents: identifier4, ExpectsError: false},
		},
	}
	return tokenizeTestCases.RunAll(b, logger)
}
