package main

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

type ApiClient struct {
	address string
	apiKey  string
	Address *url.URL
	Client  *http.Client
}

func (api *ApiClient) connSetup() error {
	if api.address == "" {
		return errors.New("No address specified")
	} else if api.apiKey == "" {
		return errors.New("No API Key specified")
	}

	addressUrl, err := url.Parse(api.address)
	if err != nil {
		return err
	}

	if !strings.HasSuffix(addressUrl.Path, "/") {
		addressUrl.Path += "/"
	}
	if !strings.HasSuffix(addressUrl.Path, "api/") {
		addressUrl.Path += "api/"
	}
	api.address = addressUrl.String()
	api.Address = addressUrl

	api.Client = http.DefaultClient
	return nil
}

func (api *ApiClient) doRequest(action, path string, params map[string]string, reqData, resData interface{}) error {
	lookupUrl := *api.Address
	parameters := url.Values{}

	if params != nil {
		for k, v := range params {
			parameters.Add(k, v)
		}
	}

	lookupUrl.RawQuery = parameters.Encode()
	lookupUrl.Path += path
	jsonValue, err := json.Marshal(reqData)

	if err != nil {
		return err
	}

	req, err := http.NewRequest(action, lookupUrl.String(), bytes.NewBuffer(jsonValue))

	if err != nil {
		return err
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Add("X-Api-Key", api.apiKey)
	response, err := api.Client.Do(req)

	if err != nil {
		return err
	}

	if response.StatusCode < 200 || response.StatusCode >= 300 {
		bodyBytes, err := ioutil.ReadAll(response.Body)

		if err == nil {
			fmt.Printf("Failing (%v) call returned:\n%v", response.StatusCode, string(bodyBytes))
		}

		return errors.New(fmt.Sprintf("Status code %v", response.StatusCode))
	}

	body, err := ioutil.ReadAll(response.Body)
	err = json.NewDecoder(bytes.NewBuffer(body)).Decode(resData)

	if err != nil {
		return err
	}

	return nil
}
