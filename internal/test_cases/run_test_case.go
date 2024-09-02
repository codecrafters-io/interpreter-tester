package testcases

import (
	"bytes"
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
	FrontMatter  RunTestCaseFrontMatter
}

type RunTestCaseFrontMatter struct {
	ExpectedErrorType string `yaml:"expected_error_type"`
}

func NewRunTestCaseFromFileContents(fileContents []byte, filePath string) RunTestCase {
	if !bytes.HasPrefix(fileContents, []byte("---\n")) {
		panic(fmt.Sprintf("CodeCrafters Internal Error: %s has malformed frontmatter: no beginning triple dashes", filePath))
	}

	fileContents = fileContents[4:]
	endingTripleDashIndex := bytes.Index(fileContents, []byte("\n---\n"))
	if endingTripleDashIndex == -1 {
		panic(fmt.Sprintf("CodeCrafters Internal Error: %s has malformed frontmatter: no ending triple dashes", filePath))
	}

	frontMatterRaw := fileContents[:endingTripleDashIndex]
	fileContents = fileContents[endingTripleDashIndex+5:]

	var frontMatter RunTestCaseFrontMatter
	err := yaml.Unmarshal(frontMatterRaw, &frontMatter)
	if err != nil {
		panic(fmt.Sprintf("CodeCrafters Internal Error: %s has malformed frontmatter: can't unmarshal", filePath))
	}
	if frontMatter.ExpectedErrorType == "" {
		panic(fmt.Sprintf("CodeCrafters Internal Error: %s has malformed frontmatter: missing expected_error_type field", filePath))
	}
	if !(frontMatter.ExpectedErrorType == "none" ||
		frontMatter.ExpectedErrorType == "compile" ||
		frontMatter.ExpectedErrorType == "runtime") {
		panic(fmt.Sprintf("CodeCrafters Internal Error: %s has malformed frontmatter field: expected_error_type shouldn't be %s", filePath, frontMatter.ExpectedErrorType))
	}

	return RunTestCase{
		FileContents: string(fileContents),
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

	expectedStdout, expectedExitCode, _ := loxapi.Run(t.FileContents)

	if t.FrontMatter.ExpectedErrorType == "none" && expectedExitCode != 0 {
		return fmt.Errorf("CodeCrafters internal error: faulty test case, expected this test case to not raise an error, but it did")
	}
	if t.FrontMatter.ExpectedErrorType == "compile" && expectedExitCode != 65 {
		return fmt.Errorf("CodeCrafters internal error: faulty test case, expected this test case to raise a compile time error, but it didn't")
	}
	if t.FrontMatter.ExpectedErrorType == "runtime" && expectedExitCode != 70 {
		return fmt.Errorf("CodeCrafters internal error: faulty test case, expected this test case to raise a runtime error, but it didn't")
	}
	if result.ExitCode != expectedExitCode {
		return fmt.Errorf("expected exit code %v, got %v", expectedExitCode, result.ExitCode)
	}
	// XXX: Stderr is not checked

	expectedStdoutLines := strings.Split(expectedStdout, "\n")
	if err = assertions.NewStdoutAssertion(expectedStdoutLines).Run(result, logger); err != nil {
		return err
	}

	logger.Successf("✓ Received exit code %d.", expectedExitCode)

	return nil
}
