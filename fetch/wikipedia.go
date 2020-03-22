package fetch

import (
	"encoding/json"
	"fmt"
)

// Wiki struct holds information retrieved from wikipedia api.
type Wiki struct {
	Type         string `json:"type"`
	Title        string `json:"title"`
	DisplayTitle string `json:"displaytitle"`
	PageID       int    `json:"pageid"`
	Description  string `json:"description"`
	Language     string `json:"lang"`
	Extract      string `json:"extract"`
}

// Wikipedia gets the summary from wikipedia api & returns it with an error (if
// exists). It takes a string as input which will be the search query
// sent to wikipedia.
func Wikipedia(q string) (Wiki, error) {
	w := Wiki{}
	api := fmt.Sprintf("https://en.wikipedia.org/api/rest_v1/page/summary/%s", q)

	body, err := getRes(api)
	if err != nil {
		return w, err
	}

	// Unmarshal json to w
	err = json.Unmarshal([]byte(body), &w)
	if err != nil {
		err = fmt.Errorf("Unmarshalling Json failed\n%s", err.Error())
		return w, err
	}

	return w, err
}
