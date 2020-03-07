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
		FetchDataDictionary(dataUrl, lang, buildVersion, path)
	}
}

func FetchDataDictionary(dataUrl string, langage string, buildVersion string, path string) {
	url := dataUrl + "/data/dictionary?lang=" + langage + "&v=" + buildVersion
	requestBody, err := json.Marshal(map[string]string{
		"lang": langage,
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
	dataMapDictionary := dataMapDictionary{
		Data: body,
		Name: langage + ".json",
	}
	err = SaveDataDictionary(&dataMapDictionary, buildVersion, path, langage)
	if err != nil {
		log.Println(err)
	}
	log.Println("Saved: ", dataMapDictionary.Name)
}

func SaveDataDictionary(data *dataMapDictionary, buildVersion string, path string, lang string) error {
	mainFolder := "/dataDictionary-v" + buildVersion
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
