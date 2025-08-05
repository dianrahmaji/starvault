package main

import (
	"os"
	"regexp"
	"strings"
)

func sanitizeStr(str string) string {
	s := strings.ReplaceAll(str, " ", "_")

	re := regexp.MustCompile(`[^\w\d_\-+]+`)
	s = re.ReplaceAllString(s, "_")

	return s
}

func createOutputDir() (string, error) {
	outputDir := "./recordings/temp"

	if err := os.MkdirAll(outputDir, 0755); err != nil {
		return "", err
	}

	return outputDir, nil
}
