package main

import (
	"flag"
	"fmt"
)

var (
	isVersion    bool
	appName      string
	buildTime    string
	buildVersion string
	gitCommitID  string
)

func init() {

	flag.BoolVar(&isVersion, "version", false, "版本信息")
	flag.Parse()

}

func main() {
	if isVersion {
		fmt.Printf("app_name: %s \r\n", appName)
		fmt.Printf("build_time: %s \r\n", buildTime)
		fmt.Printf("build_version: %s \r\n", buildVersion)
		fmt.Printf("git_commit_id: %s \r\n", gitCommitID)
		return
	}

	fmt.Println("hello word!")
}

// 编译命令
// go build -ldflags "-X main.appName=ldflagsDemo -X main.buildTime=` date +%Y-%m-%d,%H:%M:%S` -X main.buildVersion=1.1.0 -X main.gitCommitID=` git rev-parse HEAD`"
