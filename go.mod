module github.com/WingLim/icity-cli

go 1.16

require (
	github.com/WingLim/icity-sdk v0.9.1
	github.com/chzyer/readline v0.0.0-20180603132655-2972be24d48e
	github.com/manifoldco/promptui v0.8.0
	github.com/urfave/cli/v2 v2.3.0
	golang.org/x/crypto v0.0.0-20190308221718-c2843e01d9a2
)

replace github.com/manifoldco/promptui => ../promptui
