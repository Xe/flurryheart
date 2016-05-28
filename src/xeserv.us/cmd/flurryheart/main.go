package main

import (
	"github.com/namsral/flag"
	"github.com/syfaro/finch"
	_ "github.com/syfaro/finch/commands/help"
	_ "github.com/syfaro/finch/commands/info"
	_ "xeserv.us/commands/countdown"
	_ "xeserv.us/commands/printerfacts"
)

var (
	apiToken = flag.String("api-token", "", "telegram API token")
)

func init() {
	flag.String("config", "", "configuration file")
}

func main() {
	flag.Parse()

	f := finch.NewFinch(*apiToken)

	f.Start()
}
