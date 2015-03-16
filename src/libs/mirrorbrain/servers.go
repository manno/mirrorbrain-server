package mirrorbrain

import (
	"libs/database"
	"log"
	"math"
	"math/rand"
	"time"
)

/* the smaller, the smaller the effect of a raised prio in comparison to distance */
/* 5000000 -> mirror in 200km distance is preferred already when it has prio 100
* 1000000 -> mirror in 200km distance is preferred not before it has prio 300
* (compared to 100 as normal priority for other mirrors, and tested in
* Germany, which is a small country with many mirrors) */
const DISTANCE_PRIO = 2000000
const RAND_MAX = 32767

type MirrorbrainServer struct {
	database.Server
	Distance int64
	Rank     int
}

func (server MirrorbrainServer) HasCoords() bool {
	return server.Lat != 0 && server.Lng != 0
}

func (server *MirrorbrainServer) CalculateDistance(geoInfo *GeoInfo) {
	server.Distance = int64(math.Sqrt(math.Pow((server.Lat-geoInfo.Latitude), 2)+math.Pow((server.Lng-geoInfo.Longitude), 2)) * 1000)
}

func (server MirrorbrainServer) HasWildcardCountry() bool {
	return server.Country == "**"
}

func (server MirrorbrainServer) AcceptsForeign() bool {
	return !server.CountryOnly && !server.ASOnly && !server.PrefixOnly
}

func (server MirrorbrainServer) Worldwide() bool {
	return server.AcceptsForeign() && !server.RegionOnly
}

func (server MirrorbrainServer) RedirectUrl(path string) string {
	return addTrailingSlash(server.BaseUrl) + path
}

func addTrailingSlash(host string) string {
	if host[len(host)-1] != '/' {
		return host + "/"
	}
	return host
}

type MirrorbrainServers []MirrorbrainServer

func (servers MirrorbrainServers) Prepare(geoInfo *GeoInfo) {
	random := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := range servers {
		mbServer := &servers[i]
		log.Println(mbServer.HasCoords(), geoInfo.HasCoords())
		if mbServer.HasCoords() && geoInfo.HasCoords() {
			mbServer.CalculateDistance(geoInfo)
		}
		if mbServer.Score < 1 {
			mbServer.Score = 1
		}
		/* rank it (randomized, weighted by "score" value) */
		mbServer.Rank = random.Intn(RAND_MAX) * (RAND_MAX / mbServer.Score)
	}
}

type ServersByDistance []MirrorbrainServer

func (s ServersByDistance) Len() int {
	return len(s)
}

func (s ServersByDistance) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s ServersByDistance) Less(i, j int) bool {
	n := len(s)
	distprio := DISTANCE_PRIO / n
	a := int(s[i].Distance) + distprio/s[i].Score
	b := int(s[j].Distance) + distprio/s[j].Score
	return a < b
}

type ServersByRank []MirrorbrainServer

func (s ServersByRank) Len() int {
	return len(s)
}

func (s ServersByRank) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s ServersByRank) Less(i, j int) bool {
	return s[i].Rank < s[j].Rank
}
