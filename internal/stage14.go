package internal

import (
	"slices"
	"strings"

	"github.com/codecrafters-io/grep-tester/internal/interpreter_executable"
	testcases "github.com/codecrafters-io/grep-tester/internal/test_cases"

	"github.com/codecrafters-io/tester-utils/random"
	"github.com/codecrafters-io/tester-utils/test_case_harness"
)

var asciiLower = strings.Split("abcdefghijklmnopqrstuvwxyz", "")
var asciiUpper = strings.Split("ABCDEFGHIJKLMNOPQRSTUVWXYZ", "")
var asciiDigits = strings.Split("0123456789", "")

func testIdentifier(stageHarness *test_case_harness.TestCaseHarness) error {
	b := interpreter_executable.NewInterpreterExecutable(stageHarness)

	logger := stageHarness.Logger

	identifier1 := strings.Join(random.RandomSelection(5, asciiLower), "")
	identifier2 := "_" + strings.Join(random.RandomSelection(5, asciiDigits), "")
	identifier3 := "_" + strings.Join(random.RandomSelection(50, slices.Concat(asciiUpper, asciiLower, asciiDigits, []string{"_"})), "") + "_"
	identifier4 := randomStringFromCharactersNew(50, slices.Concat(Parens, Braces, SingleCharOperators, LexicalErrors, Equals, Negation, Relational, Division, Whitespace, asciiLower, asciiUpper, asciiDigits, []string{"_"}))

	commandTestCases := testcases.MultiTokenizeTestCase{
		FileContents: []string{identifier1, identifier2, identifier3, identifier4},
	}
	if err := commandTestCases.RunAll(b, logger); err != nil {
		return err
	}

	return nil
}
