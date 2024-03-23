package filemanager

import (
	"bufio"
	"encoding/json"
	"errors"
	"os"
	"time"
)

type FileManager struct {
	InputFilePath  string
	OutputFilePath string
}

func NewFileManager(inputPath, outputPath string) *FileManager {
	return &FileManager{
		InputFilePath:  inputPath,
		OutputFilePath: outputPath,
	}
}

func (fm FileManager) ReadLines() ([]string, error) {
	// no lines read initially
	lines := []string{}
	// open file
	file, err := os.Open(fm.InputFilePath)

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

func (fm FileManager) WriteJSON(data any) error {
	// create the file
	file, err := os.Create(fm.OutputFilePath)
	if err != nil {
		return errors.New("failed to create file")
	}
	defer file.Close()

	time.Sleep(3 * time.Second)

	encoder := json.NewEncoder(file)
	if err := encoder.Encode(data); err != nil {
		return errors.New("failed to convert data to JSON")
	}
	return nil
}
