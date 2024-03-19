package filemanager

import (
	"bufio"
	"encoding/json"
	"errors"
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

func WriteJSON(path string, data any) error {
	// create the file
	file, err := os.Create(path)
	if err != nil {
		return errors.New("failed to create file")
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	if err := encoder.Encode(data); err != nil {
		return errors.New("failed to convert data to JSON")
	}
	return nil
}
