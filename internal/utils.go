package internal

import (
	"fmt"
	"os"
	"regexp"
	"strings"

	"github.com/codecrafters-io/tester-utils/executable"
)

// CreateTemp creates a new temporary file in the directory dir with a name beginning with prefix
// Cleaning the file is the responsibility of the caller
// As dir is empty, it will create the file in the default temp directory
func createTempFileWithContents(contents string) (string, error) {
	tmpFile, err := os.CreateTemp("", "test.lox")
	if err != nil {
		return "", fmt.Errorf("CodeCrafters internal error. Error creating tmp file: %v", err)
	}
	_, err = tmpFile.WriteString(contents)
	if err != nil {
		return "", fmt.Errorf("CodeCrafters internal error. Error writing to tmp file: %v", err)
	}
	err = tmpFile.Close()
	if err != nil {
		return "", fmt.Errorf("CodeCrafters internal error. Error closing tmp file: %v", err)
	}

	return tmpFile.Name(), nil
}

func getOutputFromStdOut(result executable.ExecutableResult) []string {
	stdOut := strings.Split(strings.TrimRight(string(result.Stdout), "\n"), "\n")
	return stdOut
}

func getOutputFromStdErr(result executable.ExecutableResult) []string {
	filteredStdErr := []string{}
	stdErr := strings.Split(strings.TrimRight(string(result.Stderr), "\n"), "\n")
	regexp := regexp.MustCompile(`\[line [0-9]+\]`)

	for _, line := range stdErr {
		if regexp.MatchString(line) {
			filteredStdErr = append(filteredStdErr, line)
		}
	}

	return filteredStdErr
}
