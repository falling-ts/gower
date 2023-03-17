package utils

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func SetEnv(key, value string) error {
	file, err := os.Open(".env")
	if err != nil {
		return err
	}
	defer func(file *os.File) {
		_ = file.Close()
	}(file)

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if strings.HasPrefix(line, key+"=") {
			line = fmt.Sprintf("%s=%s", key, value)
		}
		lines = append(lines, line)
	}

	if err = scanner.Err(); err != nil {
		return err
	}

	if err = file.Close(); err != nil {
		return err
	}

	outputFile, err := os.Create(".env")
	if err != nil {
		return err
	}
	defer func(outputFile *os.File) {
		_ = outputFile.Close()
	}(outputFile)

	writer := bufio.NewWriter(outputFile)
	for _, line := range lines {
		_, err = writer.WriteString(line + "\n")
		if err != nil {
			return err
		}
	}

	return writer.Flush()
}
