package data

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

type dataMapDictionary struct {
	Data []byte
	Name string
}

func FetchAllDataDictionary(dataUrl string, buildVersion string, path string) {
	for _, lang := range langs {
		err := FetchDataDictionary(dataUrl, lang, buildVersion, path)
		if err != nil {
			log.Println(err)
		}
	}
}

func FetchDataDictionary(dataUrl string, langage string, buildVersion string, path string) error {
	url := dataUrl + "/data/dictionary?lang=" + langage + "&v=" + buildVersion
	requestBody, err := json.Marshal(map[string]string{
		"lang": langage,
	})
	if err != nil {
		return err
	}
	resp, err := http.Post(url, "application/json", bytes.NewBuffer(requestBody))
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	log.Printf("[%d] %s\n", resp.StatusCode, resp.Request.URL)
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	dataMapDictionary := dataMapDictionary{
		Data: body,
		Name: langage + ".json",
	}
	err = SaveDataDictionary(&dataMapDictionary, buildVersion, path, langage)
	if err != nil {
		return err
	}
	return nil
}

func SaveDataDictionary(data *dataMapDictionary, buildVersion string, path string, lang string) error {
	mainFolder := "/" + buildVersion + "/data/dictionary"
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
