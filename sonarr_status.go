package godarr

import "time"

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
	Authentication    bool      `json:"authentication"`
	StartOfWeek       int       `json:"startOfWeek"`
	URLBase           string    `json:"urlBase"`
}
