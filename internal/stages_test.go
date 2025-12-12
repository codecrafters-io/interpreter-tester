package internal

import (
	"os"
	"regexp"
	"testing"

	tester_utils_testing "github.com/codecrafters-io/tester-utils/testing"
)

func TestStages(t *testing.T) {
	os.Setenv("CODECRAFTERS_RANDOM_SEED", "1234567890")

	testCases := map[string]tester_utils_testing.TesterOutputTestCase{
		"pass_scanning_jlox": {
			UntilStageSlug:      "pq5",
			CodePath:            "../../craftinginterpreters/build/gen/chap04_scanning",
			ExpectedExitCode:    0,
			StdoutFixturePath:   "./test_helpers/fixtures/pass_scanning",
			NormalizeOutputFunc: normalizeTesterOutput,
		},
		"pass_parsing_jlox": {
			StageSlugs:          []string{"wz8", "ht8", "uh4", "yf2", "wa9", "mq1", "xe6", "th5", "ra8", "sc2"},
			CodePath:            "../../craftinginterpreters/build/gen/chap06_parsing",
			ExpectedExitCode:    0,
			StdoutFixturePath:   "./test_helpers/fixtures/pass_parsing",
			NormalizeOutputFunc: normalizeTesterOutput,
		},
		"pass_evaluating_jlox": {
			StageSlugs:          []string{"ib5", "cq1", "yu6", "gj9", "hw7", "et4", "jx8", "jy2", "bp3", "dc1", "oq9", "lv1", "iz6"},
			CodePath:            "../../craftinginterpreters/build/gen/chap07_evaluating",
			ExpectedExitCode:    0,
			StdoutFixturePath:   "./test_helpers/fixtures/pass_evaluating",
			NormalizeOutputFunc: normalizeTesterOutput,
		},
		"pass_statements_inprogress_jlox": {
			StageSlugs:          []string{"xy1", "oe4", "fi3", "yg2", "sv7", "bc1", "dw9", "pl3", "vr5", "fb4"},
			CodePath:            "../../craftinginterpreters/build/gen/chap08_statements",
			ExpectedExitCode:    0,
			StdoutFixturePath:   "./test_helpers/fixtures/pass_statements",
			NormalizeOutputFunc: normalizeTesterOutput,
		},
		"pass_statements_completed_jlox": {
			StageSlugs:          []string{"xy1", "oe4", "fi3", "yg2", "sv7", "bc1", "dw9", "pl3", "vr5", "fb4"},
			CodePath:            "../../craftinginterpreters/build/gen/chap13_inheritance",
			ExpectedExitCode:    0,
			StdoutFixturePath:   "./test_helpers/fixtures/pass_statements_final",
			NormalizeOutputFunc: normalizeTesterOutput,
		},
		"pass_control_flow_inprogress_jlox": {
			StageSlugs:          []string{"ne3", "st5", "fh8", "xj4", "wk8", "jx4", "qy3", "bw6", "vt1"},
			CodePath:            "../../craftinginterpreters/build/gen/chap09_control",
			ExpectedExitCode:    0,
			StdoutFixturePath:   "./test_helpers/fixtures/pass_control_flow",
			NormalizeOutputFunc: normalizeTesterOutput,
		},
		"pass_control_flow_completed_jlox": {
			StageSlugs:          []string{"ne3", "st5", "fh8", "xj4", "wk8", "jx4", "qy3", "bw6", "vt1"},
			CodePath:            "../../craftinginterpreters/build/gen/chap13_inheritance",
			ExpectedExitCode:    0,
			StdoutFixturePath:   "./test_helpers/fixtures/pass_control_flow_final",
			NormalizeOutputFunc: normalizeTesterOutput,
		},
		"pass_functions_inprogress_jlox": {
			StageSlugs:          []string{"av4", "pg8", "lb6", "px4", "rd2", "ey3", "fj7", "bz4", "gg6"},
			CodePath:            "../../craftinginterpreters/build/gen/chap10_functions",
			ExpectedExitCode:    0,
			StdoutFixturePath:   "./test_helpers/fixtures/pass_functions",
			NormalizeOutputFunc: normalizeTesterOutput,
		},
		"pass_functions_completed_jlox": {
			StageSlugs:          []string{"av4", "pg8", "lb6", "px4", "rd2", "ey3", "fj7", "bz4", "gg6"},
			CodePath:            "../../craftinginterpreters/build/gen/chap13_inheritance",
			ExpectedExitCode:    0,
			StdoutFixturePath:   "./test_helpers/fixtures/pass_functions_final",
			NormalizeOutputFunc: normalizeTesterOutput,
		},
		"pass_resolving_inprogress_jlox": {
			StageSlugs:          []string{"de8", "pt7", "pz7", "eh3"},
			CodePath:            "../../craftinginterpreters/build/gen/chap11_resolving",
			ExpectedExitCode:    0,
			StdoutFixturePath:   "./test_helpers/fixtures/pass_resolving",
			NormalizeOutputFunc: normalizeTesterOutput,
		},
		"pass_resolving_completed_jlox": {
			StageSlugs:          []string{"de8", "pt7", "pz7", "eh3"},
			CodePath:            "../../craftinginterpreters/build/gen/chap13_inheritance",
			ExpectedExitCode:    0,
			StdoutFixturePath:   "./test_helpers/fixtures/pass_resolving_final",
			NormalizeOutputFunc: normalizeTesterOutput,
		},
		"pass_classes_inprogress_jlox": {
			StageSlugs:          []string{"vf4", "yk8", "yf3", "qr2", "yd7", "dg2", "ou5", "eb9"},
			CodePath:            "../../craftinginterpreters/build/gen/chap12_classes",
			ExpectedExitCode:    0,
			StdoutFixturePath:   "./test_helpers/fixtures/pass_classes",
			NormalizeOutputFunc: normalizeTesterOutput,
		},
		"pass_classes_completed_jlox": {
			StageSlugs:          []string{"vf4", "yk8", "yf3", "qr2", "yd7", "dg2", "ou5", "eb9"},
			CodePath:            "../../craftinginterpreters/build/gen/chap13_inheritance",
			ExpectedExitCode:    0,
			StdoutFixturePath:   "./test_helpers/fixtures/pass_classes_final",
			NormalizeOutputFunc: normalizeTesterOutput,
		},
		"pass_inheritance_jlox": {
			StageSlugs:          []string{"mf6", "ky1", "ka5", "ab0", "qi0", "ib9"},
			CodePath:            "../../craftinginterpreters/build/gen/chap13_inheritance",
			ExpectedExitCode:    0,
			StdoutFixturePath:   "./test_helpers/fixtures/pass_inheritance",
			NormalizeOutputFunc: normalizeTesterOutput,
		},
	}

	tester_utils_testing.TestTesterOutput(t, testerDefinition, testCases)
}

func normalizeTesterOutput(testerOutput []byte) []byte {
	replacements := map[string][]*regexp.Regexp{
		"clock_exponent_notation": {regexp.MustCompile(`1.[0-9]*E[0-9]`)},
		"clock_decimal_notation":  {regexp.MustCompile(`[0-9]{7,10}\.[0-9]{4,10}`)}, // This is much more restrictive so as to not match any other numbers
	}

	for replacement, regexes := range replacements {
		for _, regex := range regexes {
			testerOutput = regex.ReplaceAll(testerOutput, []byte(replacement))
		}
	}

	return testerOutput
}
