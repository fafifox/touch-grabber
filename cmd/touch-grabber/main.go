package main

import (
	"fmt"
	"touch-grabber/pkg/data"
	"touch-grabber/pkg/versions"
)

func main() {
	config := versions.GetConfig()
	buildVersion := versions.GetBuildVersion()
	dataUrl := config.DataURL
	assetsUrl := config.AssetsURL
	fmt.Println("Assets URL: ", assetsUrl)
	fmt.Println("Data URL: ", dataUrl)
	fmt.Println("Build Version: ", buildVersion)

	//data.FetchDataMap(dataUrl, "fr", buildVersion, "./", "Jobs")
	//data.FetchDataMap(dataUrl, "en", buildVersion, "./", "Jobs")
	data.FetchAllDataMap(dataUrl, buildVersion, ".")
}
