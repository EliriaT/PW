package guardianNews

type response struct {
	Response embeddedResponse `json:"response"`
}

type embeddedResponse struct {
	Status  string   `json:"status"`
	Results []result `json:"results"`
}
type result struct {
	ID       string `json:"id"`
	WebTitle string `json:"webTitle"`
	WebUrl   string `json:"webUrl"`
}
