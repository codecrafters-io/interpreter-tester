package testcases

import (
	"fmt"
	"os"
	"strings"

	"github.com/codecrafters-io/interpreter-tester/internal/assertions"
	"github.com/codecrafters-io/interpreter-tester/internal/interpreter_executable"
	loxapi "github.com/codecrafters-io/interpreter-tester/internal/lox/api"
	"github.com/codecrafters-io/tester-utils/logger"
	"gopkg.in/yaml.v3"
)

// RunTestCase is a test case for testing the
// "run" functionality of the interpreter executable.
// A temporary file is created with R.FileContents,
// It is sent to the "run" command of the executable,
// the expected outputs are generated using the lox.Run function,
// With that the output of the executable is matched.

type RunTestCase struct {
	FileContents string
	FrontMatter  RunTestCaseFrontMater
}

type RunTestCaseFrontMater struct {
	ExpectedErrorType string `yaml:"error_type"`
}

func NewRunTestCaseFromFileContents(fileContents []byte, filePath string) RunTestCase {
	contents := string(fileContents)

	if contents[:4] != "---\n" {
		panic(fmt.Sprintf("CodeCrafters Internal Error: %s has malformed frontmatter: no beginning triple dashes", filePath))
	}

	contents = contents[4:]
	endingTripleDashIndex := strings.Index(contents, "\n---\n")
	if endingTripleDashIndex == -1 {
		panic(fmt.Sprintf("CodeCrafters Internal Error: %s has malformed frontmatter: no ending triple dashes", filePath))
	}

	frontMatterBytes := []byte(contents[:endingTripleDashIndex])
	contents = contents[endingTripleDashIndex+5:]

	var frontMatter RunTestCaseFrontMater
	err := yaml.Unmarshal(frontMatterBytes, &frontMatter)
	if err != nil {
		panic(fmt.Sprintf("CodeCrafters Internal Error: %s has malformed frontmatter: can't unmarshal", filePath))
	}
	if frontMatter.ExpectedErrorType == "" {
		panic(fmt.Sprintf("CodeCrafters Internal Error: %s has malformed frontmatter: no error_type field", filePath))
	}

	return RunTestCase{
		FileContents: contents,
		FrontMatter:  frontMatter,
	}
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

	if t.FrontMatter.ExpectedErrorType == "none" && exitCode != 0 {
		return fmt.Errorf("CodeCrafters internal error: faulty test case, expected this test case to not raise an error, but it did.")
	}
	if t.FrontMatter.ExpectedErrorType == "compile" && exitCode != 65 {
		return fmt.Errorf("CodeCrafters internal error: faulty test case, expected this test case to raise a compile time error, but it didn't")
	}
	if t.FrontMatter.ExpectedErrorType == "runtime" && exitCode != 70 {
		return fmt.Errorf("CodeCrafters internal error: faulty test case, expected this test case to raise a runtime error, but it didn't")
	}
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
