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

	// Get caller info for the caller of this function
	_, file, no, ok := runtime.Caller(1)
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
	regexPlaceholder := regexp.MustCompile(`<<(RANDOM_STRING|RANDOM_QUOTED_STRING|RANDOM_BOOLEAN|RANDOM_INTEGER)>>`)

	// Replace placeholders with random values
	result := regexPlaceholder.ReplaceAllStringFunc(program, func(match string) string {
		switch match {
		case "<<RANDOM_STRING>>":
			return random.RandomElementFromArray(STRINGS)
		case "<<RANDOM_QUOTED_STRING>>":
			return random.RandomElementFromArray(QUOTED_STRINGS)
		case "<<RANDOM_BOOLEAN>>":
			return random.RandomElementFromArray(BOOLEANS)
		case "<<RANDOM_INTEGER>>":
			return getRandIntAsString()
		default:
			return match
		}
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
