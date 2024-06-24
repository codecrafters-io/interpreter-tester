package internal

import (
	"fmt"
	"os"
	"regexp"
	"strings"

	"github.com/codecrafters-io/tester-utils/test_case_harness"
)

func testInit(stageHarness *test_case_harness.TestCaseHarness) error {
	// should exit with 0: lox tokenize test.lox

	logger := stageHarness.Logger
	logger.Infof("$ your_program.sh tokenize test.lox")

	tmpFile, err := os.CreateTemp("", "test.lox")
	if err != nil {
		return fmt.Errorf("CodeCrafters internal error. Error creating tmp file: %v", err)
	}
	_, err = tmpFile.WriteString("(())##")
	if err != nil {
		return fmt.Errorf("CodeCrafters internal error. Error writing to tmp file: %v", err)
	}
	err = tmpFile.Close()
	if err != nil {
		return fmt.Errorf("CodeCrafters internal error. Error closing tmp file: %v", err)
	}

	tmpFileName := tmpFile.Name()
	defer os.Remove(tmpFileName)

	result, err := stageHarness.Executable.Run("tokenize", tmpFileName)
	if err != nil {
		return err
	}

	if result.ExitCode != 0 {
		return fmt.Errorf("expected exit code %v, got %v", 0, result.ExitCode)
	}

	stdOut := strings.Split(strings.TrimRight(string(result.Stdout), "\n"), "\n")

	filteredStdErr := []string{}
	stdErr := strings.Split(strings.TrimRight(string(result.Stderr), "\n"), "\n")
	regexp := regexp.MustCompile(`\[line [0-9]+\]`)

	for _, line := range stdErr {
		if regexp.MatchString(line) {
			filteredStdErr = append(filteredStdErr, line)
		}
	}

	expectedStdout := []string{"LEFT_PAREN ( null", "LEFT_PAREN ( null", "RIGHT_PAREN ) null", "RIGHT_PAREN ) null", "EOF  null"}

	expectedStderr := []string{"[line 1] Error : Unexpected character: #", "[line 1] Error : Unexpected character: #"}

	if len(stdOut) != len(expectedStdout) {
		return fmt.Errorf("expected %d lines of stdout, got %d", len(expectedStdout), len(stdOut))
	}

	for i, expectedValue := range expectedStdout {
		actualValue := stdOut[i]
		if actualValue != expectedValue {
			return fmt.Errorf("Expected element #%d to be %q, got %q", i+1, expectedValue, actualValue)
		} else {
			logger.Successf("✓ %s", actualValue)
		}
	}

	if len(filteredStdErr) != len(expectedStderr) {
		return fmt.Errorf("expected %d lines of stderr, got %d", len(expectedStderr), len(filteredStdErr))
	}

	for i, expectedValue := range expectedStderr {
		actualValue := filteredStdErr[i]
		if actualValue != expectedValue {
			return fmt.Errorf("Expected error #%d to be %q, got %q", i+1, expectedValue, actualValue)
		} else {
			logger.Successf("✓ %s", actualValue)
		}
	}

	logger.Successf("✓ Received exit code %d.", 0)
	return nil
}
