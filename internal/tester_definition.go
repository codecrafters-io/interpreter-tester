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
		{
			Slug:     "iz6",
			TestFunc: testEvaluateBooleans,
		},
		{
			Slug:     "lv1",
			TestFunc: testEvaluateLiterals,
		},
		{
			Slug:     "oq9",
			TestFunc: testEvaluateParens,
		},
		{
			Slug:     "dc1",
			TestFunc: testEvaluateUnary,
		},
		{
			Slug:     "bp3",
			TestFunc: testEvaluateFactor,
		},
		{
			Slug:     "jy2",
			TestFunc: testEvaluateTerm,
		},
		{
			Slug:     "jx8",
			TestFunc: testEvaluateConcat,
		},
		{
			Slug:     "et4",
			TestFunc: testEvaluateRelational,
		},
		{
			Slug:     "hw7",
			TestFunc: testEvaluateEquality,
		},
		{
			Slug:     "gj9",
			TestFunc: testEvaluateUnaryErrors,
		},
		{
			Slug:     "yu6",
			TestFunc: testEvaluateMultErrors,
		},
		{
			Slug:     "cq1",
			TestFunc: testEvaluateAddErrors,
		},
		{
			Slug:     "ib5",
			TestFunc: testEvaluateCompErrors,
		},
		{
			Slug:     "xy1",
			TestFunc: createTestStatementFunction("s1"),
		},
		{
			Slug:     "oe4",
			TestFunc: createTestStatementFunction("s2"),
		},
		{
			Slug:     "fi3",
			TestFunc: createTestStatementFunction("s3"),
		},
		{
			Slug:     "yg2",
			TestFunc: createTestStatementFunction("s4"),
		},
		{
			Slug:     "sv7",
			TestFunc: createTestStatementFunction("s5"),
		},
		{
			Slug:     "bc1",
			TestFunc: createTestStatementFunction("s6"),
		},
		{
			Slug:     "dw9",
			TestFunc: createTestStatementFunction("s7"),
		},
		{
			Slug:     "pl3",
			TestFunc: createTestStatementFunction("s8"),
		},
		{
			Slug:     "vr5",
			TestFunc: createTestStatementFunction("s9"),
		},
		{
			Slug:     "fb4",
			TestFunc: createTestStatementFunction("s10"),
		},
	},
}
