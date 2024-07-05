package testcases

import (
	"fmt"
	"os"
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
