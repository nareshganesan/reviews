package appstore

import (
	"encoding/json"
	"net/url"
	"strconv"
	"fmt"
	"time"
)

// Ref: https://affiliate.itunes.apple.com/resources/documentation/linking-to-the-itunes-music-store/

type App struct {
	ScreenshotURLList           []string // iphone screen shots
	IpadScreenshotURLList       []string // ipad screenshots
	AppletvScreenshotURLList    []string // apple tc screen shots
	ArtworkURL60                string   // logo
	ArtworkURL512               string   // logo
	ArtworkURL100               string   // logo
	ArtistViewURL               string   // itunes app url
	Kind                        string   // entity type software
	Features                    []string // app features
	IsGameCenterEnabled         bool
	SupportedDevices            []string // supported apple devices
	Advisories                  []string // reasons for usage
	AvgUserRatingCurVersion     int      // averageUserRatingForCurrentVersion
	TrackCensoredName           string
	LanguageCodesISO2A          []string // languageCodesISO2A - language codes
	FileSizeBytes               int
	SellerURL                   string
	ContentAdvisoryRating       string
	UserRatingCountCurVersion   int    // userRatingCountCurrentVersion
	TrackViewURL                string // itunes app page
	TrackContentRating          string
	SellerName                  string // company name
	IsVppDeviceLicensingEnabled bool
	GenreIDList                 []string // genreIds
	TrackID                     int      // trackId
	TrackName                   string
	FormattedPrice              string
	Currency                    string
	WrapperType                 string // entity wrapper type software
	Version                     string
	Description                 string
	ArtistID                    int      // company id
	ArtistName                  string   // company name
	Genres                      []string // genre names list
	Price                       float64
	BundleID                    string // software package name
	ReleaseNotes                string
	ReleaseDate                 time.Time
	CurVersionReleaseDate       time.Time
	MinimumOSVersion            string
	PrimaryGenreName            string
	PrimaryGenreID              int
	AverageUserRating           int
	UserRatingCount             int
	CrawledDate                 time.Time
	CurCountryID                int
	Reviews                     []Review
}

var (
	SEARCH_URL = "https://itunes.apple.com/search"
	LOOKUP_URL = "https://itunes.apple.com/lookup"
)

// Search by name
func (a *App) SearchByName() []App {
	var Url *url.URL
	Url, _ = url.Parse(SEARCH_URL)
	params := url.Values{}
	params.Add("term", a.TrackName)
	params.Add("entity", "software")
	Url.RawQuery = params.Encode()
	jsonBlob := GetJSON(Url.String())
	results := make(map[string]interface{})
	err := json.Unmarshal(jsonBlob, &results)
	if err != nil {
		fmt.Println("error:", err)
	}
	var apps []App
	for _, inter := range results["results"].([]interface{}) {
		var app App
		bytes, _ := json.Marshal(inter)
		_ = json.Unmarshal(bytes, &app)
		apps = append(apps, app)
	}
	return apps
}

func (a *App) GetAppDetails() {
	var Url *url.URL
	Url, _ = url.Parse(LOOKUP_URL)
	params := url.Values{}
	params.Add("id", strconv.Itoa(a.TrackID))
	Url.RawQuery = params.Encode()	
	jsonBlob := GetJSON(Url.String())
	results := make(map[string]interface{})
	err := json.Unmarshal(jsonBlob, &results)
	if err != nil {
		fmt.Println("error:", err)
	}
	// var app App
	bytes, _ := json.Marshal(results["results"].([]interface{})[0])
	_ = json.Unmarshal(bytes, a)
	// fmt.Println(app.TrackID, app.TrackName)
	// return app
}

// JSON to ES index dump

// ES index backup as json - snapshot api
