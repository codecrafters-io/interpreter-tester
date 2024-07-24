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
			StageSlugs:          []string{"sc2", "ra8", "th5", "xe6", "mq1", "wa9", "yf2", "uh4", "ht8", "wz8"},
			CodePath:            "../craftinginterpreters/build/gen/chap06_parsing",
			ExpectedExitCode:    0,
			StdoutFixturePath:   "./test_helpers/fixtures/pass_parsing",
			NormalizeOutputFunc: normalizeTesterOutput,
		},
		"pass_evaluating_jlox": {
			StageSlugs:          []string{"iz6", "lv1", "oq9", "dc1"},
			CodePath:            "../craftinginterpreters/build/gen/chap07_evaluating",
			ExpectedExitCode:    0,
			StdoutFixturePath:   "./test_helpers/fixtures/pass_evaluating",
			NormalizeOutputFunc: normalizeTesterOutput,
		},
	}

	tester_utils_testing.TestTesterOutput(t, testerDefinition, testCases)
}

func normalizeTesterOutput(testerOutput []byte) []byte {
	return testerOutput
}
