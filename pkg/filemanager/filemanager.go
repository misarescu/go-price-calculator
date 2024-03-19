package filemanager

import (
	"bufio"
	"os"
)

func ReadLines(path string) ([]string, error) {
	// no lines read initially
	lines := []string{}
	// open file
	file, err := os.Open("data/prices.txt")

	if err != nil {
		return nil, err
	}

	defer file.Close()

	// create scanner
	s := bufio.NewScanner(file)

	for s.Scan() {
		lines = append(lines, s.Text())
	}

	if err = s.Err(); err != nil {
		return nil, err
	}

	return lines, nil
}
