package main

import (
	"os"
	"path"
	"strings"
)

func webrootPath(webRoot string, urlPath string) string {
	return path.Join(webRoot, urlPath)
}

func removeSlash(path string) string {
	if strings.Index(path, "/") == 0 {
		return path[1:len(path)]
	}
	return path
}

func addTrailingSlash(host string) string {
	if host[len(host)-1] != '/' {
		return host + "/"
	}
	return host
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
