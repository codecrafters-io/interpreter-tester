package testcases

import (
	"fmt"
	"os"

	"github.com/codecrafters-io/interpreter-tester/internal/assertions"
	"github.com/codecrafters-io/interpreter-tester/internal/interpreter_executable"
	"github.com/codecrafters-io/interpreter-tester/internal/lox"
	"github.com/codecrafters-io/tester-utils/logger"
)

// ParseTestCase is a test case for testing the
// "parse" functionality of the interpreter executable.
// A temporary file is created with R.FileContents,
// It is sent to the "parse" command of the executable,
// the expected outputs are generated using the lox.Parse function,
// With that the output of the executable is matched.
type ParseTestCase struct {
	FileContents string
}

func (t *ParseTestCase) Run(executable *interpreter_executable.InterpreterExecutable, logger *logger.Logger) error {
	tmpFileName, err := createTempFileWithContents(t.FileContents)
	if err != nil {
		return err
	}
	defer os.Remove(tmpFileName)

	logReadableFileContents(logger, t.FileContents)

	result, err := executable.Run("parse", tmpFileName)
	if err != nil {
		return err
	}

	expectedStdout, exitCode, _ := lox.Parse(t.FileContents)
	if result.ExitCode != exitCode {
		return fmt.Errorf("expected exit code %v, got %v", exitCode, result.ExitCode)
	}

	// We are intentionally not testing the errors lines printed to stderr
	// We will just check the exitCode here

	expectedStdoutLines := []string{expectedStdout}
	if err = assertions.NewStdoutAssertion(expectedStdoutLines).Run(result, logger); err != nil {
		return err
	}

	logger.Successf("✓ Received exit code %d.", exitCode)

	return nil
}
