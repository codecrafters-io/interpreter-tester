package internal

import (
	"os"
	"path/filepath"
	"regexp"
	"strings"
	"testing"
)

// This file contains tests to validate the formatting of .lox files.
// It ensures that all .lox files in the repository adhere to specific
// formatting rules like line length limits, no trailing spaces, and
// no trailing newlines.

// MaxLineLength defines the maximum allowed line length for .lox files
const maxLineLength = 51

// Compile regex patterns for random placeholders
var randomPatterns = map[*regexp.Regexp]string{
	regexp.MustCompile(`<<RANDOM_STRING(_[0-9]+)?>>`):       "hello",
	regexp.MustCompile(`<<RANDOM_QUOTEDSTRING(_[0-9]+)?>>`): "\"hello\"",
	regexp.MustCompile(`<<RANDOM_INTEGER(_[0-9]+)?>>`):      "99",
	regexp.MustCompile(`<<RANDOM_BOOLEAN(_[0-9]+)?>>`):      "false",
	regexp.MustCompile(`<<RANDOM_DIGIT(_[0-9]+)?>>`):        "3",
}

func findLoxFiles(t *testing.T) []string {
	// Get the absolute path of the test_programs directory
	absPath, err := filepath.Abs("../test_programs")
	if err != nil {
		t.Fatalf("Error getting absolute path: %v", err)
	}

	// Find all .lox files in the repository
	var files []string
	err = filepath.Walk(absPath, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			// Handle errors accessing the file/dir
			return err
		}
		if !info.IsDir() && filepath.Ext(path) == ".lox" { // Check if it's a file and has the .lox extension
			files = append(files, path)
		}
		// Continue walking
		return nil
	})
	if err != nil {
		// Handle errors during the walk itself
		t.Fatalf("Error finding .lox files: %v", err)
	}
	return files
}

func getFileName(file string) string {
	// Use filepath functions to safely extract components
	dir := filepath.Base(filepath.Dir(file))
	fileName := filepath.Base(file)
	return filepath.Join(dir, fileName)
}

func TestEmptyTrailingLines(t *testing.T) {
	files := findLoxFiles(t)

	for _, file := range files {
		// Read the entire file
		content, err := os.ReadFile(file)
		if err != nil {
			t.Errorf("Error reading file %s: %v", file, err)
			continue
		}

		// Skip empty files
		if len(content) == 0 {
			continue
		}

		// Check if the last byte is a newline
		if content[len(content)-1] == '\n' {
			t.Errorf("File %s ends with a newline", getFileName(file))
		}
	}
}

func TestLineLength(t *testing.T) {
	files := findLoxFiles(t)

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

			// Skip empty files
			if len(content) == 0 {
				continue
			}

			if len(line) > maxLineLength {
				t.Errorf("Line length exceeds %d in file: %s, line number: %d, line length: %d",
					maxLineLength, getFileName(file), i+1, len(line))
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

		// Skip empty files
		if len(content) == 0 {
			continue
		}

		lines := strings.Split(string(content), "\n")
		for i, line := range lines {
			if strings.HasSuffix(line, " ") {
				t.Errorf("Trailing space found in file: %s, line number: %d", getFileName(file), i+1)
			}
			if strings.HasSuffix(line, "\t") {
				t.Errorf("Trailing tab found in file: %s, line number: %d", getFileName(file), i+1)
			}
		}
	}
}
