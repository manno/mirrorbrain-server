package mirrorbrain

import (
	"fmt"
	"github.com/manno/mirrorbrain-server/database"
	"os"
	"testing"
)

const absPath = "/tmp/test"
const requestPath = "/test"

func TestRequestFile(t *testing.T) {
	requestFile := createRequestFile()
	if requestFile.Path() != "test" {
		t.Error("RequestFile does not return path", requestFile.Path())
	}
}

func TestPrepare(t *testing.T) {
	mbServers := filterServers(createRequestFile(), createServers())
	mbServers.Prepare(createGeoInfo())
	fmt.Println(mbServers)
}

func TestFilterServers(t *testing.T) {
	mbServers := filterServers(createRequestFile(), createServers())
	fmt.Println(mbServers)
	if len(mbServers) != 2 {
		t.Error("Fails to filter one server and return 2 others", mbServers)
	}
}

func createGeoInfo() (geoInfo *GeoInfo) {
	return &GeoInfo{Latitude: 47.367, Longitude: 8.55}
}

func createServers() database.Servers {
	var servers = database.Servers{
		database.Server{FileMaxsize: -1},
		database.Server{Lat: 50.933, Lng: 6.95},
		database.Server{Lat: 49.871, Lng: 8.649},
	}
	return servers
}

func createRequestFile() RequestFile {
	stat := createTestfile(absPath)
	return RequestFile{requestPath, stat}
}

func createTestfile(absPath string) os.FileInfo {
	fo, err := os.Create(absPath)
	if err != nil {
		panic(err)
	}
	fo.Close()

	stat, err := os.Stat(absPath)
	if err != nil {
		panic(err)
	}

	return stat
}
