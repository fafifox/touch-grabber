package data

import (
	"encoding/json"
	"github.com/farooch/touch-grabber/pkg/httputils"
	"log"
	"os"
	"path/filepath"
)

type assetMap struct {
	Files map[string]Files `json:"files"`
	Load  []interface{}    `json:"load"`
}

type Files struct {
	Filename string `json:"filename"`
	Version  string `json:"version"`
}

var (
	pathManifest       = "manifest.json"
	pathAssetMap       = "assetMap.json"
	pathBuildScript    = "build/script.js"
	pathBuildStyle     = "build/styles-native.css"
	pathPrimus         = "primus/primus.js"
	pathAssetsVersions = "assetsVersions.json"
	pathConfig         = "config.json"
)

func FetchAssetMap(dataUrl string, buildVersion string, savePath string) {
	//Files URLs
	urlManifest := dataUrl + "/" + pathManifest
	urlAssetMap := dataUrl + "/" + pathAssetMap
	urlBuildScript := dataUrl + "/" + pathBuildScript
	urlBuildStyle := dataUrl + "/" + pathBuildStyle
	urlPrimus := dataUrl + "/" + pathPrimus
	urlAssetsVersions := dataUrl + "/" + pathAssetsVersions
	urlConfig := dataUrl + "/" + pathConfig

	dataManifest := httputils.GetRequestBody(urlManifest)
	dataAssetMap := httputils.GetRequestBody(urlAssetMap)
	dataBuildScript := httputils.GetRequestBody(urlBuildScript)
	dataBuildStyle := httputils.GetRequestBody(urlBuildStyle)
	dataPrimus := httputils.GetRequestBody(urlPrimus)
	dataAssetsVersions := httputils.GetRequestBody(urlAssetsVersions)
	dataConfig := httputils.GetRequestBody(urlConfig)

	err := SaveAsset(dataManifest, pathManifest, buildVersion, savePath)
	if err != nil {
		log.Println(err)
	}
	err = SaveAsset(dataAssetMap, pathAssetMap, buildVersion, savePath)
	if err != nil {
		log.Println(err)
	}
	err = SaveAsset(dataBuildScript, pathBuildScript, buildVersion, savePath)
	if err != nil {
		log.Println(err)
	}
	err = SaveAsset(dataBuildStyle, pathBuildStyle, buildVersion, savePath)
	if err != nil {
		log.Println(err)
	}
	err = SaveAsset(dataPrimus, pathPrimus, buildVersion, savePath)
	if err != nil {
		log.Println(err)
	}
	err = SaveAsset(dataAssetsVersions, pathAssetsVersions, buildVersion, savePath)
	if err != nil {
		log.Println(err)
	}
	err = SaveAsset(dataConfig, pathConfig, buildVersion, savePath)
	if err != nil {
		log.Println(err)
	}

	assets := assetMap{}
	err = json.Unmarshal(dataAssetMap, &assets)
	if err != nil {
		log.Println(err)
	}
	for k := range assets.Files {
		data := httputils.GetRequestBody(dataUrl + "/" + assets.Files[k].Filename)
		err := SaveAsset(data, assets.Files[k].Filename, buildVersion, savePath)
		if err != nil {
			log.Println(err)
		}
	}
}

func SaveAsset(assetData []byte, assetName string, buildVersion string, path string) error {
	mainFolder := "/" + buildVersion + "/" + assetName
	fullPath := path + mainFolder
	dirName, fileName := filepath.Split(fullPath)
	if _, err := os.Stat(dirName); os.IsNotExist(err) {
		err = os.MkdirAll(dirName, os.ModePerm)
	}
	file, err := os.Create(dirName + "/" + fileName)
	if err != nil {
		return err
	}
	defer file.Close()
	_, err = file.Write(assetData)
	return err
}
