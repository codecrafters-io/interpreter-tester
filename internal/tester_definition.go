package internal

import (
	"github.com/codecrafters-io/tester-utils/tester_definition"
)

var testerDefinition = tester_definition.TesterDefinition{
	AntiCheatTestCases:       []tester_definition.TestCase{},
	ExecutableFileName:       "your_program.sh",
	LegacyExecutableFileName: "your_program.sh",
	TestCases: []tester_definition.TestCase{
		{
			Slug:     "ry8",
			TestFunc: testEOF,
		},
		{
			Slug:     "ol4",
			TestFunc: testParen,
		},
		{
			Slug:     "oe8",
			TestFunc: testBrace,
		},
		{
			Slug:     "xc5",
			TestFunc: testSingleChars,
		},
	},
}
