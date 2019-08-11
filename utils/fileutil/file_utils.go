package fileutil

import (
	"os"
	"path/filepath"
)

func GetFileNameList(root string) ([]string, error) {
	var files []string
	err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if info == nil {
			return err
		}
		if !info.IsDir() {
			files = append(files, info.Name())
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	return files, nil
}
