package internal

import (
	"github.com/codecrafters-io/interpreter-tester/internal/interpreter_executable"
	testcases "github.com/codecrafters-io/interpreter-tester/internal/test_cases"

	"github.com/codecrafters-io/tester-utils/test_case_harness"
)

func testParseBooleans(stageHarness *test_case_harness.TestCaseHarness) error {
	b := interpreter_executable.NewInterpreterExecutable(stageHarness)

	logger := stageHarness.Logger

	parseTestCase := testcases.MultiTestCase{
		TestCases: []testcases.TestCase{
			&testcases.ParseTestCase{FileContents: "true", ExpectsError: false},
			&testcases.ParseTestCase{FileContents: "false", ExpectsError: false},
			&testcases.ParseTestCase{FileContents: "nil", ExpectsError: false},
		},
	}
	return parseTestCase.RunAll(b, logger)
}
