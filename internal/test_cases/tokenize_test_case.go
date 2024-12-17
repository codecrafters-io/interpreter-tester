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
type TokenizeTestCase struct {
	FileContents string
	ExpectsError bool
}

func (t *TokenizeTestCase) Run(executable *interpreter_executable.InterpreterExecutable, logger *logger.Logger) error {
	tmpFileName, err := createTempFileWithContents(t.FileContents)
	if err != nil {
		return err
	}
	defer os.Remove(tmpFileName)

	logReadableFileContents(logger, t.FileContents)

	result, err := executable.Run("tokenize", tmpFileName)
	if err != nil {
		return err
	}

	expectedStdout, _, exitCode, err := loxapi.ScanTokens(t.FileContents)
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

	if err = assertions.NewStdoutAssertion(expectedStdout).Run(result, logger); err != nil {
		return err
	}

	logger.Successf("✓ Received exit code %d.", exitCode)

	return nil
}
