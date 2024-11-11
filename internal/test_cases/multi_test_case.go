package testcases

import (
	"fmt"

	"github.com/codecrafters-io/interpreter-tester/internal/interpreter_executable"
	"github.com/codecrafters-io/tester-utils/logger"
)

type MultiTestCase struct {
	TestCases []TestCase
}

func (t *MultiTestCase) RunAll(executable *interpreter_executable.InterpreterExecutable, logger *logger.Logger) error {
	for i, testCase := range t.TestCases {
		logger.UpdateSecondaryPrefix(fmt.Sprintf("test-%d", i+1))
		logger.Infof("Running test case: %d", i+1)
		if err := testCase.Run(executable, logger); err != nil {
			return err
		}
		// If err is encountered, we want the secondary prefix to be present
		logger.ResetSecondaryPrefix()
	}
	return nil
}
