package file

import (
	"os"
	"regexp"
	"strings"
)

type File struct {
	name    string
	content *[]string
	path    string
}

func FromPath(path string) (*File, error) {
	fileContent, err := os.ReadFile(path)

	if err != nil {
		return nil, err
	}

	line := ""
	var content []string

	for _, byteValue := range fileContent {
		if byteValue == 10 {
			content = append(content, line)
			line = ""
			continue
		}

		line = line + string(byteValue)
	}

	// TODO: Use last part from path
	return &File{"prueba", &content, path}, nil
}

func (f *File) longestLineLength() int {
	maxLength := 0

	for _, line := range *f.content {
		lineLength := 0
		isCountingSpaces := false

		for _, char := range line {
			if isCountingSpaces && char != ' ' {
				break
			}

			lineLength++

			if char == ' ' {
				isCountingSpaces = true
			}
		}

		if lineLength > maxLength {
			maxLength = lineLength
		}
	}

	return maxLength
}

func (f *File) align() (*File, error) {
	maxLength := f.longestLineLength()

	regex, err := regexp.Compile(" +")

	if err != nil {
		return nil, err
	}

	leftHandRegex, err := regexp.Compile(".+? ")

	if err != nil {
		return nil, err
	}

	for i, line := range *f.content {
		leftHandMatch := leftHandRegex.FindStringSubmatch(line)
		trueLength := maxLength - len(leftHandMatch[0]) + 1

		if trueLength < 0 {
			trueLength = 1
		}

		(*f.content)[i] = regex.ReplaceAllString(line, strings.Repeat(" ", trueLength))
	}

	return f, nil
}

func (f *File) store() error {
	var byteContent []byte

	for _, line := range *f.content {
		lineBytes := []byte(line + "\n")
		byteContent = append(byteContent, lineBytes...)
	}

	err := os.WriteFile(f.path, byteContent, 0o644)

	if err != nil {
		return err
	}

	return nil
}

func (f *File) Process() error {
	f, err := f.align()

	if err != nil {
		return err
	}

	err = f.store()

	if err != nil {
		return err
	}

	return nil
}
