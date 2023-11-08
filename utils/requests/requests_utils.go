package requests

import (
	"encoding/json"
	"github.com/soshika/sample-search/logger"
	"io/ioutil"
	"net/http"
	"strings"
)

type Req struct {
	URL    string `json:"url"`
	Method string `json:"method"`
}

func (r *Req) POST(jsonMap map[string]interface{}) ([]byte, error) {
	jsonBody, _ := json.Marshal(jsonMap)
	payload := strings.NewReader(string(jsonBody))

	client := &http.Client{}
	req, err := http.NewRequest(r.Method, r.URL, payload)
	if err != nil {
		logger.Error("Could not sent request to target url", err)
		return nil, err
	}
	req.Header.Add("Content-Type", "application/json")

	res, err := client.Do(req)
	if err != nil {
		logger.Error("Could not send request to target url", err)
		return nil, err
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		logger.Error("Could not send request to target url", err)
		return nil, err
	}
	return body, nil
}
