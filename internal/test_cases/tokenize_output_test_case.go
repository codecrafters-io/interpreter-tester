package testcases

import (
	"fmt"
	"os"

	"github.com/codecrafters-io/grep-tester/internal/assertions"
	"github.com/codecrafters-io/grep-tester/internal/interpreter_executable"
	"github.com/codecrafters-io/grep-tester/internal/lox"
	"github.com/codecrafters-io/tester-utils/logger"
)

// ToDo: Improve test case name

// TokenizeOutputTestCase is a test case for testing the
// "tokenize" functionality of the interpreter executable.
// A tmp file is created with R.FileContents,
// It is sent to the "tokenize" command of the executable,
// the expected outputs are generated using the lox.ScanTokens function,
// With that the output of the executable is matched.
type TokenizeOutputTestCase struct {
	FileContents string
}

func (t *TokenizeOutputTestCase) Run(executable *interpreter_executable.InterpreterExecutable, logger *logger.Logger) error {
	tmpFileName, err := createTempFileWithContents(t.FileContents)
	if err != nil {
		return err
	}
	defer os.Remove(tmpFileName)

	logger.Debugf("File contents: %q", t.FileContents)
	result, err := executable.Run("tokenize", tmpFileName)
	if err != nil {
		return err
	}

	// No matter what, we always expect the exit code to be 0
	if result.ExitCode != 0 {
		return fmt.Errorf("expected exit code %v, got %v", 0, result.ExitCode)
	}

	stdOut := getOutputFromStdOut(result)
	stdErr := getOutputFromStdErr(result)

	expectedStdout, expectedStderr, err := lox.ScanTokens(t.FileContents)
	if err != nil {
		return fmt.Errorf("CodeCrafters internal error: %v", err)
	}

	stdoutAssertionResult, err := assertions.NewOrderedStringArrayAssertion(expectedStdout, "stdout").Run(stdOut)
	for _, line := range stdoutAssertionResult {
		logger.Successf(line)
	}
	if err != nil {
		return err
	}

	if len(expectedStderr) > 0 {
		stderrAssertionResult, err := assertions.NewOrderedStringArrayAssertion(expectedStderr, "stderr").Run(stdErr)
		for _, line := range stderrAssertionResult {
			logger.Successf(line)
		}
		if err != nil {
			return err
		}
	}

	logger.Successf("✓ Received exit code %d.", 0)

	return nil
}
