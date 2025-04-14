package testcases

import (
	"fmt"
	"os"

	"github.com/codecrafters-io/interpreter-tester/internal/assertions"
	"github.com/codecrafters-io/interpreter-tester/internal/interpreter_executable"
	loxapi "github.com/codecrafters-io/interpreter-tester/internal/lox/api"
	"github.com/codecrafters-io/tester-utils/logger"
)

// TokenizeTestCase is a test case for testing the
// "tokenize" functionality of the interpreter executable.
// A temporary file is created with R.FileContents,
// It is sent to the "tokenize" command of the executable,
// the expected outputs are generated using the lox.ScanTokens function,
// With that the output of the executable is matched.
// Default value for HighlightWhitespaceWhileLogging is false.
// We explicitly set it to true in the tests involving whitespace.
type TokenizeTestCase struct {
	FileContents                    string
	ExpectsError                    bool
	HighlightWhitespaceWhileLogging bool
}

func (t *TokenizeTestCase) Run(executable *interpreter_executable.InterpreterExecutable, logger *logger.Logger) error {
	tmpFileName, err := createTempFileWithContents(t.FileContents)
	if err != nil {
		return err
	}
	defer os.Remove(tmpFileName)

	if t.HighlightWhitespaceWhileLogging {
		logReadableFileContentsPreservingWhitespace(logger, t.FileContents)
	} else {
		logReadableFileContents(logger, t.FileContents)
	}

	result, err := executable.Run("tokenize", tmpFileName)
	if err != nil {
		return err
	}

	expectedStdout, expectedStderr, exitCode, err := loxapi.ScanTokens(t.FileContents)
	if err != nil {
		return fmt.Errorf("CodeCrafters internal error: %v", err)
	}

	if t.ExpectsError && exitCode == 0 {
		return fmt.Errorf("CodeCrafters internal error: faulty test case, expected this test case to raise an error, but it didn't")
	}
	if !t.ExpectsError && exitCode != 0 {
		return fmt.Errorf("CodeCrafters internal error: faulty test case, expected this test case to not raise an error, but it did")
	}

	if result.ExitCode != exitCode {
		return fmt.Errorf("expected %v (exit code %v), got exit code %v", exitCodeToErrorTypeMapping[exitCode], exitCode, result.ExitCode)
	}

	if len(expectedStderr) > 0 {
		if err := assertions.NewStderrAssertion(expectedStderr).Run(result, logger); err != nil {
			return err
		}
	}

	if err = assertions.NewStdoutAssertion(expectedStdout).Run(result, logger); err != nil {
		return err
	}

	logger.Successf("✓ Received exit code %d.", exitCode)

	return nil
}
