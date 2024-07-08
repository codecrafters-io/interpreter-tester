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
			StdoutFixturePath:   "./test_helpers/fixtures/pass_all",
			NormalizeOutputFunc: normalizeTesterOutput,
		},
	}

	tester_utils_testing.TestTesterOutput(t, testerDefinition, testCases)
}

func normalizeTesterOutput(testerOutput []byte) []byte {
	return testerOutput
}
