package helper

import (
	"fmt"
	"io"
	"os"
	"strings"
)

func ReadFile(path string, number int) []string {
	file, err := os.Open(path)
	content, err := io.ReadAll(file)

	if err != nil {
		fmt.Println("Cannot open file")
	}
	allSections := strings.Split(string(content), "---")

	poem := strings.TrimSpace(allSections[number-1])
	lines := strings.Split(poem, "\n")
	return lines
}
