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
			TestFunc: createTestForRunCommandStage("s1"),
		},
		{
			Slug:     "oe4",
			TestFunc: createTestForRunCommandStage("s2"),
		},
		{
			Slug:     "fi3",
			TestFunc: createTestForRunCommandStage("s3"),
		},
		{
			Slug:     "yg2",
			TestFunc: createTestForRunCommandStage("s4"),
		},
		{
			Slug:     "sv7",
			TestFunc: createTestForRunCommandStage("s5"),
		},
		{
			Slug:     "bc1",
			TestFunc: createTestForRunCommandStage("s6"),
		},
		{
			Slug:     "dw9",
			TestFunc: createTestForRunCommandStage("s7"),
		},
		{
			Slug:     "pl3",
			TestFunc: createTestForRunCommandStage("s8"),
		},
		{
			Slug:     "vr5",
			TestFunc: createTestForRunCommandStage("s9"),
		},
		{
			Slug:     "fb4",
			TestFunc: createTestForRunCommandStage("s10"),
		},
		{
			Slug:     "ne3",
			TestFunc: createTestForRunCommandStage("c1"),
		},
		{
			Slug:     "st5",
			TestFunc: createTestForRunCommandStage("c2"),
		},
		{
			Slug:     "fh8",
			TestFunc: createTestForRunCommandStage("c3"),
		},
		{
			Slug:     "xj4",
			TestFunc: createTestForRunCommandStage("c4"),
		},
		{
			Slug:     "wk8",
			TestFunc: createTestForRunCommandStage("c5"),
		},
		{
			Slug:     "jx4",
			TestFunc: createTestForRunCommandStage("c6"),
		},
		{
			Slug:     "qy3",
			TestFunc: createTestForRunCommandStage("c7"),
		},
		{
			Slug:     "bw6",
			TestFunc: createTestForRunCommandStage("c8"),
		},
		{
			Slug:     "vt1",
			TestFunc: createTestForRunCommandStage("c9"),
		},
		{
			Slug:     "av4",
			TestFunc: testClock, // This needs a special test function because it has an integer assertion, for clock()
		},
		{
			Slug:     "pg8",
			TestFunc: createTestForRunCommandStage("f2"),
		},
		{
			Slug:     "lb6",
			TestFunc: createTestForRunCommandStage("f3"),
		},
		{
			Slug:     "px4",
			TestFunc: createTestForRunCommandStage("f4"),
		},
		{
			Slug:     "rd2",
			TestFunc: createTestForRunCommandStage("f5"),
		},
		{
			Slug:     "ey3",
			TestFunc: createTestForRunCommandStage("f6"),
		},
		{
			Slug:     "fj7",
			TestFunc: createTestForRunCommandStage("f7"),
		},
		{
			Slug:     "bz4",
			TestFunc: createTestForRunCommandStage("f8"),
		},
		{
			Slug:     "gg6",
			TestFunc: createTestForRunCommandStage("f9"),
		},
		{
			Slug:     "de8",
			TestFunc: createTestForRunCommandStage("r1"),
		},
		{
			Slug:     "pt7",
			TestFunc: createTestForRunCommandStage("r2"),
		},
		{
			Slug:     "pz7",
			TestFunc: createTestForRunCommandStage("r3"),
		},
		{
			Slug:     "eh3",
			TestFunc: createTestForRunCommandStage("r4"),
		},
		{
			Slug:     "vf4",
			TestFunc: createTestForRunCommandStage("cl1"),
		},
		{
			Slug:     "yk8",
			TestFunc: createTestForRunCommandStage("cl2"),
		},
		{
			Slug:     "yf3",
			TestFunc: createTestForRunCommandStage("cl3"),
		},
		{
			Slug:     "qr2",
			TestFunc: createTestForRunCommandStage("cl4"),
		},
		{
			Slug:     "yd7",
			TestFunc: createTestForRunCommandStage("cl5"),
		},
		{
			Slug:     "dg2",
			TestFunc: createTestForRunCommandStage("cl6"),
		},
		{
			Slug:     "ou5",
			TestFunc: createTestForRunCommandStage("cl7"),
		},
		{
			Slug:     "eb9",
			TestFunc: createTestForRunCommandStage("cl8"),
		},
		{
			Slug:     "mf6",
			TestFunc: createTestForRunCommandStage("i1"),
		},
		{
			Slug:     "ky1",
			TestFunc: createTestForRunCommandStage("i2"),
		},
		{
			Slug:     "ka5",
			TestFunc: createTestForRunCommandStage("i3"),
		},
		{
			Slug:     "ab0",
			TestFunc: createTestForRunCommandStage("i4"),
		},
		{
			Slug:     "qi0",
			TestFunc: createTestForRunCommandStage("i5"),
		},
		{
			Slug:     "ib9",
			TestFunc: createTestForRunCommandStage("i6"),
		},
	},
}
