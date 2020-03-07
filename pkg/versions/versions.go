package versions

import (
	"encoding/json"
	"log"
	"os"
	"regexp"
	"time"
	"touch-grabber/pkg/httputils"
)

const (
	TOUCH_PROXY_URL = "https://proxyconnection.touch.dofus.com" //get from config
	LANGAGE         = "fr"
)

var (
	TOUCH_CONFIG_URL   = TOUCH_PROXY_URL + "/config.json"
	TOUCH_BUILD_SCRIPT = TOUCH_PROXY_URL + "/build/script.js"
	TOUCH_ASSETS       = TOUCH_PROXY_URL + "/assetsVersions.json"
	APP_URL            = "https://itunes.apple.com/lookup?country=" + LANGAGE + "&id=1041406978&lang=" + LANGAGE + "&limit=1&t=" + string(time.Now().UTC().UnixNano()/1e6)
)

/**
Get application informations from the App Store
*/
func GetApplication() *Application {
	data := httputils.GetRequestBody(APP_URL)
	app := Application{}
	jsonErr := json.Unmarshal(data, &app)
	if jsonErr != nil {
		log.Fatal(jsonErr)
	}
	return &app
}

/**
get build version from script.js
https://proxyconnection.touch.dofus.com/build/script.js
*/
func GetBuildVersion() string {
	data := httputils.GetRequestBody(TOUCH_BUILD_SCRIPT)
	reg := regexp.MustCompile(`.*buildVersion=("|')([0-9]*\.[0-9]*\.[0-9]*)("|')`)
	buildVersion := reg.FindSubmatch(data)
	return string(buildVersion[2])
}

/**
Get Assets - Get and Parse then return Assets struct
https://proxyconnection.touch.dofus.com/assetsVersions.json
*/
func GetAssets() *Assets {
	data := httputils.GetRequestBody(TOUCH_ASSETS)
	assets := Assets{}
	jsonErr := json.Unmarshal(data, &assets)
	if jsonErr != nil {
		log.Fatal(jsonErr)
	}
	return &assets
}

/**
Get Config - Get and Parse then return Config struct
https://proxyconnection.touch.dofus.com/config.json
*/
func GetConfig() *Config {
	data := httputils.GetRequestBody(TOUCH_CONFIG_URL)
	config := Config{}
	jsonErr := json.Unmarshal(data, &config)
	if jsonErr != nil {
		log.Fatal(jsonErr)
	}
	return &config
}

func GetVersions(path string) {
	//Fetch versions
	appVersion := GetApplication().Results[0].Version
	buildVersion := GetBuildVersion()
	assetsVersion := GetAssets().AssetsVersion
	config := GetConfig()
	assetsUrl := config.AssetsURL
	dataUrl := config.DataURL
	//Display results
	log.Println("App Version: ", appVersion)
	log.Println("Build Version: ", buildVersion)
	log.Println("Assets Version: ", assetsVersion)
	log.Println("Assets URL: ", assetsUrl)
	log.Println("Data URL: ", dataUrl)
	//Parse results in json
	client, err := json.Marshal(Client{
		AppVersion:    appVersion,
		BuildVersion:  buildVersion,
		AssetsVersion: assetsVersion,
		AssetsURL:     assetsUrl,
		DataURL:       dataUrl,
	})
	if err != nil {
		log.Println(err)
	}
	//Save results
	err = SaveData(buildVersion, path, client)
	if err != nil {
		log.Println(err)
	}
}

//Save results as "version.json"
func SaveData(buildVersion string, path string, client []byte) error {
	mainFolder := "/dataMap-v" + buildVersion
	path = path + mainFolder
	if _, err := os.Stat(path); os.IsNotExist(err) {
		err = os.Mkdir(path, os.ModePerm)
	}
	file, err := os.Create(path + "/version.json")
	if err != nil {
		return err
	}
	defer file.Close()
	_, err = file.Write(client)
	return err
}
