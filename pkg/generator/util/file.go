package util

import "os"

func MakeDir(dir string) error {
	return os.MkdirAll(dir, 0755)
}

func MakeFile(file string) (*os.File, error) {
	return os.Create(file)
}
