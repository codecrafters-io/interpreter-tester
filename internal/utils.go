package internal

import (
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"runtime"
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
func GetTestProgramsForCurrentStage() []string {
	var testPrograms []string

	// Get caller info for the caller of this function (testStatements1())
	// GetTestProgramsForCurrentStage() <- GetTestProgramsForCurrentStageWithRandomValues() <- testStatements1()
	// We skip over 2 frames, to get to testStatements1()
	_, file, no, ok := runtime.Caller(2)
	if !ok {
		panic(fmt.Sprintf("CodeCrafters Internal Error: Encountered error while getting caller info: %s#%d", file, no))
	}

	parentDir := filepath.Dir(file)
	stageIdentifier := strings.Split(strings.Split(file, ".")[0], "_")[1]
	testDir := filepath.Join(parentDir, "test_programs", stageIdentifier)
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

func GetTestProgramsForCurrentStageWithRandomValues() []string {
	var testPrograms []string

	for _, program := range GetTestProgramsForCurrentStage() {
		testPrograms = append(testPrograms, ReplacePlaceholdersWithRandomValues(program))
	}

	return testPrograms
}
