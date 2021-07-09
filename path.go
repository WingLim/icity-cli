package main

import "os/user"

var (
	home        = homePath()
	iCityPath   = home + "/.icity"
	cookiesPath = iCityPath + "/cookies.json"
)

func homePath() string {
	u, err := user.Current()
	if err != nil {
		return ""
	}
	return u.HomeDir
}
