package utils

import (
	"io/ioutil"
	"strings"
)

// ReadLines reads a file, splits it into lines by the given splitter
// and returns an array of strings
func ReadLines(path, splitOn string) ([]string, error) {
	content, readFileErr := ioutil.ReadFile(path)
	if readFileErr != nil {
		return nil, readFileErr
	}
	lines := strings.Split(string(content), splitOn)

	return lines, nil
}
