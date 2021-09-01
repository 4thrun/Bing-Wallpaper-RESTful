package xmlhandler

// definition of struct Response
type Response struct {
	SDate         string `json:"sdate"`
	EDate         string `json:"edate"`
	URL           string `json:"url"`
	Copyright     string `json:"copyright"`
	CopyrightLink string `json:"copyright_link"`
}
