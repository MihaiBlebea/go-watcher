package snapshot

import (
	"strings"
)

func isExecutable(filePath string) bool {
	fileName := filePath

	if strings.Contains(filePath, "/") == true {
		parts := strings.Split(filePath, "/")
		fileName = parts[len(parts)-1]
	}

	if strings.Contains(fileName, ".") {
		return false
	}

	return true
}
