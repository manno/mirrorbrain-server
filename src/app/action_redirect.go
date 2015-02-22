package main

import (
	"libs/database"
	"log"
	"net/http"
)

func sendRedirect(w http.ResponseWriter, r *http.Request, path string) {
	// TODO decline redirect for:
	//   small files cfg->min_sie
	//   excluded files cfg->exclude_filemask
	//   excluded source ips cfg->exclude_ips
	//   excluded network cfg->exclude_networks
	//   excluded mime type cfg->exclude_mime
	//   excluded user agent cfg->exclude_agents
	// TODO since we don't send directly, what happens?

	path = removeSlash(path)
	servers, err := database.FindServers(path)
	if err != nil {
		http.Error(w, "internal error", http.StatusInternalServerError)
		return
	}
	dumpServers(servers)
	server := mirrorbrain.ChooseServer(servers)
	http.Redirect(w, r, addTrailingSlash(server.BaseUrl)+path, http.StatusFound)
}

func dumpServers(servers database.Servers) {
	for _, server := range servers {
		log.Printf("%d %s (%s)\n", server.Score, server.Identifier, server.BaseUrl)
	}
}
