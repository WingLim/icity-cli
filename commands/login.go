package commands

import (
	"errors"
	"fmt"
	"log"
	"os"

	icity "github.com/WingLim/icity-sdk"
	"github.com/manifoldco/promptui"
	"github.com/urfave/cli/v2"
)

var LoginCommand = &cli.Command{
	Name:  "login",
	Usage: "Login to iCity",
	Action: func(context *cli.Context) error {
		user := doLogin(context)
		if user == nil {
			return errors.New("login failed")
		}
		fmt.Println("login success")
		return nil
	},
	Flags: []cli.Flag{
		&cli.StringFlag{
			Name:    "config",
			Aliases: []string{"c"},
			Usage:   "config.json file path",
		},
	},
}

func doLogin(context *cli.Context) *icity.User {
	configPath := context.String("config")

	var user *icity.User
	if configPath != "" {
		user = icity.LoginWithConfig(configPath, icity.WithSaveCookies(cookiesPath))
	} else {
		usernameInput := promptui.Prompt{
			Label: "Username",
		}
		username, err := usernameInput.Run()
		if err != nil {
			log.Fatal(err)
			return nil
		}

		passwordInput := promptui.Prompt{
			Label: "Password",
			Mask:  '*',
		}
		password, err := passwordInput.Run()
		if err != nil {
			log.Fatal(err)
			return nil
		}

		user = icity.Login(username, password, icity.WithSaveCookies(cookiesPath))
	}
	return user
}

func isLogin() (*icity.User, bool) {
	user := icity.LoginWithCookies(cookiesPath)
	if user != nil {
		return user, true
	}
	fmt.Println("please use icity login to login")
	return nil, false
}

func getUser() *icity.User {
	if user, ok := isLogin(); ok {
		return user
	}
	os.Exit(1)
	return nil
}
