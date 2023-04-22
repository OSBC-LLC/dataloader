package string

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"runtime"
)

func getRootPath() (string, error) {
	_, currentFilePath, _, ok := runtime.Caller(0)
	if !ok {
		return "", fmt.Errorf("Error getting the current file path")
	}
	rootPath := filepath.Join(filepath.Dir(currentFilePath), "..", "..")
	absRootPath, err := filepath.Abs(rootPath)
	if err != nil {
		return "", err
	}
	return absRootPath, nil
}

func readFileLines(filePath string) ([]string, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return lines, nil
}
