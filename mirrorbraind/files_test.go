package main

import (
	"os"
	"testing"
)

func TestPrintDirectoryList(t *testing.T) {
	config.Setup()
	printDirectoryList(os.Stdout, "/srv/ftp/congress/2012/mp4-h264-LQ-iProd", "/congress/2012/mp4")
	printDirectoryList(os.Stdout, "/srv/ftp/", "/")
	printDirectoryList(os.Stdout, "/srv/ftp", "/")

}
