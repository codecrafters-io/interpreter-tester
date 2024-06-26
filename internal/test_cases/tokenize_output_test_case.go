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

	logger.Infof("Writing contents to ./test.lox:")
	logger.UpdateSecondaryPrefix("test.lox")
	logger.Infof(t.FileContents)
	logger.ResetSecondaryPrefix()

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
	logCount := len(stdoutAssertionResult)
	if err != nil {
		// If there is an error, the last line should be error log
		// All lines before that should be success logs
		for _, line := range stdoutAssertionResult[:logCount-1] {
			logger.Successf(line)
		}
		logger.Errorf(stdoutAssertionResult[logCount-1])
		return err
	}

	for _, line := range stdoutAssertionResult {
		logger.Successf(line)
	}

	if len(expectedStderr) > 0 {
		stderrAssertionResult, err := assertions.NewOrderedStringArrayAssertion(expectedStderr, "stderr").Run(stdErr)
		logCount = len(stderrAssertionResult)
		if err != nil {
			// If there is an error, the last line should be error log
			// All lines before that should be success logs
			for _, line := range stderrAssertionResult[:logCount-1] {
				logger.Successf(line)
			}
			logger.Errorf(stderrAssertionResult[logCount-1])
			return err
		}
		for _, line := range stderrAssertionResult {
			logger.Successf(line)
		}
	}

	logger.Successf("✓ Received exit code %d.", 0)

	return nil
}
