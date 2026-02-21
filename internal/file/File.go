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

func (f *File) longestLeftSideLength() int {
	maxLength := 0

	for _, line := range f.content {
		leftHandMatch := leftHandRegex.FindString(line)

		if len(leftHandMatch) > maxLength {
			maxLength = len(leftHandMatch)
		}
	}

	return maxLength
}

func (f *File) align() {
	longestLeftSideLength := f.longestLeftSideLength()

	for i, line := range f.content {
		leftHandMatch := leftHandRegex.FindString(line)
		trueLength := longestLeftSideLength - len(leftHandMatch) + 1

		if trueLength < 0 {
			trueLength = 1
		}

		(f.content)[i] = spaceRegex.ReplaceAllString(line, strings.Repeat(" ", trueLength))
	}
}

func (f *File) Process() error {
	f.align()
	output := strings.Join(f.content, "\n") + "\n"
	return os.WriteFile(f.path, []byte(output), 0o644)
}
