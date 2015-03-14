package mirrorbrain

import (
	"os"
	"strings"
)

type RequestFile struct {
	RequestPath string
	StatInfo    os.FileInfo
}

func (rf RequestFile) Path() string {
	return removeSlash(rf.RequestPath)
}

func removeSlash(path string) string {
	if strings.Index(path, "/") == 0 {
		return path[1:len(path)]
	}
	return path
}
