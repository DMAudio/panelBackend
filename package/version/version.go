package version

import (
	"fmt"
	"runtime"

	"github.com/DMAudio/panelBackend/package/log"
)

// Info contains versioning information.
type Info struct {
	GitTag       string `json:"gitTag"`
	GitCommit    string `json:"gitCommit"`
	GitTreeState string `json:"gitTreeState"`
	BuildDate    string `json:"buildDate"`
	GoVersion    string `json:"goVersion"`
	Compiler     string `json:"compiler"`
	Platform     string `json:"platform"`
}

// String returns info as a human-friendly version string.
func (info Info) String() string {
	return info.GitTag
}

func Get() Info {
	return Info{
		GitTag:       gitTag,
		GitCommit:    gitCommit,
		GitTreeState: gitTreeState,
		BuildDate:    buildDate,
		GoVersion:    runtime.Version(),
		Compiler:     runtime.Compiler,
		Platform:     fmt.Sprintf("%s/%s", runtime.GOOS, runtime.GOARCH),
	}
}

func Print() {
	info := Get()
	log.Info("GitTag:", info.GitTag)
	log.Info("GitCommit:", info.GitCommit)
	log.Info("GitTreeState:", info.GitTreeState)
	log.Info("BuildDate:", info.BuildDate)
	log.Info("GoVersion:", info.GoVersion)
	log.Info("Compiler:", info.Compiler)
	log.Info("Platform:", info.Platform)

	log.Info()
}
