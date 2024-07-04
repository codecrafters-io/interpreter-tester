package testcases

import (
	"fmt"
	"os"
	"regexp"
	"strings"

	"github.com/codecrafters-io/interpreter-tester/internal/assertions"
	"github.com/codecrafters-io/interpreter-tester/internal/interpreter_executable"
	"github.com/codecrafters-io/interpreter-tester/internal/lox"
	"github.com/codecrafters-io/tester-utils/logger"
)

// TokenizeTestCase is a test case for testing the
// "tokenize" functionality of the interpreter executable.
// A temporary file is created with R.FileContents,
// It is sent to the "tokenize" command of the executable,
// the expected outputs are generated using the lox.ScanTokens function,
// With that the output of the executable is matched.
type TokenizeTestCase struct {
	FileContents string
}

func (t *TokenizeTestCase) Run(executable *interpreter_executable.InterpreterExecutable, logger *logger.Logger) error {
	tmpFileName, err := createTempFileWithContents(t.FileContents)
	if err != nil {
		return err
	}
	defer os.Remove(tmpFileName)

	logger.Infof("Writing contents to ./test.lox:")
	existingSecondaryPrefix := logger.GetSecondaryPrefix()
	// existingSecondaryPrefix can be "" or hold our intended log prefix
	logger.UpdateSecondaryPrefix(existingSecondaryPrefix + "[test.lox] ")

	// If the file contents contain a singe %, it will be decoded as a format specifier
	// And it will add a `(MISSING)` to the log line
	printableFileContents := strings.ReplaceAll(t.FileContents, "%", "%%")
	printableFileContents = strings.ReplaceAll(printableFileContents, "\t", "<|TAB|>")

	regex1 := regexp.MustCompile("[ ]+\n")
	regex2 := regexp.MustCompile("[ ]+$")
	printableFileContents = regex1.ReplaceAllString(printableFileContents, "<|SPACE|>")
	printableFileContents = regex2.ReplaceAllString(printableFileContents, "<|SPACE|>")

	logger.Infof(printableFileContents)

	logger.UpdateSecondaryPrefix(existingSecondaryPrefix)

	result, err := executable.Run("tokenize", tmpFileName)
	if err != nil {
		return err
	}

	stdout := getStdoutLinesFromExecutableResult(result)
	stderr := getStderrLinesFromExecutableResult(result)

	expectedStdout, expectedStderr, exitCode, err := lox.ScanTokens(t.FileContents)
	if err != nil {
		return fmt.Errorf("CodeCrafters internal error: %v", err)
	}

	// No matter what, we always expect the exit code to be 0
	if result.ExitCode != exitCode {
		return fmt.Errorf("expected exit code %v, got %v", exitCode, result.ExitCode)
	}

	if len(expectedStderr) > 0 {
		stderrAssertionResult, err := assertions.NewStderrAssertion(expectedStderr).Run(stderr)
		logCount := len(stderrAssertionResult)
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

	stdoutAssertionResult, err := assertions.NewStdoutAssertion(expectedStdout).Run(stdout)
	logCount := len(stdoutAssertionResult)
	if err != nil {
		// If there is an error, the last line should be error log
		// All lines before that should be success logs
		if logCount > 1 {
			for _, line := range stdoutAssertionResult[:logCount-1] {
				logger.Successf(line)
			}
			logger.Errorf(stdoutAssertionResult[logCount-1])
		}
		return err
	}
	for _, line := range stdoutAssertionResult {
		logger.Successf(line)
	}

	logger.Successf("✓ Received exit code %d.", exitCode)

	return nil
}
