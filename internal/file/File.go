package file

import (
	"os"
	"regexp"
	"strings"
)

var (
	spaceRegex    = regexp.MustCompile(" +")
	leftHandRegex = regexp.MustCompile(".+? ")
)

type File struct {
	content []string
	path    string
}

func New(path string) (*File, error) {
	fileContent, err := os.ReadFile(path)

	if err != nil {
		return nil, err
	}

	content := strings.Split(string(fileContent), "\n")

	// TODO: Use last part from path
	return &File{
		content: content,
		path:    path,
	}, nil
}

func (f *File) longestLineLength() int {
	maxLength := 0

	for _, line := range f.content {
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

func (f *File) align() {
	maxLength := f.longestLineLength()

	for i, line := range f.content {
		leftHandMatch := leftHandRegex.FindString(line)
		trueLength := maxLength - len(leftHandMatch) + 1

		if trueLength < 0 {
			trueLength = 1
		}

		(f.content)[i] = spaceRegex.ReplaceAllString(line, strings.Repeat(" ", trueLength))
	}
}

func (f *File) store() error {
	output := strings.Join(f.content, "\n") + "\n"

	err := os.WriteFile(f.path, []byte(output), 0o644)

	if err != nil {
		return err
	}

	return nil
}

func (f *File) Process() error {
	f.align()
	return f.store()
}
