package testcases

import (
	"fmt"
	"github.com/codecrafters-io/tester-utils/logger"
	"os"
	"regexp"
	"strings"
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

func (t *ParseTestCase) logReadableFileContents(logger *logger.Logger) {
	logger.Infof("Writing contents to ./test.lox:")

	// If the file contents contain a singe %, it will be decoded as a format specifier
	// And it will add a `(MISSING)` to the log line
	printableFileContents := strings.ReplaceAll(t.FileContents, "%", "%%")
	printableFileContents = strings.ReplaceAll(printableFileContents, "\t", "<|TAB|>")

	regex1 := regexp.MustCompile("[ ]+\n")
	regex2 := regexp.MustCompile("[ ]+$")
	printableFileContents = regex1.ReplaceAllString(printableFileContents, "<|SPACE|>")
	printableFileContents = regex2.ReplaceAllString(printableFileContents, "<|SPACE|>")

	for _, line := range strings.Split(printableFileContents, "\n") {
		logger.Infof("[test.lox] " + line)
	}
}
