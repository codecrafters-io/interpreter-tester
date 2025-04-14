package testcases

import (
	"fmt"
	"os"
	"regexp"
	"strings"

	"github.com/codecrafters-io/tester-utils/logger"
	"github.com/fatih/color"
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

func logReadableFileContents(logger *logger.Logger, fileContents string) {
	logger.Infof("Writing contents to ./test.lox:")

	// If the file contents contain a single %, it will be decoded as a format specifier
	// And it will add a `(MISSING)` to the log line
	printableFileContents := strings.ReplaceAll(fileContents, "%", "%%")
	printableFileContents = strings.ReplaceAll(printableFileContents, "\t", "<|TAB|>")

	regex1 := regexp.MustCompile("[ ]+\n")
	regex2 := regexp.MustCompile("[ ]+$")
	printableFileContents = regex1.ReplaceAllString(printableFileContents, "\n")
	printableFileContents = regex2.ReplaceAllString(printableFileContents, "<|SPACE|>")

	// This is of the form "test-N"
	oldPrefix := logger.GetSecondaryPrefix()
	testNumber := strings.TrimPrefix(oldPrefix, "test-")
	logger.UpdateSecondaryPrefix(fmt.Sprintf("test-%s.lox", testNumber))
	defer logger.UpdateSecondaryPrefix(oldPrefix)

	if len(printableFileContents) == 0 {
		logger.Plainf("<|EMPTY FILE|>")
	} else {
		for _, line := range strings.Split(printableFileContents, "\n") {
			if strings.Contains(line, "//") {
				code := (strings.Split(line, "//")[0])
				comment := "//" + (strings.Split(line, "//")[1])
				formattedLine := code + color.YellowString(comment)
				logger.Plainf(formattedLine)
			} else {
				logger.Plainf(line)
			}
		}
	}
}

func logReadableFileContentsPreservingWhitespace(logger *logger.Logger, fileContents string) {
	logger.Infof("Writing contents to ./test.lox:")

	// If the file contents contain a single %, it will be decoded as a format specifier
	// And it will add a `(MISSING)` to the log line
	printableFileContents := strings.ReplaceAll(fileContents, "%", "%%")
	printableFileContents = strings.ReplaceAll(printableFileContents, "\t", "<|TAB|>")
	printableFileContents = strings.ReplaceAll(printableFileContents, " ", "<|SPACE|>")

	// This is of the form "test-N"
	oldPrefix := logger.GetSecondaryPrefix()
	testNumber := strings.TrimPrefix(oldPrefix, "test-")
	logger.UpdateSecondaryPrefix(fmt.Sprintf("test-%s.lox", testNumber))
	defer logger.UpdateSecondaryPrefix(oldPrefix)

	if len(printableFileContents) == 0 {
		logger.Plainf("<|EMPTY FILE|>")
	} else {
		for _, line := range strings.Split(printableFileContents, "\n") {
			logger.Plainf(line)
		}
	}
}

var exitCodeToErrorTypeMapping = map[int]string{
	0:  "no error",
	65: "compile error",
	70: "runtime error",
}
