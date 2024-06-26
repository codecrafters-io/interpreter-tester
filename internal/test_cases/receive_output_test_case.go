package testcases

import (
	"fmt"
	"os"

	"github.com/codecrafters-io/grep-tester/internal/assertions"
	"github.com/codecrafters-io/grep-tester/internal/interpreter_executable"
	"github.com/codecrafters-io/tester-utils/logger"
)

// ToDo: Improve test case name

// ReceiveOutputTestCase is a test case for testing the
// command functionality of the interpreter executable.
// A tmp file is created with R.FileContents,
// and the output of the command is compared with the expected outputs,
// For the expected output, the expectedStdout and expectedStderr are used.
type ReceiveOutputTestCase struct {
	FileContents   string
	Command        string
	ExpectedStdout []string
	ExpectedStderr []string
}

func (t *ReceiveOutputTestCase) Run(executable *interpreter_executable.InterpreterExecutable, logger *logger.Logger) error {
	tmpFileName, err := createTempFileWithContents(t.FileContents)
	if err != nil {
		return err
	}
	defer os.Remove(tmpFileName)

	logger.Infof("Writing contents to ./test.lox:")
	logger.UpdateSecondaryPrefix("test.lox")
	logger.Infof(t.FileContents)
	logger.ResetSecondaryPrefix()

	result, err := executable.Run(t.Command, tmpFileName)
	if err != nil {
		return err
	}

	// No matter what, we always expect the exit code to be 0
	if result.ExitCode != 0 {
		return fmt.Errorf("expected exit code %v, got %v", 0, result.ExitCode)
	}

	stdOut := getOutputFromStdOut(result)
	stdErr := getOutputFromStdErr(result)

	stdoutAssertionResult, err := assertions.NewOrderedStringArrayAssertion(t.ExpectedStdout, "stdout").Run(stdOut)
	for _, line := range stdoutAssertionResult {
		logger.Successf(line)
	}
	if err != nil {
		return err
	}

	stderrAssertionResult, err := assertions.NewOrderedStringArrayAssertion(t.ExpectedStderr, "stderr").Run(stdErr)
	for _, line := range stderrAssertionResult {
		logger.Successf(line)
	}
	if err != nil {
		return err
	}

	logger.Successf("✓ Received exit code %d.", 0)

	return nil
}
