package main

import (
	"time"

	"github.com/juju/errors"
)

type Status struct {
	Version           string    `json:"version"`
	BuildTime         time.Time `json:"buildTime"`
	IsDebug           bool      `json:"isDebug"`
	IsProduction      bool      `json:"isProduction"`
	IsAdmin           bool      `json:"isAdmin"`
	IsUserInteractive bool      `json:"isUserInteractive"`
	StartupPath       string    `json:"startupPath"`
	AppData           string    `json:"appData"`
	OsVersion         string    `json:"osVersion"`
	IsMono            bool      `json:"isMono"`
	IsLinux           bool      `json:"isLinux"`
	IsWindows         bool      `json:"isWindows"`
	Branch            string    `json:"branch"`
	StartOfWeek       int       `json:"startOfWeek"`
	URLBase           string    `json:"urlBase"`
}

func (sc *SonarrClient) SystemStatus() (*Status, error) {
	rv := &Status{}
	err := sc.api.doRequest("GET", "system/status", map[string]string{}, nil, rv)

	if err != nil {
		return nil, errors.Annotate(err, "Failed to get system status.")
	}

	return rv, nil
}
