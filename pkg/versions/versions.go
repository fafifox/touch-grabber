package versions

import (
	"encoding/json"
	"log"
	"regexp"
	"time"
	"touch-grabber/pkg/httputils"
)

const (
	TOUCH_PROXY_URL = "https://proxyconnection.touch.dofus.com"
	LANGAGE = "fr"
)

var (
	TOUCH_CONFIG_URL   = TOUCH_PROXY_URL + "/config.json"
	TOUCH_BUILD_SCRIPT = TOUCH_PROXY_URL + "/build/script.js"
	TOUCH_ASSETS       = TOUCH_PROXY_URL + "/assetsVersions.json"
	APP_URL            = "https://itunes.apple.com/lookup?country="+ LANGAGE +"&id=1041406978&lang="+ LANGAGE +"&limit=1&t=" + string(time.Now().UTC().UnixNano() / 1e6)
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

func GetVersions() {

}
