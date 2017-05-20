package main

import (
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/dokku/dokku/plugins/config"
	columnize "github.com/ryanuber/columnize"
)

const (
	helpHeader = `Usage: dokku config [<app>|--global]

Display all global or app-specific config vars

Additional commands:`

	helpContent = `
	config (<app>|--global), Pretty-print an app or global environment
	config:get (<app>|--global) KEY, Display a global or app-specific config value
	config:set (<app>|--global) [--encoded] [--no-restart] KEY1=VALUE1 [KEY2=VALUE2 ...], Set one or more config vars
	config:unset (<app>|--global) KEY1 [KEY2 ...], Unset one or more config vars
	config:export (<app>|--global) [--envfile], Export a global or app environment
	config:keys (<app> | --global) [--merged], Show keys set in environment
	config:bundle (<app> | --global) [--merged], Bundle environment into tarfile
`
)

func main() {
	flag.Usage = usage
	flag.Parse()

	cmd := flag.Arg(0)
	switch cmd {
	case "config", "config:show":
		target := flag.Arg(1)
		env := config.GetConfig(target, false)
		fmt.Print(config.PrettyPrintEnvEntries("", env.Map()))
	case "config:help":
		usage()
	case "help":
		fmt.Print(helpContent)
	default:
		dokkuNotImplementExitCode, err := strconv.Atoi(os.Getenv("DOKKU_NOT_IMPLEMENTED_EXIT"))
		if err != nil {
			fmt.Println("failed to retrieve DOKKU_NOT_IMPLEMENTED_EXIT environment variable")
			dokkuNotImplementExitCode = 10
		}
		os.Exit(dokkuNotImplementExitCode)
	}
}

func usage() {
	config := columnize.DefaultConfig()
	config.Delim = ","
	config.Prefix = "\t"
	config.Empty = ""
	content := strings.Split(helpContent, "\n")[1:]
	fmt.Println(helpHeader)
	fmt.Println(columnize.Format(content, config))
}
