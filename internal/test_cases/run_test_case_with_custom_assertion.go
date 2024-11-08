package testcases

import (
	"fmt"
	"os"
	"slices"

	"github.com/codecrafters-io/interpreter-tester/internal/assertions"
	"github.com/codecrafters-io/interpreter-tester/internal/interpreter_executable"
	"github.com/codecrafters-io/tester-utils/logger"
)

// CustomAssertionRunTestCase is a test case for testing the
// "run" functionality of the interpreter executable.
// A temporary file is created with R.FileContents,
// It is sent to the "run" command of the executable,
// the expected outputs are to be passed as ExpectedResult,
// Assertion is the custom assertion to be used to validate the output.
type CustomAssertionRunTestCase struct {
	FileContents   string
	FrontMatter    RunTestCaseFrontMatter
	Assertion      assertions.Assertion
	ExpectedResult GoloxExecutionResult
}

type GoloxExecutionResult struct {
	Stdout   string
	Stderr   string
	Exitcode int
}

func (t *CustomAssertionRunTestCase) Run(executable *interpreter_executable.InterpreterExecutable, logger *logger.Logger) error {
	tmpFileName, err := createTempFileWithContents(t.FileContents)
	if err != nil {
		return err
	}
	defer os.Remove(tmpFileName)

	logReadableFileContents(logger, t.FileContents)

	result, err := executable.Run("run", tmpFileName)
	if err != nil {
		return err
	}

	if err := validateExpectedErrorTypeAndExitCode(t.FrontMatter.ExpectedErrorType, t.ExpectedResult.Exitcode, result.ExitCode); err != nil {
		return err
	}

	if err = t.Assertion.Run(result, logger); err != nil {
		return err
	}

	logger.Successf("✓ Received exit code %d.", t.ExpectedResult.Exitcode)

	return nil
}

func validateExpectedErrorTypeAndExitCode(expectedErrorType string, expectedExitCode int, actualExitCode int) error {
	if !slices.Contains([]string{"none", "compile", "runtime"}, expectedErrorType) {
		return fmt.Errorf("CodeCrafters internal error: Unknown expected_error_type value: %s", expectedErrorType)
	}
	if expectedErrorType == "none" && expectedExitCode != 0 {
		return fmt.Errorf("CodeCrafters internal error: faulty test case, expected this test case to not raise an error, but it did")
	}
	if expectedErrorType == "compile" && expectedExitCode != 65 {
		return fmt.Errorf("CodeCrafters internal error: faulty test case, expected this test case to raise a compile time error, but it didn't")
	}
	if expectedErrorType == "runtime" && expectedExitCode != 70 {
		return fmt.Errorf("CodeCrafters internal error: faulty test case, expected this test case to raise a runtime error, but it didn't")
	}
	if actualExitCode != expectedExitCode {
		return fmt.Errorf("expected exit code %v, got %v", expectedExitCode, actualExitCode)
	}
	return nil
}
