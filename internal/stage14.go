package internal

import (
	"github.com/codecrafters-io/interpreter-tester/internal/interpreter_executable"
	testcases "github.com/codecrafters-io/interpreter-tester/internal/test_cases"

	"github.com/codecrafters-io/tester-utils/random"
	"github.com/codecrafters-io/tester-utils/test_case_harness"
)

var Identifiers = []string{"_hello", "world_", "f00", "6ar", "6az", "foo", "bar", "baz"}

func testIdentifier(stageHarness *test_case_harness.TestCaseHarness) error {
	b := interpreter_executable.NewInterpreterExecutable(stageHarness)

	logger := stageHarness.Logger

	identifier1 := joinWith(random.RandomElementsFromArray(Identifiers[6:], 2), " ")
	identifier2 := "_123" + joinWith(random.RandomElementsFromArray(Identifiers, 5), " ")
	identifier3 := `message = "Hello, World!"
number = 123`
	identifier4 := `{
// This is a complex test case
str1 = "Test" 
str2 = "Case"
num1 = 100
num2 = 200.00
result = (str1 == "Test" , str2 != "Fail") && (num1 + num2) >= 300 && (a - b) < 10
}`

	tokenizeTestCases := testcases.MultiTokenizeTestCase{
		FileContents: []string{identifier1, identifier2, identifier3, identifier4},
	}
	return tokenizeTestCases.RunAll(b, logger)
}
