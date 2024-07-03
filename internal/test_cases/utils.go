package testcases

import (
	"fmt"
	"os"
	"regexp"
	"strings"

	"github.com/codecrafters-io/tester-utils/executable"
)

func createTempFileWithContents(contents string) (string, error) {
	tmpFileName := "test.lox"
	if _, err := os.Stat(tmpFileName); err == nil {
		err := os.Remove(tmpFileName)
		if err != nil {
			return "", fmt.Errorf("CodeCrafters internal error. Error removing tmp file: %v", err)
		}
	}

	tmpFile, err := os.Create(tmpFileName)
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

func getStdoutLinesFromExecutableResult(result executable.ExecutableResult) []string {
	stdout := strings.Split(strings.TrimRight(string(result.Stdout), "\n"), "\n")
	return stdout
}

func getStderrLinesFromExecutableResult(result executable.ExecutableResult) []string {
	var filteredStdErr []string
	stderr := strings.Split(strings.TrimRight(string(result.Stderr), "\n"), "\n")
	regex := regexp.MustCompile(`\[line [0-9]+\]`)

	for _, line := range stderr {
		if regex.MatchString(line) {
			filteredStdErr = append(filteredStdErr, line)
		}
	}

	return filteredStdErr
}
