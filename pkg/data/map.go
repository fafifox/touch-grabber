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
			err := FetchDataMap(dataUrl, lang, buildVersion, path, typeName)
			if err != nil {
				log.Println(err)
			}
		}
	}
}

func FetchDataMap(dataUrl string, langage string, buildVersion string, path string, typeName string) error {
	url := dataUrl + "/data/map?lang=" + langage + "&v=" + buildVersion
	requestBody, err := json.Marshal(map[string]string{
		"class": typeName,
	})
	if err != nil {
		return err
	}
	resp, err := http.Post(url, "application/json", bytes.NewBuffer(requestBody))
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	log.Printf("[%d] %s Class: %s\n", resp.StatusCode, resp.Request.URL, typeName)
	if resp.StatusCode == 200 {
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return err
		}
		dataMapData := dataMapData{
			Data: body,
			Name: typeName + ".json",
		}
		err = SaveDataMap(&dataMapData, buildVersion, path, langage)
		if err != nil {
			return err
		}
		return nil
	}
	return nil
}

func SaveDataMap(data *dataMapData, buildVersion string, path string, lang string) error {
	mainFolder := "/" + buildVersion + "/data/map"
	langFolder := "/" + lang
	path = path + mainFolder
	pathLang := path + langFolder
	if _, err := os.Stat(path); os.IsNotExist(err) {
		err = os.MkdirAll(path, os.ModePerm)
	}
	if _, err := os.Stat(pathLang); os.IsNotExist(err) {
		err = os.MkdirAll(pathLang, os.ModePerm)
	}
	file, err := os.Create(pathLang + "/" + data.Name)
	if err != nil {
		return err
	}
	defer file.Close()
	_, err = file.Write(data.Data)
	return err
}
