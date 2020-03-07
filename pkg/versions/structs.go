package versions

import "time"

type Client struct {
	AppVersion    string `json:"appVersion"`
	BuildVersion  string `json:"buildVersion"`
	AssetsVersion string `json:"assetsVersion"`
	AssetsURL     string `json:"assetsUrl"`
	DataURL       string `json:"dataUrl"`
}

type Assets struct {
	AssetsVersion     string   `json:"assetsVersion"`
	StaticDataVersion string   `json:"staticDataVersion"`
	ChangedFiles      []string `json:"changedFiles"`
}

type Application struct {
	ResultCount int `json:"resultCount"`
	Results     []struct {
		IsGameCenterEnabled                bool          `json:"isGameCenterEnabled"`
		ScreenshotUrls                     []string      `json:"screenshotUrls"`
		IpadScreenshotUrls                 []string      `json:"ipadScreenshotUrls"`
		AppletvScreenshotUrls              []interface{} `json:"appletvScreenshotUrls"`
		ArtworkURL60                       string        `json:"artworkUrl60"`
		ArtworkURL512                      string        `json:"artworkUrl512"`
		ArtworkURL100                      string        `json:"artworkUrl100"`
		ArtistViewURL                      string        `json:"artistViewUrl"`
		SupportedDevices                   []string      `json:"supportedDevices"`
		Advisories                         []string      `json:"advisories"`
		Kind                               string        `json:"kind"`
		Features                           []string      `json:"features"`
		TrackCensoredName                  string        `json:"trackCensoredName"`
		LanguageCodesISO2A                 []string      `json:"languageCodesISO2A"`
		FileSizeBytes                      string        `json:"fileSizeBytes"`
		SellerURL                          string        `json:"sellerUrl"`
		ContentAdvisoryRating              string        `json:"contentAdvisoryRating"`
		AverageUserRatingForCurrentVersion float64       `json:"averageUserRatingForCurrentVersion"`
		UserRatingCountForCurrentVersion   int           `json:"userRatingCountForCurrentVersion"`
		TrackViewURL                       string        `json:"trackViewUrl"`
		TrackContentRating                 string        `json:"trackContentRating"`
		TrackID                            int           `json:"trackId"`
		TrackName                          string        `json:"trackName"`
		GenreIds                           []string      `json:"genreIds"`
		FormattedPrice                     string        `json:"formattedPrice"`
		PrimaryGenreName                   string        `json:"primaryGenreName"`
		IsVppDeviceBasedLicensingEnabled   bool          `json:"isVppDeviceBasedLicensingEnabled"`
		ReleaseDate                        time.Time     `json:"releaseDate"`
		MinimumOsVersion                   string        `json:"minimumOsVersion"`
		SellerName                         string        `json:"sellerName"`
		CurrentVersionReleaseDate          time.Time     `json:"currentVersionReleaseDate"`
		ReleaseNotes                       string        `json:"releaseNotes"`
		PrimaryGenreID                     int           `json:"primaryGenreId"`
		Currency                           string        `json:"currency"`
		Version                            string        `json:"version"`
		WrapperType                        string        `json:"wrapperType"`
		ArtistID                           int           `json:"artistId"`
		ArtistName                         string        `json:"artistName"`
		Genres                             []string      `json:"genres"`
		Price                              float64       `json:"price"`
		Description                        string        `json:"description"`
		BundleID                           string        `json:"bundleId"`
		AverageUserRating                  float64       `json:"averageUserRating"`
		UserRatingCount                    int           `json:"userRatingCount"`
	} `json:"results"`
}

type Config struct {
	AssetsURL string `json:"assetsUrl"`
	Analytics struct {
		AnkAnalytics struct {
			Version   string `json:"version"`
			Debug     bool   `json:"debug"`
			Analog    bool   `json:"analog"`
			UseURLMap bool   `json:"useUrlMap"`
		} `json:"ankAnalytics"`
		WizAnalytics struct {
			Localytics struct {
				Token   string `json:"token"`
				Options struct {
					SessionTimeoutSeconds int    `json:"sessionTimeoutSeconds"`
					Namespace             string `json:"namespace"`
				} `json:"options"`
			} `json:"localytics"`
		} `json:"wizAnalytics"`
	} `json:"analytics"`
	DisabledFeatures struct {
		ShopByServerIDList []string `json:"shopByServerIdList"`
	} `json:"disabledFeatures"`
	Adjust struct {
		AppToken    string `json:"appToken"`
		Environment string `json:"environment"`
	} `json:"adjust"`
	Logging struct {
		Groups    bool `json:"groups"`
		LogLevels struct {
			Time      int `json:"time"`
			Verbose   int `json:"verbose"`
			Debug     int `json:"debug"`
			Info      int `json:"info"`
			Notice    int `json:"notice"`
			Warning   int `json:"warning"`
			Error     int `json:"error"`
			Critical  int `json:"critical"`
			Alert     int `json:"alert"`
			Emergency int `json:"emergency"`
		} `json:"logLevels"`
		Config struct {
			Console []string `json:"console"`
			Server  []string `json:"server"`
		} `json:"config"`
		DisableOverride bool `json:"disableOverride"`
	} `json:"logging"`
	DataURL string `json:"dataUrl"`
	Haapi   struct {
		ID       int    `json:"id"`
		URL      string `json:"url"`
		Hostname string `json:"hostname"`
	} `json:"haapi"`
	ServerLanguages []string `json:"serverLanguages"`
	Notification    struct {
		Push struct {
			AppID     string `json:"appId"`
			ProjectID string `json:"projectId"`
		} `json:"push"`
	} `json:"notification"`
	Recaptcha struct {
		ProxyURL  string `json:"proxyUrl"`
		AnkamaURL string `json:"ankamaUrl"`
	} `json:"recaptcha"`
	FailoverLanguage string `json:"failoverLanguage"`
	UIURL            string `json:"uiUrl"`
	Language         string `json:"language"`
	SessionID        string `json:"sessionId"`
}
