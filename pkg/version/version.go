package version

import (
	"runtime/debug"

	"github.com/prometheus/common/version"
)

var (
	VERSION = "unknown"
)

func SetPrometheusVersion() {
	version.Version = VERSION
}

func GetTime() string {
	info, ok := debug.ReadBuildInfo()
	if ok {
		for _, s := range info.Settings {
			if s.Key == "vcs.time" {
				return s.Value
			}
		}
	}
	return "unknown"
}
