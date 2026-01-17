package fs

import (
	"os"
	"path/filepath"
)

func ReadFile(pathSegments ...string) ([]byte, error) {
	path := filepath.Join(pathSegments...)
	return os.ReadFile(path)
}