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
	// FileContents is the contents of the .lox file to run
	FileContents string

	// ExpectedExitCode is the expected exit code after running the file
	ExpectedExitCode int

	// OutputAssertion is the assertion to check the output of the executable
	OutputAssertion assertions.Assertion
}

type RunTestCaseFrontMatter struct {
	ExpectedErrorType string `yaml:"expected_error_type"`
}

func NewRunTestCaseFromFilePath(filePath string) RunTestCase {
	fileContents, err := os.ReadFile(filePath)
	if err != nil {
		panic(fmt.Sprintf("CodeCrafters Internal Error: Encountered error while reading test file: %s", err))
	}

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
	err = yaml.Unmarshal(frontMatterRaw, &frontMatter)
	if err != nil {
		panic(fmt.Sprintf("CodeCrafters Internal Error: %s has malformed frontmatter: can't unmarshal", filePath))
	}

	if frontMatter.ExpectedErrorType == "" {
		panic(fmt.Sprintf("CodeCrafters Internal Error: %s has malformed frontmatter: missing expected_error_type field", filePath))
	}

	if frontMatter.ExpectedErrorType != "none" &&
		frontMatter.ExpectedErrorType != "compile" &&
		frontMatter.ExpectedErrorType != "runtime" {
		panic(fmt.Sprintf("CodeCrafters Internal Error: %s has malformed frontmatter field: expected_error_type shouldn't be %s", filePath, frontMatter.ExpectedErrorType))
	}

	expectedExitCode := map[string]int{
		"none":    0,
		"compile": 65,
		"runtime": 70,
	}[frontMatter.ExpectedErrorType]

	return RunTestCase{
		FileContents:     string(fileContents),
		ExpectedExitCode: expectedExitCode,
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

	ourLoxStdout, ourLoxExitCode, _ := loxapi.Run(t.FileContents)

	if t.ExpectedExitCode != ourLoxExitCode {
		return fmt.Errorf("CodeCrafters internal error: faulty test case, expected %d exit code, our lox returned %d", t.ExpectedExitCode, ourLoxExitCode)
	}
	if result.ExitCode != t.ExpectedExitCode {
		return fmt.Errorf("expected %v (exit code %v), got exit code %v", exitCodeToErrorTypeMapping[t.ExpectedExitCode], t.ExpectedExitCode, result.ExitCode)
	}

	if t.OutputAssertion == nil {
		expectedStdoutLines := strings.Split(ourLoxStdout, "\n")
		t.OutputAssertion = assertions.NewStdoutAssertion(expectedStdoutLines)
	}

	if err = t.OutputAssertion.Run(result, logger); err != nil {
		return err
	}

	logger.Successf("✓ Received exit code %d.", t.ExpectedExitCode)

	return nil
}
