package main

import (
	"fmt"
	"touch-grabber/pkg/data"
	"touch-grabber/pkg/versions"
)

func main() {
	config := versions.GetConfig()
	buildVersion := versions.GetBuildVersion()
	//appVersion := versions.GetApplication().Results[0].Version
	//assetsVersion := versions.GetAssets().AssetsVersion
	//staticDataVersion := versions.GetAssets().StaticDataVersion
	//assetsUrl := config.AssetsURL
	dataUrl := config.DataURL
	//fmt.Println("Build Version: ", buildVersion)
	//fmt.Println("Application Version: ", appVersion)
	//fmt.Println("Assets Version: ", assetsVersion)
	//fmt.Println("Static Data Version: ", staticDataVersion)
	//fmt.Println("Assets URL: ", assetsUrl)
	fmt.Println("Data URL: ", dataUrl)
	//data.FetchDataMap(dataUrl, "fr", buildVersion, "Jobs")
	data.FetchAllDataMap(dataUrl, "fr", buildVersion)
}
