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
	fmt.Println("Data URL: ", dataUrl)
	//data.FetchDataMap(dataUrl, "fr", buildVersion, "./", "Jobs")
	data.FetchAllDataMap(dataUrl, "fr", buildVersion, "./")
}
