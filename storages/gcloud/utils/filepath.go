package utils

import (
	"path"
	"path/filepath"
	"strings"
)

func GetFileName(basename string) string {
	return filepath.Base(basename)
}

func GetFileExtension(basename string) string {
	return strings.TrimLeft(filepath.Ext(basename), ".")
}

func GetChildrenFolder(rootFolder, folderPath string) []string {
	dir := path.Dir(rootFolder)
	folder := strings.TrimSpace(strings.Replace(dir, folderPath, "", 1))
	if folder == "" {
		return nil
	}
	result := strings.Split(strings.Replace(folder, "/", "", 1), "/")
	return result
}

func GetFolder(basename string) string {
	dir := path.Dir(basename)
	return dir
}
