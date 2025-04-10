package internal

import (
	"os"
	"path/filepath"
	"regexp"
	"strings"
	"testing"
)

func findLoxFiles(t *testing.T) []string {
	// Get the absolute path of the test_programs directory
	absPath, err := filepath.Abs("../test_programs")
	if err != nil {
		t.Fatalf("Error getting absolute path: %v", err)
	}

	// Find all .lox files in the repository
	files, err := filepath.Glob(filepath.Join(absPath, "*/*.lox"))
	if err != nil {
		t.Fatalf("Error finding .lox files: %v", err)
	}

	return files
}

func getFileName(t *testing.T, file string) string {
	fileNameParts := strings.Split(file, string(os.PathSeparator))
	dir := fileNameParts[len(fileNameParts)-2]
	fileName := fileNameParts[len(fileNameParts)-1]
	return dir + string(os.PathSeparator) + fileName
}

func TestEmptyTrailingLines(t *testing.T) {
	files := findLoxFiles(t)

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
			t.Errorf("File %s ends with a newline", getFileName(t, file))
		}
	}
}

func TestLineLength(t *testing.T) {
	const maxLineLength = 51
	files := findLoxFiles(t)

	// Compile regex patterns for random placeholders
	randomPatterns := map[*regexp.Regexp]string{
		regexp.MustCompile(`<<RANDOM_STRING(_[0-9]+)?>>`):       "hello",
		regexp.MustCompile(`<<RANDOM_QUOTEDSTRING(_[0-9]+)?>>`): "\"hello\"",
		regexp.MustCompile(`<<RANDOM_INTEGER(_[0-9]+)?>>`):      "99",
		regexp.MustCompile(`<<RANDOM_BOOLEAN(_[0-9]+)?>>`):      "false",
		regexp.MustCompile(`<<RANDOM_DIGIT(_[0-9]+)?>>`):        "3",
	}

	for _, file := range files {
		content, err := os.ReadFile(file)
		if err != nil {
			t.Errorf("Error reading file %s: %v", file, err)
			continue
		}

		lines := strings.Split(string(content), "\n")
		for i, line := range lines {
			// Replace all random placeholders using regex
			for pattern, replacement := range randomPatterns {
				line = pattern.ReplaceAllString(line, replacement)
			}

			if len(line) > maxLineLength {
				t.Errorf("Line length exceeds %d in file: %s, line number: %d, line length: %d",
					maxLineLength, getFileName(t, file), i+1, len(line))
			}
		}
	}
}

func TestNoTrailingSpaces(t *testing.T) {
	files := findLoxFiles(t)

	for _, file := range files {
		content, err := os.ReadFile(file)
		if err != nil {
			t.Errorf("Error reading file %s: %v", file, err)
			continue
		}

		lines := strings.Split(string(content), "\n")
		for i, line := range lines {
			if strings.HasSuffix(line, " ") {
				t.Errorf("Trailing space found in file: %s, line number: %d", getFileName(t, file), i+1)
			}
			if strings.HasSuffix(line, "\t") {
				t.Errorf("Trailing tab found in file: %s, line number: %d", getFileName(t, file), i+1)
			}
		}
	}
}
