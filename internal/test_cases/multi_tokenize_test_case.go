package testcases

import (
	"github.com/codecrafters-io/interpreter-tester/internal/interpreter_executable"
	"github.com/codecrafters-io/tester-utils/logger"
)

type MultiTokenizeTestCase struct {
	FileContents []string
}

func (t *MultiTokenizeTestCase) RunAll(executable *interpreter_executable.InterpreterExecutable, logger *logger.Logger) error {
	for i, fileContents := range t.FileContents {
		logger.Infof("Running test case: %d", i+1)
		testCase := TokenizeTestCase{
			FileContents: fileContents}
		if err := testCase.Run(executable, logger); err != nil {
			return err
		}
		logger.Infof("Test case %d passed\n", i+1)
	}
	return nil
}
