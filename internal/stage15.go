package internal

import (
	"slices"
	"strings"

	"github.com/codecrafters-io/interpreter-tester/internal/interpreter_executable"
	testcases "github.com/codecrafters-io/interpreter-tester/internal/test_cases"

	"github.com/codecrafters-io/tester-utils/random"
	"github.com/codecrafters-io/tester-utils/test_case_harness"
)

func testReservedWords(stageHarness *test_case_harness.TestCaseHarness) error {
	b := interpreter_executable.NewInterpreterExecutable(stageHarness)

	logger := stageHarness.Logger

	allKeywords := slices.Concat(KEYWORDS, CAPITALIZED_KEYWORDS)
	k1 := random.RandomElementFromArray(KEYWORDS)
	k2 := strings.Join(random.RandomElementsFromArray(allKeywords, len(allKeywords)), " ")
	k3 := `var greeting = "Hello"
if (greeting == "Hello") {
    return true
} else {
    return false
}`
	k4 := `var result = (a + b) > 7 or "Success" != "Failure" or x >= 5
while (result) {
    var counter = 0
    counter = counter + 1
    if (counter == 10) {
        return nil
    }
}`

	tokenizeTestCases := testcases.MultiTestCase{
		TestCases: []testcases.TestCase{
			&testcases.TokenizeTestCase{FileContents: k1, ExpectsError: false},
			&testcases.TokenizeTestCase{FileContents: k2, ExpectsError: false},
			&testcases.TokenizeTestCase{FileContents: k3, ExpectsError: false},
			&testcases.TokenizeTestCase{FileContents: k4, ExpectsError: false},
		},
	}
	return tokenizeTestCases.RunAll(b, logger)
}
