package storage

import (
	"log"
	"runtime/debug"
	"strconv"
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

func GetOption[T []byte | string | int | bool](option string, defaultValue T) (value T, err error) {
	val, err := dataManager.LoadObjectProp("options", option)

	switch any(value).(type) {
	case string:
		return any(val).(T), err

	case int:
		if err != nil {
			return defaultValue, err
		}

		parsedValue, err := strconv.ParseInt(string(val), 10, 16)

		if err != nil {
			return defaultValue, err
		}

		return any(int(parsedValue)).(T), nil

	case bool:
		return any(string(val) == "true").(T), nil

	default:
		return value, err
	}
}

func SetOption(option string, value string) (err error) {
	return dataManager.SaveObjectProp("options", option, []byte(value))
}
