package internal

import (
	"os"
	"path/filepath"
	"regexp"
	"strings"
	"testing"
)

func TestEmptyTrailingLines(t *testing.T) {
	os.Chdir("../test_programs")
	// Find all .lox files in the repository
	files, err := filepath.Glob("*/*.lox")
	if err != nil {
		t.Fatalf("Error finding .lox files: %v", err)
	}

	for _, file := range files {
		// Skip empty files
		info, err := os.Stat(file)
		if err != nil {
			t.Errorf("Error getting file info for %s: %v", file, err)
			continue
		}
		if info.Size() == 0 {
			continue
		}

		// Read the last byte of the file
		f, err := os.Open(file)
		if err != nil {
			t.Errorf("Error opening file %s: %v", file, err)
			continue
		}
		defer f.Close()

		// Seek to the last byte
		_, err = f.Seek(-1, 2)
		if err != nil {
			t.Errorf("Error seeking in file %s: %v", file, err)
			continue
		}

		// Read the last byte
		buf := make([]byte, 1)
		_, err = f.Read(buf)
		if err != nil {
			t.Errorf("Error reading last byte of file %s: %v", file, err)
			continue
		}

		// Check if the last byte is a newline
		if buf[0] == '\n' {
			t.Errorf("File %s ends with a newline", file)
		}
	}
}

