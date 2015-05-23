package main

import (
	"html/template"
	"io"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

const DirDateFormat = "02-Jan-2006 15:04"

type dirEntry struct {
	Name    string
	ModTime string
	Size    string
}

type dirListing struct {
	Entries    []dirEntry
	TopDir     bool
	CurrentDir string
}

const fileIndexTemplate = `
<html>
<head><title>Index of {{.CurrentDir}}</title></head>
<body bgcolor="white">
<h1>Index of {{.CurrentDir}}</h1>
<hr><pre>{{ if not .TopDir }}<a href="../">../</a>{{ end }}
{{range .Entries}}<a href="{{.Name}}">{{.Name}}</a>        {{.ModTime}}        {{.Size}}
{{end}}</pre><hr>
</body>
</html>
`

// TODO support WEBDAV propfind for better directory listings
func printDirectoryList(w io.Writer, path string, requestPath string) {
	if strings.Index(path, config.WebRoot) != 0 {
		log.Fatal("directory listing outside webroot requested")
	}
	files, _ := ioutil.ReadDir(path)
	listing := new(dirListing)
	listing.TopDir = pathEqual(path, config.WebRoot)
	listing.CurrentDir = requestPath
	for _, f := range files {
		entry := new(dirEntry)
		if f.IsDir() {
			entry.Name = f.Name() + "/"
			entry.Size = "-"
		} else {
			entry.Name = f.Name()
			entry.Size = strconv.FormatInt(f.Size(), 10)
		}
		entry.ModTime = f.ModTime().Format(DirDateFormat)
		listing.Entries = append(listing.Entries, *entry)
	}

	t, _ := template.New("file index template").Parse(fileIndexTemplate)
	t.Execute(w, listing)
}
