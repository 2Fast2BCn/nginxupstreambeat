package main

import (
	"os"

	"github.com/elastic/beats/libbeat/beat"

	"github.com/2Fast2BCn/nginxupstreambeat/beater"
)

var Version = "0.0.1"
var Name = "nginxupstreambeat"

func main() {
	err := beat.Run(Name, Version, beater.New)
	if err != nil {
		os.Exit(1)
	}
}
