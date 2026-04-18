package file 

import (
	"os"
	"path/filepath"
)
func GetCurPath() string {
	exe, err := os.Executable()
	if err != nil {
		return ""
	}
	return filepath.Dir(exe)
}
func GetCWD() string {
	t, err := os.Getwd()
	if err != nil {
		return ""
	}
	return t
}
