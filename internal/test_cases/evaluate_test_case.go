package testcases

import (
	"fmt"
	"os"

	"github.com/codecrafters-io/interpreter-tester/internal/assertions"
	"github.com/codecrafters-io/interpreter-tester/internal/interpreter_executable"
	"github.com/codecrafters-io/interpreter-tester/internal/lox"
	"github.com/codecrafters-io/tester-utils/logger"
)

// EvaluateTestCase is a test case for testing the
// "evaluate" functionality of the interpreter executable.
// A temporary file is created with R.FileContents,
// It is sent to the "evaluate" command of the executable,
// the expected outputs are generated using the lox.Evaluate function,
// With that the output of the executable is matched.
type EvaluateTestCase struct {
	FileContents string
}

func (t *EvaluateTestCase) Run(executable *interpreter_executable.InterpreterExecutable, logger *logger.Logger) error {
	tmpFileName, err := createTempFileWithContents(t.FileContents)
	if err != nil {
		return err
	}
	defer os.Remove(tmpFileName)

	logReadableFileContents(logger, t.FileContents)

	result, err := executable.Run("evaluate", tmpFileName)
	if err != nil {
		return err
	}

	expectedStdout, exitCode, _ := lox.Evaluate(t.FileContents)
	if result.ExitCode != exitCode {
		return fmt.Errorf("expected exit code %v, got %v", exitCode, result.ExitCode)
	}

	// ToDo: Confirm
	// We are intentionally not testing the errors lines printed to stderr
	// We will just check the exitCode here

	expectedStdoutLines := []string{expectedStdout}
	if err = assertions.NewStdoutAssertion(expectedStdoutLines).Run(result, logger); err != nil {
		return err
	}

	logger.Successf("✓ Received exit code %d.", exitCode)

	return nil
}
