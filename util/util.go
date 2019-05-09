package util

import (
	"os"
	"path/filepath"
)

// GetCurrentDirectory returns the current path.
func GetCurrentDirectory() string {
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		return ""
	}
	return dir
}
