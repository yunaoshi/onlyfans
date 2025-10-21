package version

//
// Versioning: Use semantic versioning (MAJOR.MINOR.PATCH, e.g. 1.2.3).
// Tag releases with the version (e.g. v1.2.3) and update the version for each release.
// Show only the version/tag to users, not commit hashes.

//
// Build Guidelines:
// -----------------
// To set version and tag at build time, use Go ldflags:
//
//   go build -ldflags "-X 'onlyfans/internal/version.Version=1.2.3' -X 'onlyfans/internal/version.Tag=v1.2.3'"
//
// You can automate this in CI/CD:
//   VERSION=$(git describe --tags --abbrev=0)
//   go build -ldflags "-X 'onlyfans/internal/version.Version=$VERSION' -X 'onlyfans/internal/version.Tag=$VERSION'"
//
// The version string shown to users will only display the tag (if set) or the version. Commit hash is not included in user output.
//
// If not set, the package will attempt to get these values from git at runtime (if available).

import (
	"os/exec"
	"strings"
)

// Set by build flags, e.g.:
// go build -ldflags "-X 'onlyfans/internal/version.Version=1.2.3' -X 'onlyfans/internal/version.CommitHash=abc1234' -X 'onlyfans/internal/version.Tag=v1.2.3'"
var (
	version    = "0.0.0" // fallback version
	commitHash = "unknown"
	tag        = ""
)

func String() string {
	if version == "0.0.0" || version == "" {
		// Try to get git tag at runtime
		out, err := exec.Command("git", "describe", "--tags", "--abbrev=0").Output()
		if err == nil {
			t := strings.TrimSpace(string(out))
			if t != "" {
				return t
			}
		}
		return "unknown (version not set at build time)"
	}
	return version
}
