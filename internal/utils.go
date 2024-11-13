package internal

import (
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"strings"

	testcases "github.com/codecrafters-io/interpreter-tester/internal/test_cases"
	"github.com/codecrafters-io/tester-utils/random"
)

func getRandInt() int {
	return random.RandomInt(10, 100)
}

func getRandIntAsString() string {
	return fmt.Sprintf("%d", getRandInt())
}

func getRandString() string {
	return random.RandomElementFromArray(STRINGS)
}

func getRandBoolean() string {
	return random.RandomElementFromArray(BOOLEANS)
}

// From the "statements & state" extension onwards, we'll need to pass multi line programs for testing.
// To make it easier to work with and track changes, all the test programs are stored in the test_programs directory.
// Every stage has its own directory. (For example, stage_s1.go has a test_programs/s1 directory.)
// This function reads all the files in the test_programs/<<stage_id>> directory and returns their contents as a slice of strings.
func GetTestCasesForCurrentStage(stageIdentifier string) []testcases.RunTestCase {
	var testCases []testcases.RunTestCase

	// Construct the path to the test_programs directory
	parentDir := filepath.Join(os.Getenv("TESTER_DIR"), "test_programs")
	testDir := filepath.Join(parentDir, stageIdentifier)
	files, err := os.ReadDir(testDir)
	if err != nil {
		panic(fmt.Sprintf("CodeCrafters Internal Error: Encountered error while reading test directory: %s", err))
	}

	for _, file := range files {
		filePath := filepath.Join(testDir, file.Name())
		testCases = append(testCases, testcases.NewRunTestCaseFromFilePath(filePath))
	}
	return testCases
}

func ReplacePlaceholdersWithRandomValues(program string) string {
	regexPlaceholder := regexp.MustCompile(`<<(RANDOM_STRING|RANDOM_QUOTEDSTRING|RANDOM_BOOLEAN|RANDOM_INTEGER|RANDOM_DIGIT)(_\d+)?>>`)

	generatedValues := make(map[string]string)
	seenValues := make(map[string]bool)

	placeholderTypes := []string{"RANDOM_STRING", "RANDOM_QUOTEDSTRING", "RANDOM_BOOLEAN", "RANDOM_INTEGER", "RANDOM_DIGIT"}
	placeholderIDs := []string{"0", "1", "2", "3", "4"}
	for _, placeholderID := range placeholderIDs {
		for _, placeholderType := range placeholderTypes {
			var value string

			for {
				value = generateRandomValueForPlaceholderType(placeholderType)
				// until we get a value that wasn't seen before we keep on generating new values
				// There are only 2 unique booleans, so we don't need to worry about that
				if !seenValues[value] || placeholderType == "RANDOM_BOOLEAN" {
					break
				}
			}
			seenValues[value] = true
			generatedValues[placeholderType+placeholderID] = value
		}
	}

	// Replace placeholders with random values
	result := regexPlaceholder.ReplaceAllStringFunc(program, func(match string) string {
		// match looks like: <<RANDOM_DTYPE_N>>
		parts := strings.Split(match[2:len(match)-2], "_")
		placeholderType := strings.Join(parts[0:2], "_") // first 2 parts are guaranteed to be RANDOM & DTYPE
		placeholderID := "0"

		if len(parts) > 2 {
			placeholderID = parts[2]
		}

		key := placeholderType + placeholderID
		if value, exists := generatedValues[key]; exists {
			return value
		} else {
			panic(fmt.Sprintf("CodeCrafters Internal Error: Placeholder %s not found in generated values", key))
		}
	})

	return result
}

func generateRandomValueForPlaceholderType(placeholderType string) string {
	var value string

	switch placeholderType {
	case "RANDOM_STRING":
		value = random.RandomElementFromArray(STRINGS)
	case "RANDOM_QUOTEDSTRING":
		value = random.RandomElementFromArray(QUOTED_STRINGS)
	case "RANDOM_BOOLEAN":
		value = random.RandomElementFromArray(BOOLEANS)
	case "RANDOM_INTEGER": // Between 10 - 100
		value = getRandIntAsString()
	case "RANDOM_DIGIT": // Between 2 - 6
		value = fmt.Sprintf("%d", random.RandomInt(2, 7))
	default:
		value = placeholderType
	}
	return value
}

func GetTestCasesForCurrentStageWithRandomValues(stageIdentifier string) []testcases.TestCase {
	var testCases []testcases.TestCase

	for _, t := range GetTestCasesForCurrentStage(stageIdentifier) {
		t.FileContents = ReplacePlaceholdersWithRandomValues(t.FileContents)
		testCases = append(testCases, &t)
	}

	return testCases
}
