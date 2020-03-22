package fetch

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

func getRes(api string) ([]byte, error) {
	var body []byte

	c := http.Client{
		// TODO: timeout should be configurable by the user
		Timeout: time.Second * 64,
	}

	req, err := http.NewRequest(http.MethodGet, api, nil)
	if err != nil {
		return body, err
	}

	req.Header.Set("User-Agent", "Andinus / Indus - framagit.org/andinus/indus")

	res, err := c.Do(req)
	if err != nil {
		return body, err
	}
	defer res.Body.Close()

	if res.StatusCode != 200 {
		return body, fmt.Errorf("Unexpected response status code received: %d %s",
			res.StatusCode, http.StatusText(res.StatusCode))
	}

	body, err = ioutil.ReadAll(res.Body)
	if err != nil {
		return body, err
	}
	return body, err
}
