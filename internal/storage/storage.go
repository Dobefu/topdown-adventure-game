// Package storage handles the saving and loading of game data.
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

// GetOption retrieves an option from the saved storage.
func GetOption[T []byte | string | int | bool](option string, defaultValue T) (value T, err error) {
	val, err := dataManager.LoadObjectProp("options", option)

	switch any(value).(type) {
	case string:
		return any(val).(T), err

	case int:
		if err != nil {
			return defaultValue, err
		}

		valStr := string(val)

		// If the string is empty, there is probably no value to get.
		// In this case, let's just return the default value.
		if valStr == "" {
			return defaultValue, nil
		}

		parsedValue, err := strconv.ParseInt(valStr, 10, 16)

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

// SetOption sets an option in saved storage.
func SetOption(option string, value string) (err error) {
	return dataManager.SaveObjectProp("options", option, []byte(value))
}
