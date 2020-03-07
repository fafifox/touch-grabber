package data

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

type dataMapData struct {
	Data []byte
	Name string
}

func FetchAllDataMap(dataUrl string, buildVersion string, path string) {
	for _, lang := range langs {
		for _, typeName := range typeNames {
			FetchDataMap(dataUrl, lang, buildVersion, path, typeName)
		}
	}
}

func FetchDataMap(dataUrl string, langage string, buildVersion string, path string, typeName string) {
	url := dataUrl + "/data/map?lang=" + langage + "&v=" + buildVersion
	requestBody, err := json.Marshal(map[string]string{
		"class": typeName,
	})
	if err != nil {
		log.Println(err)
	}
	resp, err := http.Post(url, "application/json", bytes.NewBuffer(requestBody))
	if err != nil {
		log.Println(err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println(err)
	}
	dataMapData := dataMapData{
		Data: body,
		Name: typeName + ".json",
	}
	err = SaveDataMap(&dataMapData, buildVersion, path, langage)
	if err != nil {
		log.Println(err)
	}
	log.Println("Saved: ", dataMapData.Name)
}

func SaveDataMap(data *dataMapData, buildVersion string, path string, lang string) error {
	mainFolder := "/dataMap-v" + buildVersion
	langFolder := "/" + lang
	path = path + mainFolder
	pathLang := path + langFolder
	if _, err := os.Stat(path); os.IsNotExist(err) {
		err = os.Mkdir(path, os.ModePerm)
	}
	if _, err := os.Stat(pathLang); os.IsNotExist(err) {
		err = os.Mkdir(pathLang, os.ModePerm)
	}
	file, err := os.Create(pathLang + "/" + data.Name)
	if err != nil {
		return err
	}
	defer file.Close()
	_, err = file.Write(data.Data)
	return err
}
