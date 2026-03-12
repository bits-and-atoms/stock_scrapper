package filemanager

import (
	"bufio"
	"encoding/json"
	"os"
)

type FileManager struct {
	InputFilePath  string
	OutputFilePath string
}

func (fm *FileManager) Set(inp, out string) {
	fm.InputFilePath = inp
	fm.OutputFilePath = out
}
func (fm FileManager) ReadLines(items *[]string) error {
	file, err := os.Open(fm.InputFilePath)
	if err != nil {
		return err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	buf := make([]byte, 0, 1024*1024)
	scanner.Buffer(buf, 1024*1024)
	for scanner.Scan() {
		*items = append(*items, scanner.Text())
	}

	return scanner.Err()
}
func (fm FileManager) WriteJson(data any) error {
	file, err := os.Create(fm.OutputFilePath)
	if err != nil {
		return err
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "  ")
	return encoder.Encode(data)
}
