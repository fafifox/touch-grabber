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

func FetchAssetMap(dataUrl string, buildVersion string, savePath string) {
	url := dataUrl + "/assetMap.json"
	data := httputils.GetRequestBody(url)
	assets := assetMap{}
	err := json.Unmarshal(data, &assets)
	if err != nil {
		log.Println(err)
	}
	for k := range assets.Files {
		log.Println(assets.Files[k].Filename)
		data := httputils.GetRequestBody(dataUrl + "/" + assets.Files[k].Filename)
		err := SaveAsset(data, assets.Files[k].Filename, buildVersion, savePath)
		if err != nil {
			log.Println(err)
		}
	}
}

func SaveAsset(assetData []byte, assetName string, buildVersion string, path string) error {
	mainFolder := "/" + buildVersion + "/" + assetName // 1.45.4/assets/ui/preloadMap/screenshot.png
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
