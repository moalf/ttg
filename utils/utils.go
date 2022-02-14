// Utility package
package utils

import "os"

// FileExists checks for existence of a file path. It returns true if it does exist, otherwise returns false.
func FileExists(path string) (bool, error) {
	stat, err := os.Stat(path)
	if err == nil && !stat.IsDir() && stat.Size() > 0 {
		return true, nil
	}
	return false, err
}
