package main

import (
	"os"
	"path"
)

func webrootPath(webRoot string, urlPath string) string {
	return path.Join(webRoot, urlPath)
}

func isDir(path string) bool {
	fileInfo, err := os.Stat(path)
	if err != nil {
		return false
	}
	return fileInfo.IsDir()
}

func pathEqual(p1 string, p2 string) bool {
	return path.Clean(p1) == path.Clean(p2)
}
