package testcases

import (
	"fmt"
	"os"
	"strings"

	"github.com/codecrafters-io/interpreter-tester/internal/assertions"
	"github.com/codecrafters-io/interpreter-tester/internal/interpreter_executable"
	loxapi "github.com/codecrafters-io/interpreter-tester/internal/lox/api"
	"github.com/codecrafters-io/tester-utils/logger"
)

// RunTestCase is a test case for testing the
// "run" functionality of the interpreter executable.
// A temporary file is created with R.FileContents,
// It is sent to the "run" command of the executable,
// the expected outputs are generated using the lox.Run function,
// With that the output of the executable is matched.
type RunTestCase struct {
	FileContents string
}

func (t *RunTestCase) Run(executable *interpreter_executable.InterpreterExecutable, logger *logger.Logger) error {
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

	expectedStdout, exitCode, _ := loxapi.Run(t.FileContents)
	if result.ExitCode != exitCode {
		return fmt.Errorf("expected exit code %v, got %v", exitCode, result.ExitCode)
	}
	// XXX: Stderr is not checked

	expectedStdoutLines := strings.Split(expectedStdout, "\n")
	if err = assertions.NewStdoutAssertion(expectedStdoutLines).Run(result, logger); err != nil {
		return err
	}

	logger.Successf("✓ Received exit code %d.", exitCode)

	return nil
}
