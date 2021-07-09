package main

import "os/user"

var (
	home      = homePath()
	iCityPath = home + "/.icity"
)

func homePath() string {
	u, err := user.Current()
	if err != nil {
		return ""
	}
	return u.HomeDir
}
