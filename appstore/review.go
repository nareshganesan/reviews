package appstore

import (
	"bytes"
	"encoding/xml"
	"fmt"
	"github.com/antchfx/xmlquery"
	// "github.com/beevik/etree"
	"github.com/metakeule/fmtdate"
	"net/url"
	"strconv"
	"strings"
	"time"
)

type RootLevel struct {
	XMLName xml.Name `xml:"Document"`
	// Raw string `xml:",innerxml"`
	View view
}

type view struct {
	Ele        string `xml:",innerxml"`
	ScrollView scrollview
}

type scrollview struct {
	Ele      string `xml:",innerxml"`
	VBoxView vboxview
}

type vboxview struct {
	Ele  string `xml:",innerxml"`
	View view1
}

type view1 struct {
	Ele        string `xml:",innerxml"`
	MatrixView matrixview
}

type matrixview struct {
	Ele      string `xml:",innerxml"`
	VBoxView vboxview1
}

type vboxview1 struct {
	Ele      string `xml:",innerxml"`
	VBoxView vboxview2
}

type vboxview2 struct {
	Ele      string `xml:",innerxml"`
	VBoxView []vboxview3
	HBoxView []hboxview2
}

type vboxview3 struct {
	Ele      string `xml:",innerxml"`
	HBoxView []hboxview1
	TextView textview
}

type hboxview1 struct {
	Ele      string `xml:",innerxml"`
	TextView []textview
}

type hboxview2 struct {
	Ele      string `xml:",innerxml"`
	TextView textview
}

type textview struct {
	Ele string `xml:",innerxml"`
}

type setfontstyle struct {
	Ele string `xml:",innerxml"`
}

type setfontstyle1 struct {
	Ele string `xml:",chardata"`
	B   b
}

type b struct {
	Ele string `xml:",innerxml"`
}

type Review struct {
	Title      string
	Desciption string
	CreatedAt  time.Time
	User       string
	Rating     int
	ID         string
	AppVersion string
}

func (a *App) GetReviews(pageNo int) {

	var existingReviews []Review

	if a.CurCountryID <= 0 {
		a.CurCountryID = 143441
	}
	if pageNo < 0 {
		pageNo = 0
	}

	if len(existingReviews) > 0 {
		a.Reviews = existingReviews
	} else {
		var Url *url.URL
		Url, _ = url.Parse(APPSTORE_URL)
		params := DefaultAPIParams()
		params.Add("id", strconv.Itoa(a.TrackID))
		params.Add("pageNumber", strconv.Itoa(pageNo))
		Url.RawQuery = params.Encode()
		xml := GetXML(Url.String(), a.CurCountryID)
		a.Reviews = ParseReviews(xml)
	}
}

func (a *App) GetAllReviews() {
	pageNo := 0
	var Url *url.URL
	Url, _ = url.Parse(APPSTORE_URL)
	params := DefaultAPIParams()
	params.Add("id", strconv.Itoa(a.TrackID))
	params.Add("pageNumber", strconv.Itoa(pageNo))
	Url.RawQuery = params.Encode()
	fmt.Println("page: ", 0)
	xmlData := GetXML(Url.String(), a.CurCountryID)
	totalPages := a.GetTotalReviewPages(xmlData)
	var reviews []Review
	reviews = ParseReviews(xmlData)
	for i := 1; i < totalPages; i++ {
		pageparams := DefaultAPIParams()
		pageparams.Add("id", strconv.Itoa(a.TrackID))
		pageparams.Add("pageNumber", strconv.Itoa(i))
		Url.RawQuery = pageparams.Encode()
		xmlData = GetXML(Url.String(), a.CurCountryID)
		currentPageReviews := ParseReviews(xmlData)
		fmt.Println("page: ", i)
		for _, review := range currentPageReviews {
			reviews = append(reviews, review)
		}
	}
	fmt.Println(len(reviews))
	a.Reviews = reviews
}

func (a *App) GetTotalReviewPages(xmlData []byte) int {
	var pages int
	rbody := RootLevel{}
	xml.Unmarshal(xmlData, &rbody)
	// totalPagesNode
	totalPagesWrapperNode := rbody.View.ScrollView.VBoxView.
		View.MatrixView.VBoxView.VBoxView.
		HBoxView[1].TextView.Ele
	totalPagesNode := setfontstyle{}
	totalPages := b{}
	xml.Unmarshal([]byte(totalPagesWrapperNode), &totalPagesNode)
	xml.Unmarshal([]byte(totalPagesNode.Ele), &totalPages)
	// fmt.Println(totalPages.Ele)
	if strings.Contains(totalPages.Ele, "Page") {
		data := strings.Split(totalPages.Ele, " ")
		length := len(data)
		pages, _ = strconv.Atoi(data[length-1])
	} else {
		pages = 1
	}
	fmt.Println(pages)
	return pages
}

func ParseReviews(xmlData []byte) []Review {

	body := bytes.NewReader(xmlData)
	doc, _ := xmlquery.Parse(body)
	reviewsNode := xmlquery.Find(doc, "//Document/View/ScrollView/VBoxView/View/MatrixView/VBoxView/VBoxView/VBoxView")

	var reviews []Review
	rbody := RootLevel{}
	xml.Unmarshal(xmlData, &rbody)

	for idx, reviewNode := range reviewsNode {
		urlNode := xmlquery.FindOne(reviewNode, "//HBoxView/HBoxView/LoadFrameURL")
		var review Review
		// reviewID
		var reviewID, title, user, version string
		var rating int
		var date time.Time
		idNode := xmlquery.FindOne(urlNode, "//reviewId")
		if idNode != nil {
			fmt.Println("id is present")
			fmt.Println(idNode)
		} else {
			for _, el := range urlNode.Attr {
				if el.Name.Local == "url" {
					reviewURL, _ := url.Parse(el.Value)
					reviewID = reviewURL.Query()["userReviewId"][0]
				}
			}
		}
		// title
		titleNode := xmlquery.FindOne(reviewNode, "//HBoxView/TextView/SetFontStyle/b")
		title = titleNode.InnerText()
		// rating
		ratingNode := xmlquery.FindOne(reviewNode, "//HBoxView/HBoxView/HBoxView")
		for _, el := range ratingNode.Attr {
			if el.Name.Local == "alt" {
				rtg := strings.Split(el.Value, " ")[0]
				rating, _ = strconv.Atoi(rtg)
			}
		}
		// review user
		userNode := xmlquery.FindOne(reviewNode, "//HBoxView/TextView/SetFontStyle/GotoURL/b")
		user = strings.TrimSpace(userNode.InnerText())
		// version and date
		versionWrapperNode := rbody.View.ScrollView.VBoxView.
			View.MatrixView.VBoxView.VBoxView.
			VBoxView[idx].HBoxView[1].TextView[0].Ele

		versionNode := setfontstyle{}
		xml.Unmarshal([]byte(versionWrapperNode), &versionNode)
		versionText := strings.Split(versionNode.Ele, "\n")
		var ranText string
		for _, s := range versionText {
			s = strings.TrimSpace(s)
			if s != "" && s != "by" && s != "-" {
				if s != " " && s != "\n" && s != "\t" && s != "\r" {
					ranText += s
					if version == "" && strings.Contains(s, "Version") {
						version = strings.TrimSpace(s)
						version = strings.Split(version, " ")[1]
					}
					if !strings.Contains(s, "Version") && ranText != "" {
						date, _ = fmtdate.Parse("MMM DD, YYYY", strings.TrimSpace(s))
					}
				}
			}
		}
		// review description
		descriptionEle := rbody.View.ScrollView.VBoxView.
			View.MatrixView.VBoxView.VBoxView.
			VBoxView[idx].TextView.Ele
		desc := setfontstyle{}
		xml.Unmarshal([]byte(descriptionEle), &desc)
		description := desc.Ele
		// create review object
		review.ID = reviewID
		review.CreatedAt = date
		review.AppVersion = version
		review.Title = title
		review.User = user
		review.Rating = rating
		review.Desciption = description
		reviews = append(reviews, review)
	}
	return reviews
}
