package appstore

import (
	"fmt"
	"github.com/gocolly/colly"
	"net"
	"net/url"
	"strconv"
	"net/http"
	"time"
)

var (
	APPSTORE_URL = "http://ax.phobos.apple.com.edgesuite.net/WebObjects/MZStore.woa/wa/viewContentsUserReviews"
)

func DefaultAPIParams() url.Values {
	params := url.Values{}
	params.Add("sortOrdering", strconv.Itoa(4))
	params.Add("onlyLatestVersion", "false")
	params.Add("type", "Purple Software")
	return params
}

func userAgent(name string) string {
	userAgent := "default user agent"
	if name == "chrome" {
		userAgent = "chrome"
	} else if name == "itunes" {
		userAgent = "iTunes/9.2 (Macintosh; U; Mac OS X 10.6)"
	}
	return userAgent
}

func countryIDHeader(countryID int) string {
	countryIDHead := fmt.Sprintf("%d,5", countryID)
	return countryIDHead
}

func GetCrawler() *colly.Collector {
	c := colly.NewCollector()
	// uses defualt golang httpclient
	c.WithTransport(&http.Transport{
		Proxy: http.ProxyFromEnvironment,
		DialContext: (&net.Dialer{
			Timeout:   30 * time.Second,
			KeepAlive: 30 * time.Second,
			DualStack: true,
		}).DialContext,
		MaxIdleConns:          100,
		IdleConnTimeout:       90 * time.Second,
		TLSHandshakeTimeout:   10 * time.Second,
		ExpectContinueTimeout: 1 * time.Second,
	})

	// one way to set collector configs
	c.AllowURLRevisit = true
	return c
}

func GetXML(url string, countryID int) []byte {
	var xml []byte
	c := GetCrawler()
	c.OnRequest(func(r *colly.Request) {
		r.Headers.Set("User-Agent", userAgent("itunes"))
		r.Headers.Set("X-Apple-Store-Front", countryIDHeader(countryID))
	})
	c.OnResponse(func(r *colly.Response) {
		xml = r.Body
	})
	c.Visit(url)
	c.Wait()
	return xml
}

func GetJSON(url string) []byte {
	var json []byte
	c := GetCrawler()
	c.OnRequest(func(r *colly.Request) {
		r.Headers.Set("User-Agent", userAgent("itunes"))
	})
	c.OnResponse(func(r *colly.Response) {
		json = r.Body
	})
	c.Visit(url)
	c.Wait()
	return json
}
