package internal

import (
	"fmt"

	"github.com/codecrafters-io/interpreter-tester/internal/interpreter_executable"
	testcases "github.com/codecrafters-io/interpreter-tester/internal/test_cases"
	"github.com/codecrafters-io/tester-utils/random"
	"github.com/codecrafters-io/tester-utils/test_case_harness"
)

func testParseStrings(stageHarness *test_case_harness.TestCaseHarness) error {
	b := interpreter_executable.NewInterpreterExecutable(stageHarness)

	logger := stageHarness.Logger

	randomWords := []string{"foo", "bar", "baz", "quz", "hello", "world"}
	stringLiteral1 := "\"" + randomSelection(2, randomWords, " ") + "\""
	stringLiteral2 := "\"'" + randomSelection(1, randomWords, "") + "'\""
	stringLiteral3 := "\"// " + randomSelection(1, randomWords, "") + "\""
	stringLiteral4 := "\"/* " + fmt.Sprint(random.RandomInt(10, 100)) + " */\""

	parseTestCase := testcases.MultiParseTestCase{
		FileContents: []string{stringLiteral1, stringLiteral2, stringLiteral3, stringLiteral4},
	}
	return parseTestCase.RunAll(b, logger)
}
