package mirrorbrain

import "libs/database"

type ServerSelection struct {
	Chosen  MirrorbrainServer
	FoundIn string
}

func ChooseServer(requestFile RequestFile, requestIp string, servers database.Servers) ServerSelection {
	geoInfo := GeoLookup(requestIp)
	mirrorbrainServers := filterServers(requestFile, servers)
	if geoInfo.HasCoords() {
		mirrorbrainServers.CalculateDistance(geoInfo)
	}

	// TODO smarter lists?
	//serverList := ServerList{SamePrefix: make([]MirrorbrainServer, 0, len(mirrorbrainServers))}
	serverList := new(ServerList)

	// TODO use search prio to exit early if only a redirect is required, don't build all lists
	serverList.BuildServerLists(mirrorbrainServers, geoInfo)

	// TODO only sort for metalink/mirrorlist
	if geoInfo.HasCoords() {
		serverList.SortByDistance()
	} else {
		serverList.SortByRank()
	}

	serverSelection := selectServer(*serverList, geoInfo)
	return serverSelection
}

// TODO servers should not contain illegal entries (i.e. null values)
func filterServers(requestFile RequestFile, servers database.Servers) (mirrorbrainServers MirrorbrainServers) {
	for _, server := range servers {
		//  skip mirror because of server.FileMaxsize
		if server.FileMaxsize >= requestFile.StatInfo.Size() {
			var mbServer = MirrorbrainServer{server, 0, 0}
			mirrorbrainServers = append(mirrorbrainServers, mbServer)
		}
	}
	return mirrorbrainServers
}

// since lists are already sorted, just select first one
func selectServer(serverList ServerList, geoInfo *GeoInfo) (serverSelection ServerSelection) {
	if len(serverList.SamePrefix) > 0 {
		serverSelection.FoundIn = "prefix"
		// TODO Does mirrorbrain just duplicates logic from sort, or is find_closest_dist an improvement
		// since it respects Rank?
		serverSelection.Chosen = serverList.SamePrefix[0]
	} else if len(serverList.SameAsn) > 0 {
		serverSelection.FoundIn = "AS"
		serverSelection.Chosen = serverList.SameAsn[0]
	} else if len(serverList.SameCountry) > 0 {
		serverSelection.Chosen = serverList.SameCountry[0]
		if serverSelection.Chosen.Country == geoInfo.CountryCode {
			serverSelection.FoundIn = "country"
		} else {
			serverSelection.FoundIn = "other_country"
		}
	} else if len(serverList.SameRegion) > 0 {
		serverSelection.FoundIn = "region"
		serverSelection.Chosen = serverList.SameRegion[0]
	} else if len(serverList.Elsewhere) > 0 {
		serverSelection.FoundIn = "other"
		serverSelection.Chosen = serverList.Elsewhere[0]
	}
	return serverSelection
}
