package internal

import (
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"strings"

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
func GetTestProgramsForCurrentStage(stageIdentifier string) []string {
	var testPrograms []string

	parentDir := "./internal/test_programs"
	testDir := filepath.Join(parentDir, stageIdentifier)
	files, err := os.ReadDir(testDir)
	if err != nil {
		panic(fmt.Sprintf("CodeCrafters Internal Error: Encountered error while reading test directory: %s", err))
	}

	for _, file := range files {
		filePath := filepath.Join(testDir, file.Name())
		contents, err := os.ReadFile(filePath)
		if err != nil {
			panic(fmt.Sprintf("CodeCrafters Internal Error: Encountered error while reading test file: %s", err))
		}
		testPrograms = append(testPrograms, string(contents))
	}
	return testPrograms
}

func ReplacePlaceholdersWithRandomValues(program string) string {
	regexPlaceholder := regexp.MustCompile(`<<(RANDOM_STRING|RANDOM_QUOTEDSTRING|RANDOM_BOOLEAN|RANDOM_INTEGER)(_\d+)?>>`)

	generatedValues := make(map[string]string)

	// Replace placeholders with random values
	result := regexPlaceholder.ReplaceAllStringFunc(program, func(match string) string {
		// match looks like: <<RANDOM_DTYPE_N>>
		parts := strings.Split(match[2:len(match)-2], "_")
		placeholderType := strings.Join(parts[0:2], "_") // first 2 parts are guaranteed to be RANDOM & DTYPE
		placeholderID := ""

		if len(parts) > 2 {
			placeholderID = parts[2]
		}

		key := placeholderType + placeholderID
		if value, exists := generatedValues[key]; exists {
			return value
		}

		var newValue string
		switch placeholderType {
		case "RANDOM_STRING":
			newValue = random.RandomElementFromArray(STRINGS)
		case "RANDOM_QUOTEDSTRING":
			newValue = random.RandomElementFromArray(QUOTED_STRINGS)
		case "RANDOM_BOOLEAN":
			newValue = random.RandomElementFromArray(BOOLEANS)
		case "RANDOM_INTEGER":
			newValue = getRandIntAsString()
		default:
			return match
		}

		generatedValues[key] = newValue
		return newValue
	})

	return result
}

func GetTestProgramsForCurrentStageWithRandomValues(stageIdentifier string) []string {
	var testPrograms []string

	for _, program := range GetTestProgramsForCurrentStage(stageIdentifier) {
		testPrograms = append(testPrograms, ReplacePlaceholdersWithRandomValues(program))
	}

	return testPrograms
}
