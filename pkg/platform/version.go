package platform

import (
	"strings"

	"github.com/blang/semver/v4"
)

var (
	MinimumVersionTag = "v4.3.4"
	MinimumVersion    = semver.MustParse(strings.TrimPrefix(MinimumVersionTag, "v"))
)

// LatestCompatibleVersion returns the latest compatible version of test-release-oss
// This is a simplified version for testing purposes
func LatestCompatibleVersion() string {
	return MinimumVersionTag
}
