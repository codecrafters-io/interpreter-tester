package testcases

import (
	"github.com/codecrafters-io/interpreter-tester/internal/interpreter_executable"
	"github.com/codecrafters-io/tester-utils/logger"
)

type TestCase interface {
	Run(executable *interpreter_executable.InterpreterExecutable, logger *logger.Logger) error
}
