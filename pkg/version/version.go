package version

import "github.com/prometheus/common/version"

var (
	VERSION   = "unknown"
	BUILDTIME = "unknown"
)

func SetPrometheusVersion() {
	version.Version = VERSION
	version.BuildDate = BUILDTIME
}
