package xmlhandler

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/beevik/etree"
)

const (
	BingURL = "https://cn.bing.com"
	BingAPI = "https://cn.bing.com/HPImageArchive.aspx?format=xml&idx=%d&n=1&mkt=%s"
)

var Resolution map[string]string
var Markets map[string]bool

func init() {
	Resolution = map[string]string{
		"1366": "1366x768.jpg",  // HD
		"1920": "1920x1080.jpg", // Full HD
		"3840": "UHD.jpg",       // 3840×2160
	}
	Markets = map[string]bool{
		"en-US": true,
		"zh-CN": true,
		"ja-JP": true,
		"en-AU": true,
		"en-UK": true,
		"de-DE": true,
		"en-NZ": true,
		"en-CA": true,
	}
}

// Get: parse .XML from Bing API
func Get(index uint, market string, resolution string) (*Response, error) {
	if _, ok := Resolution[resolution]; !ok {
		return nil, fmt.Errorf("resolution %s not supported", resolution)
	}
	if _, ok := Markets[market]; !ok {
		return nil, fmt.Errorf("market %s not supported", market)
	}

	// new http.Client
	var client = &http.Client{
		Timeout: 5 * time.Second,
	}

	// new http.Request
	var request, err1 = http.NewRequest(http.MethodGet, fmt.Sprintf(BingAPI, index, market), nil)
	if err1 != nil {
		return nil, fmt.Errorf("http.NewRequest error: %s", err1)
	}
	request.Header.Add("Referer", "https://cn.bing.com")
	request.Header.Add(
		"User-Agent",
		"Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/92.0.4515.159 Safari/537.36 Edg/92.0.902.84",
	)

	// request
	var resp, err2 = client.Do(request)
	if err2 != nil {
		return nil, fmt.Errorf("client.Do error: %s", err2)
	}
	defer resp.Body.Close()

	// parse body
	// nice etree!!! (github.com/beevik/etree)
	var body, err3 = ioutil.ReadAll(resp.Body)
	if err3 != nil {
		return nil, fmt.Errorf("ioutil.ReadAll error: %s", err3)
	}
	var doc = etree.NewDocument()
	if err4 := doc.ReadFromBytes(body); err4 != nil {
		return nil, fmt.Errorf("ReadFromBytes error: %s", err4)
	}

	// return Response
	img := doc.SelectElement("images").SelectElement("image")
	return &Response{
		SDate:         img.SelectElement("startdate").Text(),
		EDate:         img.SelectElement("enddate").Text(),
		URL:           fmt.Sprintf("%s%s_%s", BingURL, img.SelectElement("urlBase").Text(), Resolution[resolution]),
		Copyright:     img.SelectElement("copyright").Text(),
		CopyrightLink: img.SelectElement("copyrightlink").Text(),
	}, nil
}
