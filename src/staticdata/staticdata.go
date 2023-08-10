package staticdata

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func getJSON(url string, dest interface{}) error {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return err
	}

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}

	defer res.Body.Close()

	if res.StatusCode != 200 {
		return fmt.Errorf("staticdata: unexpected status code %d", res.StatusCode)
	}

	return json.NewDecoder(res.Body).Decode(dest)
}
