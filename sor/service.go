package sor

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/davidwrpayne/wellmetersystem/schema"
	"net/http"
)

type Service interface {
	Publish(measurement *schema.Measurement) error
}

const postMeasurementPath = "/watertank/measurment/new"
const contentType = "application/json"

type HTTPSystemOfRecord struct {
	token      string
	host       string
	httpClient *http.Client
}

var _ Service = (*HTTPSystemOfRecord)(nil)

func NewHttpSystemOfRecord(hostUrl, bearerToken string) *HTTPSystemOfRecord {
	return &HTTPSystemOfRecord{
		token:      bearerToken,
		host:       hostUrl,
		httpClient: http.DefaultClient,
	}
}

func (h HTTPSystemOfRecord) Publish(measurement *schema.Measurement) error {
	postUrl := fmt.Sprintf("%s%s", h.host, postMeasurementPath)
	body, err := json.Marshal(measurement)
	if err != nil {
		return err
	}

	resp, err := h.httpClient.Post(postUrl, contentType, bytes.NewReader(body))
	if err != nil {
		return err
	}
	if resp.StatusCode != 200 && resp.StatusCode != 204 { // OK or OK Created
		return errors.New("Failed to post successful JSON body")
	}

	return nil // successful post
}

type LoggingSOR struct {
	token string
	url   string
}

func (l LoggingSOR) Publish(measurement *schema.Measurement) error {
	fmt.Printf("Publish called\n")
	fmt.Printf("measurement %v\n", measurement)
	fmt.Printf("would have POSTED to %s\n", l.url)
	fmt.Printf("using bearer token POSTED to %s\n", l.token)
	return nil
}

var _ Service = (*LoggingSOR)(nil)

func NewLoggingSystemOfRecord(url, bearerToken string) *LoggingSOR {
	return &LoggingSOR{
		url:   url,
		token: bearerToken,
	}
}
