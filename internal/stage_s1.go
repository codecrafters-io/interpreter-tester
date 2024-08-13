package internal

import (
	"github.com/codecrafters-io/interpreter-tester/internal/interpreter_executable"
	testcases "github.com/codecrafters-io/interpreter-tester/internal/test_cases"
	"github.com/codecrafters-io/tester-utils/test_case_harness"
)

func testStatements1(stageHarness *test_case_harness.TestCaseHarness) error {
	b := interpreter_executable.NewInterpreterExecutable(stageHarness)

	logger := stageHarness.Logger

	fileContents := GetTestProgramsForCurrentStage()
	var replacedFileContents []string
	for _, fileContent := range fileContents {
		replacedFileContent := ReplacePlaceholdersWithRandomValues(fileContent)
		replacedFileContents = append(replacedFileContents, replacedFileContent)
	}

	// fileContents = GetTestProgramsForCurrentStageWithRandomValues()

	runTestCases := testcases.MultiRunTestCase{
		FileContents: replacedFileContents,
	}

	return runTestCases.RunAll(b, logger)
}
