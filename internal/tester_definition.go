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
		{
			Slug:     "ea6",
			TestFunc: testErrors,
		},
		{
			Slug:     "mp7",
			TestFunc: testEquality,
		},
		{
			Slug:     "bu3",
			TestFunc: testNegation,
		},
		{
			Slug:     "et2",
			TestFunc: testRelational,
		},
		{
			Slug:     "ml2",
			TestFunc: testComments,
		},
		{
			Slug:     "er2",
			TestFunc: testWhitespace,
		},
		{
			Slug:     "tz7",
			TestFunc: testErrorsMulti,
		},
		{
			Slug:     "ue7",
			TestFunc: testStrings,
		},
		{
			Slug:     "kj0",
			TestFunc: testNumbers,
		},
		{
			Slug:     "ey7",
			TestFunc: testIdentifier,
		},
		{
			Slug:     "pq5",
			TestFunc: testReservedWords,
		},
		{
			Slug:     "sc2",
			TestFunc: testParseBooleans,
		},
		{
			Slug:     "ra8",
			TestFunc: testParseNumbers,
		},
		{
			Slug:     "th5",
			TestFunc: testParseStrings,
		},
		{
			Slug:     "xe6",
			TestFunc: testParseParens,
		},
		{
			Slug:     "mq1",
			TestFunc: testParseUnary,
		},
		{
			Slug:     "wa9",
			TestFunc: testParseFactor,
		},
		{
			Slug:     "yf2",
			TestFunc: testParseTerms,
		},
		{
			Slug:     "uh4",
			TestFunc: testParseComparison,
		},
		{
			Slug:     "ht8",
			TestFunc: testParseEquality,
		},
		{
			Slug:     "wz8",
			TestFunc: testParseErrors,
		},
	},
}
