package internal

import (
	"github.com/codecrafters-io/grep-tester/internal/interpreter_executable"
	testcases "github.com/codecrafters-io/grep-tester/internal/test_cases"

	"github.com/codecrafters-io/tester-utils/test_case_harness"
)

func testInit(stageHarness *test_case_harness.TestCaseHarness) error {
	b := interpreter_executable.NewInterpreterExecutable(stageHarness)

	logger := stageHarness.Logger

	logger.Infof("\nRunning test: 1")
	commandTestCase1 := testcases.ReceiveOutputTestCase{
		FileContents:   "(",
		Command:        "tokenize",
		ExpectedStdout: []string{"LEFT_PAREN ( null", "EOF  null"},
		ExpectedStderr: []string{},
	}
	if err := commandTestCase1.Run(b, logger); err != nil {
		return err
	}

	logger.Infof("\nRunning test: 2")
	commandTestCase2 := testcases.ReceiveOutputTestCase{
		FileContents:   "))",
		Command:        "tokenize",
		ExpectedStdout: []string{"RIGHT_PAREN ) null", "RIGHT_PAREN ) null", "EOF  null"},
		ExpectedStderr: []string{},
	}
	if err := commandTestCase2.Run(b, logger); err != nil {
		return err
	}

	logger.Infof("\nRunning test: 3")
	commandTestCase3 := testcases.ReceiveOutputTestCase{
		FileContents:   "(())",
		Command:        "tokenize",
		ExpectedStdout: []string{"LEFT_PAREN ( null", "LEFT_PAREN ( null", "RIGHT_PAREN ) null", "RIGHT_PAREN ) null", "EOF  null"},
		ExpectedStderr: []string{},
	}
	if err := commandTestCase3.Run(b, logger); err != nil {
		return err
	}

	logger.Infof("\nRunning test: 4")
	commandTestCase4 := testcases.ReceiveOutputTestCase{
		FileContents:   ")(())()(())(",
		Command:        "tokenize",
		ExpectedStdout: []string{"RIGHT_PAREN ) null", "LEFT_PAREN ( null", "LEFT_PAREN ( null", "RIGHT_PAREN ) null", "RIGHT_PAREN ) null", "LEFT_PAREN ( null", "RIGHT_PAREN ) null", "LEFT_PAREN ( null", "LEFT_PAREN ( null", "RIGHT_PAREN ) null", "RIGHT_PAREN ) null", "LEFT_PAREN ( null", "EOF  null"},
		ExpectedStderr: []string{},
	}
	if err := commandTestCase4.Run(b, logger); err != nil {
		return err
	}

	// ToDo: No need to test errors here
	logger.Infof("\nRunning test: X")
	commandTestCaseX := testcases.ReceiveOutputTestCase{
		FileContents:   "(())##",
		Command:        "tokenize",
		ExpectedStdout: []string{"LEFT_PAREN ( null", "LEFT_PAREN ( null", "RIGHT_PAREN ) null", "RIGHT_PAREN ) null", "EOF  null"},
		ExpectedStderr: []string{"[line 1] Error: Unexpected character: #", "[line 1] Error: Unexpected character: #"},
	}
	if err := commandTestCaseX.Run(b, logger); err != nil {
		return err
	}

	return nil
}
