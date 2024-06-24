package internal

import (
	"fmt"
	"os"

	"github.com/codecrafters-io/grep-tester/internal/assertions"
	"github.com/codecrafters-io/grep-tester/internal/interpreter_executable"

	"github.com/codecrafters-io/tester-utils/test_case_harness"
)

func testInit(stageHarness *test_case_harness.TestCaseHarness) error {
	b := interpreter_executable.NewInterpreterExecutable(stageHarness)

	logger := stageHarness.Logger

	tmpFileName, err := createTempFileWithContents("(())##")
	if err != nil {
		return err
	}
	defer os.Remove(tmpFileName)

	result, err := b.Run("tokenize", tmpFileName)
	if err != nil {
		return err
	}

	if result.ExitCode != 0 {
		return fmt.Errorf("expected exit code %v, got %v", 0, result.ExitCode)
	}

	stdOut := getOutputFromStdOut(result)
	stdErr := getOutputFromStdErr(result)

	expectedStdout := []string{"LEFT_PAREN ( null", "LEFT_PAREN ( null", "RIGHT_PAREN ) null", "RIGHT_PAREN ) null", "EOF  null"}

	expectedStderr := []string{"[line 1] Error : Unexpected character: #", "[line 1] Error : Unexpected character: #"}

	assertionResult, err := assertions.NewOrderedStringArrayAssertion(expectedStdout, "stdout").Run(stdOut)
	for _, line := range assertionResult {
		logger.Successf(line)
	}
	if err != nil {
		return err
	}

	assertionResult, err = assertions.NewOrderedStringArrayAssertion(expectedStderr, "stderr").Run(stdErr)
	for _, line := range assertionResult {
		logger.Successf(line)
	}
	if err != nil {
		return err
	}

	logger.Successf("✓ Received exit code %d.", 0)
	return nil
}
