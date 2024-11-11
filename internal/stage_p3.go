package internal

import (
	"fmt"
	"strings"

	"github.com/codecrafters-io/interpreter-tester/internal/interpreter_executable"
	testcases "github.com/codecrafters-io/interpreter-tester/internal/test_cases"
	"github.com/codecrafters-io/tester-utils/random"
	"github.com/codecrafters-io/tester-utils/test_case_harness"
)

func testParseStrings(stageHarness *test_case_harness.TestCaseHarness) error {
	b := interpreter_executable.NewInterpreterExecutable(stageHarness)

	logger := stageHarness.Logger

	stringLiteral1 := "\"" + strings.Join(random.RandomElementsFromArray(STRINGS, 2), " ") + "\""
	stringLiteral2 := "\"'" + random.RandomElementFromArray(STRINGS) + "'\""
	stringLiteral3 := "\"// " + random.RandomElementFromArray(STRINGS) + "\""
	stringLiteral4 := "\"" + fmt.Sprint(random.RandomInt(10, 100)) + "\""

	parseTestCase := testcases.MultiTestCase{
		TestCases: []testcases.TestCase{
			&testcases.ParseTestCase{FileContents: stringLiteral1, ExpectsError: false},
			&testcases.ParseTestCase{FileContents: stringLiteral2, ExpectsError: false},
			&testcases.ParseTestCase{FileContents: stringLiteral3, ExpectsError: false},
			&testcases.ParseTestCase{FileContents: stringLiteral4, ExpectsError: false},
		},
	}
	return parseTestCase.RunAll(b, logger)
}
