package translator

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"time"
)

type TranslatorAPI interface {
	GetTranslation(data string) (translatedData string, err error)
}

type Translator struct {
	URL string
}

func NewTranslator(translatorURL string) *Translator {
	return &Translator{
		URL: translatorURL,
	}
}

func (t *Translator) GetTranslation(data string) (translatedData string, err error) {
	var translateReq struct {
		Text string `json:"text"`
	}
	translateReq.Text = data
	payload, err := json.Marshal(translateReq)
	if err != nil {
		return translatedData, fmt.Errorf("Error marshalling translator payload %w", err)
	}

	var translateRes struct {
		Contents struct {
			Translated  string `json:"translated"`
			Text        string `json:"text"`
			Translation string `json:"translation"`
		} `json:"contents"`
	}

	err = requestJSON(http.MethodPost, t.URL, "", bytes.NewReader(payload), &translateRes)
	return translateRes.Contents.Translated, err
}

const (
	contentTypeJSON           = "application/json"
	thidpartyTimeoutInSeconds = 5
)

func requestJSON(method, url, path string, body io.Reader, obj interface{}) error {
	req, err := http.NewRequest(method, url+path, body)
	if err != nil {
		return fmt.Errorf("Error constructing HTTP request %w", err)
	}

	q := req.URL.Query()
	req.URL.RawQuery = q.Encode()
	req.Header.Add("content-type", contentTypeJSON)

	client := http.Client{
		Timeout: thidpartyTimeoutInSeconds * time.Second,
	}

	resp, err := client.Do(req)
	if err != nil {
		return fmt.Errorf("Error getting response %w", err)
	}

	defer resp.Body.Close()

	b, _ := ioutil.ReadAll(resp.Body)
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("Unexpected response  %d: %s", resp.StatusCode, string(b))
	}

	if err := json.Unmarshal(b, obj); err != nil {
		return fmt.Errorf("Error decoding response %w, %s", err, string(b))
	}
	return nil
}
