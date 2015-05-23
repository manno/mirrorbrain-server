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

// format request, i.e: ?meta4 or .torrent
func (rf RequestFile) HasExtension(extension string) bool {
	pos := strings.LastIndex(rf.RequestPath, extension)
	if pos > -1 && pos == len(rf.RequestPath)-len(extension) {
		if rf.RequestPath[pos-1] == '.' || rf.RequestPath[pos-1] == '?' {
			return true
		}
	}
	return false
}

func removeSlash(path string) string {
	if strings.Index(path, "/") == 0 {
		return path[1:len(path)]
	}
	return path
}
