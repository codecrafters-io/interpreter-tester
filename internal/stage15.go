package internal

import (
	"slices"
	"strings"

	"github.com/codecrafters-io/grep-tester/internal/interpreter_executable"
	testcases "github.com/codecrafters-io/grep-tester/internal/test_cases"

	"github.com/codecrafters-io/tester-utils/random"
	"github.com/codecrafters-io/tester-utils/test_case_harness"
)

var Keywords = []string{"and", "class", "else", "false", "for", "fun", "if", "nil", "or", "return", "super", "this", "true", "var", "while"}
var KeywordsCapitalized = []string{"AND", "CLASS", "ELSE", "FALSE", "FOR", "FUN", "IF", "NIL", "OR", "RETURN", "SUPER", "THIS", "TRUE", "VAR", "WHILE"}

func testReservedWords(stageHarness *test_case_harness.TestCaseHarness) error {
	b := interpreter_executable.NewInterpreterExecutable(stageHarness)

	logger := stageHarness.Logger

	k1 := Keywords[0]
	k2 := strings.Join(Keywords[0:], " ")
	k3 := strings.Join(KeywordsCapitalized[0:], "	")
	k4 := "(" + random.RandomStringFromCharacters(1, slices.Concat(Identifiers, Strings, Keywords, KeywordsCapitalized)) + random.RandomStringFromCharacters(1, slices.Concat(SingleCharOperators, LexicalErrors, Relational, Whitespace)) + random.RandomStringFromCharacters(1, slices.Concat(Identifiers, Strings, Keywords, KeywordsCapitalized)) + random.RandomStringFromCharacters(1, slices.Concat(SingleCharOperators, LexicalErrors, Relational, Whitespace)) + random.RandomStringFromCharacters(1, slices.Concat(Identifiers, Strings, Keywords, KeywordsCapitalized)) + random.RandomStringFromCharacters(1, slices.Concat(SingleCharOperators, LexicalErrors, Relational, Whitespace)) + random.RandomStringFromCharacters(1, slices.Concat(Identifiers, Strings, Keywords, KeywordsCapitalized)) + random.RandomStringFromCharacters(1, slices.Concat(SingleCharOperators, LexicalErrors, Relational, Whitespace)) + ")"

	commandTestCases := testcases.MultiTokenizeTestCase{
		FileContents: []string{k1, k2, k3, k4},
	}
	if err := commandTestCases.RunAll(b, logger); err != nil {
		return err
	}

	return nil
}
