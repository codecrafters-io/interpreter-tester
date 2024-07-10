package testcases

import (
	"fmt"

	"github.com/codecrafters-io/interpreter-tester/internal/interpreter_executable"
	"github.com/codecrafters-io/tester-utils/logger"
)

type MultiParseTestCase struct {
	FileContents []string
}

func (t *MultiParseTestCase) RunAll(executable *interpreter_executable.InterpreterExecutable, logger *logger.Logger) error {
	for i, fileContents := range t.FileContents {
		logger.UpdateSecondaryPrefix(fmt.Sprintf("test-%d", i+1))
		logger.Infof("Running test case: %d", i+1)
		testCase := ParseTestCase{
			FileContents: fileContents}
		if err := testCase.Run(executable, logger); err != nil {
			return err
		}
		// If err is encountered, we want the secondary prefix to be present
		logger.ResetSecondaryPrefix()
	}
	return nil
}
