package main

import (
	"flag"
	"os"

	"github.com/DMAudio/panelBackend/common/config"
	"github.com/DMAudio/panelBackend/internal/api_http"
	"github.com/DMAudio/panelBackend/internal/service"
	"github.com/DMAudio/panelBackend/package/libzap"
	"github.com/DMAudio/panelBackend/package/log"
	"github.com/DMAudio/panelBackend/package/utils"
)

var (
	// LastCommit 版本commit id
	LastCommit = "No Commit hash Provided"

	// Branch 版本构建分支
	Branch = "No Branch Provided"

	// BuildTime 版本构建时间
	BuildTime = "No Build Time Provided"

	// GoVersion 版本构建时的go编译器版本号
	GoVersion = "No Go Version Provided"
)

func init() {
	config.RegisterFlags()

	checkVersion := flag.Bool("version", false, "check the version")
	flag.Parse()

	if *checkVersion {
		log.Printf("Branch: %s\n", Branch)
		log.Printf("Last Commit: %s\n", LastCommit)
		log.Printf("Build Time: %s\n", BuildTime)
		log.Printf("Go Version: %s\n", GoVersion)
		os.Exit(0)
	}
}

func main() {
	flushOutput := libzap.Setup()

	config.Load()
	api_http.Load()

	var waitShutdown = make(chan bool)
	utils.RuntimeHandleQuitSignal(func() {
		service.Shutdown()
		waitShutdown <- true
	})

	go service.Serve()
	<-waitShutdown

	_ = flushOutput()
}
