package client

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
	RawAddress string
	ApiKey     string
	Address    *url.URL
	Client     *http.Client
}

func (api *ApiClient) ConnSetup() error {
	if api.RawAddress == "" {
		return errors.New("No RawAddress specified")
	} else if api.ApiKey == "" {
		return errors.New("No API Key specified")
	}

	RawAddressUrl, err := url.Parse(api.RawAddress)
	if err != nil {
		return err
	}

	if !strings.HasSuffix(RawAddressUrl.Path, "/") {
		RawAddressUrl.Path += "/"
	}
	if !strings.HasSuffix(RawAddressUrl.Path, "api/") {
		RawAddressUrl.Path += "api/"
	}
	api.RawAddress = RawAddressUrl.String()
	api.Address = RawAddressUrl

	api.Client = http.DefaultClient
	return nil
}

func (api *ApiClient) DoRequest(action, path string, params map[string]string, reqData, resData interface{}) error {
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
	req.Header.Add("X-Api-Key", api.ApiKey)
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
