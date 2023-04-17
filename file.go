package gFile

import (
	"os"
)

func IsExists(path string) bool {
	_, err := os.Stat(path)
	if err != nil {
		if os.IsExist(err) {
			return true
		}
		return false
	}
	return true
}

func CreateDirIfNotExists(dir string, perm os.FileMode) error {
	exist := IsExists(dir)
	if !exist {
		err := os.Mkdir(dir, perm)
		if err != nil {
			return err
		}
	}

	return nil
}

func IsDir(path string) bool {
	s, err := os.Stat(path)
	if err != nil {
		return false
	}
	return s.IsDir()
}

func IsFile(path string) bool {
	return !IsDir(path)
}
