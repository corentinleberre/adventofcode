package util

import (
	"bufio"
	"log"
	"os"
)

func ParseFile(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var lines []string
	for scanner.Scan() {
		var value = scanner.Text()
		lines = append(lines, value)
	}
	return lines, scanner.Err()
}
