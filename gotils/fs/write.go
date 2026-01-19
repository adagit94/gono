package fs

import (
	"os"
	"path/filepath"
)

func WriteFile(data []byte, perm os.FileMode, pathSegments ...string) error {
	path := filepath.Join(pathSegments...)
	return os.WriteFile(path, data, perm)
}
