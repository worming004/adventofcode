package parser

import (
	"bufio"
	"os"
)

func Parse(filePath string) []string {
	f, err := os.Open(filePath)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	var result []string

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		t := scanner.Text()
		result = append(result, t)
	}

	return result
}
