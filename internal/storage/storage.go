package storage

import (
	"log"
	"runtime/debug"
	"strings"

	"github.com/quasilyte/gdata/v2"
)

var (
	dataManager *gdata.Manager
)

func init() {
	buildInfo, ok := debug.ReadBuildInfo()

	if !ok {
		log.Fatalln("could not get the build info")
	}

	pathParts := strings.Split(buildInfo.Main.Path, "/")

	appName := pathParts[len(pathParts)-1]

	var err error
	dataManager, err = gdata.Open(gdata.Config{
		AppName: appName,
	})

	if err != nil {
		log.Fatal(err)
	}
}
