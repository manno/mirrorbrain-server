package mirrorbrain

import (
	"libs/database"
	"log"

	"github.com/abh/geoip"
)

var file = "/usr/share/GeoIP/GeoIPCity.dat"
var geoIP *geoip.GeoIP

type GeoInfo struct {
	RequestIp     string
	AS            string
	Prefix        string
	CountryName   string
	CountryCode   string
	ContinentCode string
	Latitude      float64
	Longitude     float64
	RegionCode    string
	RegionName    string
}

func (geoInfo GeoInfo) HasCoords() bool {
	return geoInfo.Latitude != 0 && geoInfo.Longitude != 0
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
	if geoIpRecord == nil {
		log.Println("Failed to acquire GeoIP record for:", ip)
		return geoInfo
	}

	// TODO convervt empty values to --
	geoInfo.CountryName = geoIpRecord.CountryName
	geoInfo.CountryCode = geoIpRecord.CountryCode
	geoInfo.CountryName = geoIpRecord.CountryName
	geoInfo.ContinentCode = geoIpRecord.ContinentCode
	geoInfo.Latitude = float64(geoIpRecord.Latitude)
	geoInfo.Longitude = float64(geoIpRecord.Longitude)
	geoInfo.RegionCode = geoIpRecord.Region
	//geoInfo.RegionName = geoIP.GetRegionName(geoIpRecord.CountryCode, geoIpRecord.Region)

	return geoInfo
}
