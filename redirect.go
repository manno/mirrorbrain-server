package main

import (
	"log"
	"net/http"
	"os"
	"path"
)

const WebRoot = "/usr/share/pixmaps"

func confPort() string {
	var port = os.Getenv("PORT")
	if port == "" {
		return "8080"
	}
	return port
}

func viewHandler(w http.ResponseWriter, r *http.Request) {
	var requestPath = path.Join(WebRoot, r.URL.Path)
	log.Println(r.RemoteAddr, r.URL.Path, requestPath)
	if _, err := os.Stat(requestPath); err == nil {
		sendRedirect(w, r, path.Clean(r.URL.Path))
	} else {
		http.NotFound(w, r)
	}
}

func sendRedirect(w http.ResponseWriter, r *http.Request, path string) {
	//fmt.Fprintf(w, "%s", path)
	http.Redirect(w, r, "http://otherhost"+path, http.StatusFound)
}

func main() {
	var port = ":" + confPort()
	log.Printf("Listening on %s", confPort())

	http.HandleFunc("/", viewHandler)
	log.Fatal(http.ListenAndServe(port, nil))
}
