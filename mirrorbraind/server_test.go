package main

import "testing"

func TestIsExtension(t *testing.T) {
	if isExtension("meta4", "/test/file.mov?torrent") {
		t.Error("Expected to be false")
	}
	if !isExtension("torrent", "/test/file.mov.torrent") {
		t.Error("Expected to be true")
	}
	if !isExtension("meta4", "/test/file.mov?meta4") {
		t.Error("Expected to be true")
	}
}
