package mirrorbrain

import (
	"libs/database"
	"log"

	"github.com/abh/geoip"
)

var file = "/usr/share/GeoIP/GeoIP.dat"
var geoIP *geoip.GeoIP

type GeoInfo struct {
	RequestIp     string
	AS            string
	Prefix        string
	CountryName   string
	CountryCode   string
	ContinentCode string
	Latitude      float32
	Longitude     float32
	RegionCode    string
	RegionName    string
}

func GeoSetup() {
	gi, err := geoip.Open(file)
	if err != nil {
		log.Fatal("Could not open GeoIP database\n")
	}
	geoIP = gi
}

// TODO ipv6
func GeoLookup(ip string) *GeoInfo {
	geoInfo := new(GeoInfo)
	geoInfo.RequestIp = ip

	// lookup as number for request ip
	pfx, as, err := database.SelectAsn(ip)
	if err != nil {
		geoInfo.AS = as
		geoInfo.Prefix = pfx
	}

	// lookup geoip for request ip
	geoIpRecord := geoIP.GetRecord(ip)
	geoInfo.CountryName = geoIpRecord.CountryName
	geoInfo.CountryCode = geoIpRecord.CountryCode
	geoInfo.CountryName = geoIpRecord.CountryName
	geoInfo.ContinentCode = geoIpRecord.ContinentCode
	geoInfo.Latitude = geoIpRecord.Latitude
	geoInfo.Longitude = geoIpRecord.Longitude
	geoInfo.RegionCode = geoIpRecord.Region
	//geoInfo.RegionName = geoIP.GetRegionName(geoIpRecord.CountryCode, geoIpRecord.Region)

	return geoInfo
}
