package main

import (
	"libs/database"
	"log"
	"net/http"
	"strings"
)

func sendRedirect(w http.ResponseWriter, r *http.Request, path string) {
	path = removeSlash(path)
	servers, err := database.FindServers(path)
	if err != nil {
		http.Error(w, "internal error", http.StatusInternalServerError)
		return
	}
	dumpServers(servers)
	server := selectServer(servers)
	http.Redirect(w, r, addTrailingSlash(server.BaseUrl)+path, http.StatusFound)
}

func dumpServers(servers database.Servers) {
	for _, server := range servers {
		log.Printf("%d %s (%s)\n", server.Score, server.Identifier, server.BaseUrl)
	}
}

func selectServer(servers database.Servers) database.Server {
	max := servers[0]

	for _, server := range servers {
		if server.Score > max.Score {
			max = server
		}
	}
	return max
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
