package internal

import (
	"slices"

	"github.com/codecrafters-io/interpreter-tester/internal/interpreter_executable"
	testcases "github.com/codecrafters-io/interpreter-tester/internal/test_cases"

	"github.com/codecrafters-io/tester-utils/test_case_harness"
)

var Keywords = []string{"and", "class", "else", "false", "for", "fun", "if", "nil", "or", "return", "super", "this", "true", "var", "while"}
var KeywordsCapitalized = []string{"AND", "CLASS", "ELSE", "FALSE", "FOR", "FUN", "IF", "NIL", "OR", "RETURN", "SUPER", "THIS", "TRUE", "VAR", "WHILE"}

func testReservedWords(stageHarness *test_case_harness.TestCaseHarness) error {
	b := interpreter_executable.NewInterpreterExecutable(stageHarness)

	logger := stageHarness.Logger

	k1 := randomSelection(1, Keywords, "")
	k2 := randomSelection(5, slices.Concat(Keywords, KeywordsCapitalized), " ")
	k3 := `var greeting = "Hello"
if (greeting == "Hello") {
    return true
} else {
    return false
}`
	k4 := `var result = (a + b) > 7 && "Success" != "Failure" or x >= 5
while (result) {
    var counter = 0
    counter = counter + 1
    if (counter == 10) {
        return nil
    }
}`

	commandTestCases := testcases.MultiTokenizeTestCase{
		FileContents: []string{k1, k2, k3, k4},
	}
	if err := commandTestCases.RunAll(b, logger); err != nil {
		return err
	}

	return nil
}
