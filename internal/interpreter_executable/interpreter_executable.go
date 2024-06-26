package interpreter_executable

import (
	"fmt"
	"path"
	"strings"

	executable "github.com/codecrafters-io/tester-utils/executable"
	logger "github.com/codecrafters-io/tester-utils/logger"
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
	if b.args == nil || len(b.args) == 0 {
		b.logger.Infof("$ ./%s", path.Base(b.executable.Path))
	} else {
		var log string
		log += fmt.Sprintf("$ ./%s", path.Base(b.executable.Path))
		for _, arg := range b.args {
			// Instead of printing the full path of the file, we can just print the file name, the path is super long
			// XXX: Confirm with the team if this is the desired behavior
			if strings.Contains(arg, ".lox") {
				log += " " + "test.lox"
			} else {
				log += " " + arg
			}
		}
		b.logger.Infof(log)
	}

	result, err := b.executable.Run(b.args...)
	if err != nil {
		return executable.ExecutableResult{}, err
	}

	return result, nil
}