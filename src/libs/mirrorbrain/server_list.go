package mirrorbrain

import (
	"fmt"
	"sort"
)

type ServerList struct {
	All         MirrorbrainServers
	SamePrefix  MirrorbrainServers
	SameAsn     MirrorbrainServers
	SameCountry MirrorbrainServers
	SameRegion  MirrorbrainServers
	Elsewhere   MirrorbrainServers
}

func (serverList *ServerList) BuildServerLists(allServers MirrorbrainServers, geoInfo *GeoInfo) {
	serverList.All = allServers

	for _, server := range allServers {
		fmt.Println(server)
		found := false

		if server.Prefix == geoInfo.Prefix {
			serverList.SamePrefix = append(serverList.SamePrefix, server)
			found = true
		}
		if server.Asn == geoInfo.AS {
			serverList.SameAsn = append(serverList.SameAsn, server)
			found = true
		}
		if server.Country == geoInfo.CountryName || server.HasWildcardCountry() {
			serverList.SameCountry = append(serverList.SameCountry, server)
			found = true
		}
		//  TODO fallback country codes
		if server.Region == geoInfo.RegionName && server.AcceptsForeign() {
			serverList.SameRegion = append(serverList.SameRegion, server)
			found = true
		}
		if !found && server.Worldwide() {
			serverList.Elsewhere = append(serverList.Elsewhere, server)
		}
	}
	// fallback logic: same country empty? use fallback country
	// fallback: empty list of all servers use cfg->fallbacks
}

func (serverList *ServerList) SortByDistance() {
	sort.Sort(ServersByDistance(serverList.SamePrefix))
	sort.Sort(ServersByDistance(serverList.SameAsn))
	sort.Sort(ServersByDistance(serverList.SameCountry))
	sort.Sort(ServersByDistance(serverList.SameRegion))
	sort.Sort(ServersByDistance(serverList.Elsewhere))
}

func (serverList *ServerList) SortByRank() {
	sort.Sort(ServersByRank(serverList.SamePrefix))
	sort.Sort(ServersByRank(serverList.SameAsn))
	sort.Sort(ServersByRank(serverList.SameCountry))
	sort.Sort(ServersByRank(serverList.SameRegion))
	sort.Sort(ServersByRank(serverList.Elsewhere))
}
