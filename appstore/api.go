package appstore

import (
	"fmt"
)

func SearchApps(name string) []App {
	var app App
	fmt.Println("Search apps called")
	app.TrackName = name
	return app.SearchByName()
}

func GetAppDetails(appID int) App {
	var app App
	fmt.Println("Get app details called")
	app.TrackID = appID
	app.GetAppDetails()
	return app
}

func GetReviews(appID, countryID, pageNo int) App {
	var app App
	fmt.Println("Search reviews called")
	if appID <= 0 {
		appID = 474990205
	}
	app.TrackID = appID
	app.CurCountryID = countryID
	app.GetReviews(pageNo)
	return app
}

func GetAllReviews(appID, countryID int) App {
	var app App
	fmt.Println("Get all reviews called")
	if appID <= 0 {
		appID = 368677368 //default to UBER
	}
	app.TrackID = appID
	app.CurCountryID = countryID
	app.GetAllReviews()
	return app
}
