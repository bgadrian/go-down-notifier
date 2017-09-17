package godown

import (
	"errors"
	"net/http"
	"strconv"
	"time"
)

//HTTPRequest Returns error if the URL request fails.
func HTTPRequest(timeoutSeconds int, userAgent, method, url string) error {
	tr := &http.Transport{
		MaxIdleConns:    10,
		IdleConnTimeout: time.Duration(timeoutSeconds) * time.Second,
	}
	client := &http.Client{
		Transport: tr,
		Timeout:   time.Duration(timeoutSeconds) * time.Second,
	}

	req, err := http.NewRequest(method, url, nil)
	if err != nil {
		return err
	}

	req.Header.Set("User-Agent", userAgent)

	resp, err := client.Do(req)
	if err != nil {
		return err
	}

	defer resp.Body.Close()
	if resp.StatusCode < 200 || resp.StatusCode >= 400 {
		return errors.New("status code " + strconv.Itoa(resp.StatusCode))
	}

	return nil
}

//PortRequest ...
func portRequest() {
	//TODO you can use https://github.com/janosgyerik/portping/blob/master/portping.go
}
