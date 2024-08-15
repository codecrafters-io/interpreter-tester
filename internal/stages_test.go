package internal

import (
	"os"
	"testing"

	tester_utils_testing "github.com/codecrafters-io/tester-utils/testing"
)

func TestStages(t *testing.T) {
	os.Setenv("CODECRAFTERS_RANDOM_SEED", "1234567890")

	testCases := map[string]tester_utils_testing.TesterOutputTestCase{
		"pass_scanning_jlox": {
			UntilStageSlug:      "pq5",
			CodePath:            "../craftinginterpreters/build/gen/chap04_scanning",
			ExpectedExitCode:    0,
			StdoutFixturePath:   "./test_helpers/fixtures/pass_scanning",
			NormalizeOutputFunc: normalizeTesterOutput,
		},
		"pass_parsing_jlox": {
			StageSlugs:          []string{"wz8", "ht8", "uh4", "yf2", "wa9", "mq1", "xe6", "th5", "ra8", "sc2"},
			CodePath:            "../craftinginterpreters/build/gen/chap06_parsing",
			ExpectedExitCode:    0,
			StdoutFixturePath:   "./test_helpers/fixtures/pass_parsing",
			NormalizeOutputFunc: normalizeTesterOutput,
		},
		"pass_evaluating_jlox": {
			StageSlugs:          []string{"ib5", "cq1", "yu6", "gj9", "hw7", "et4", "jx8", "jy2", "bp3", "dc1", "oq9", "lv1", "iz6"},
			CodePath:            "../craftinginterpreters/build/gen/chap07_evaluating",
			ExpectedExitCode:    0,
			StdoutFixturePath:   "./test_helpers/fixtures/pass_evaluating",
			NormalizeOutputFunc: normalizeTesterOutput,
		},
		"pass_statements_jlox": {
			StageSlugs:          []string{"xy1", "oe4", "fi3", "yg2", "sv7", "bc1", "dw9", "pl3", "vr5", "fb4"},
			CodePath:            "../craftinginterpreters/build/gen/chap08_statements",
			ExpectedExitCode:    0,
			StdoutFixturePath:   "./test_helpers/fixtures/pass_statements",
			NormalizeOutputFunc: normalizeTesterOutput,
		},
	}

	tester_utils_testing.TestTesterOutput(t, testerDefinition, testCases)
}

func normalizeTesterOutput(testerOutput []byte) []byte {
	return testerOutput
}
