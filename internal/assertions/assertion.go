package assertions

import (
	"github.com/codecrafters-io/tester-utils/executable"
	"github.com/codecrafters-io/tester-utils/logger"
)

type Assertion interface {
	Run(result executable.ExecutableResult, logger *logger.Logger) error
}
