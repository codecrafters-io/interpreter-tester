package interpreter_executable

import (
	"fmt"
	"path"

	"github.com/codecrafters-io/tester-utils/executable"
	"github.com/codecrafters-io/tester-utils/logger"
	"github.com/codecrafters-io/tester-utils/test_case_harness"
)

type InterpreterExecutable struct {
	executable *executable.Executable
	logger     *logger.Logger
	args       []string
}

func NewInterpreterExecutable(stageHarness *test_case_harness.TestCaseHarness) *InterpreterExecutable {
	b := &InterpreterExecutable{
		executable: stageHarness.NewExecutable(),
		logger:     stageHarness.Logger,
	}

	return b
}

func (b *InterpreterExecutable) Run(args ...string) (executable.ExecutableResult, error) {
	b.args = args
	var log string
	log += fmt.Sprintf("$ ./%s", path.Base(b.executable.Path))
	for _, arg := range b.args {
		log += " " + arg
	}
	b.logger.Infof(log)

	return b.executable.Run(b.args...)
}
