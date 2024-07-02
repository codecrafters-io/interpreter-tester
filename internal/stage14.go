package internal

import (
	"slices"
	"strings"

	"github.com/codecrafters-io/interpreter-tester/internal/interpreter_executable"
	testcases "github.com/codecrafters-io/interpreter-tester/internal/test_cases"

	"github.com/codecrafters-io/tester-utils/random"
	"github.com/codecrafters-io/tester-utils/test_case_harness"
)

var Identifiers = []string{"_hello", "world_", "f00", "6ar", "6az", "foo", "bar", "baz"}

func testIdentifier(stageHarness *test_case_harness.TestCaseHarness) error {
	b := interpreter_executable.NewInterpreterExecutable(stageHarness)

	logger := stageHarness.Logger

	identifier1 := strings.Join(random.RandomSelection(3, Identifiers), " ")
	identifier2 := "_123" + strings.Join(random.RandomSelection(5, Identifiers), " ")
	identifier3 := strings.Join(random.RandomSelection(5, Identifiers), " ") + "_123"
	identifier4 := "{{" + strings.Join(random.RandomSelection(1, Identifiers), "") + random.RandomStringFromCharacters(2, slices.Concat(SingleCharOperators, LexicalErrors, Relational, Whitespace)) + strings.Join(random.RandomSelection(1, Identifiers), "") + random.RandomStringFromCharacters(2, slices.Concat(SingleCharOperators, LexicalErrors, Relational, Whitespace)) + strings.Join(random.RandomSelection(1, Identifiers), "") + "}}"

	commandTestCases := testcases.MultiTokenizeTestCase{
		FileContents: []string{identifier1, identifier2, identifier3, identifier4},
	}
	if err := commandTestCases.RunAll(b, logger); err != nil {
		return err
	}

	return nil
}
