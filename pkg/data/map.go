package data

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

//https://proxyconnection.touch.dofus.com/data/map?lang=fr&v=1.45.4


//Invoke-WebRequest -Uri "https://proxyconnection.touch.dofus.com/data/map?lang=fr&v=1.45.4"
//-Method "POST"
//-Headers @{
//	"Origin"="file://";
//	"Accept-Encoding"="gzip, deflate, br";
//	"Accept-Language"="fr";
//	"User-Agent"="Mozilla/5.0 (Linux; Android 5.1.1; SHIELD Tablet Build/LMY48C; wv) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/61.0.3163.113 Mobile Safari/537.36";
//	"Accept"="*/*";
//	"Cookie"="io=f8DC6oCt7ST3FEA6BTAG"
//	}
//-ContentType "application/json"
//-Body "{`"class`":`"TYPE`"}"
//-outfile "TYPE.json"

type dataMapData struct {
	Data      []byte
	Name      string
}

func FetchAllDataMap(dataUrl string, langage string, buildVersion string) {
	for _, typeName := range typeNames {
		FetchDataMap(dataUrl, langage, buildVersion, typeName)
	}
}

func FetchDataMap(dataUrl string, langage string, buildVersion string, typeName string) {
	url := dataUrl + "/data/map?lang=" + langage + "&v=" + buildVersion
	requestBody, err := json.Marshal(map[string]string{
		"class" : typeName,
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
	err = SaveDataMap(&dataMapData)
	if err != nil {
		log.Println(err)
	}
	log.Println("Saved: ", dataMapData.Name)
}

func SaveDataMap(data *dataMapData) error {
	path := "output"
	if _, err := os.Stat(path); os.IsNotExist(err) {
		err = os.Mkdir(path, os.ModePerm)
	}
	file, err := os.Create(path + "/" + data.Name)
	if err != nil {
		return err
	}
	defer file.Close()
	_, err = file.Write(data.Data)
	return err
}