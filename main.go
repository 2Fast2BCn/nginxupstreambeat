package main

import (
	"os"

	"github.com/elastic/beats/libbeat/beat"
	"github.com/2Fast2BCn/nginxupstreambeat/beater"
)

var Name = "nginxupstreambeat"

func main() {
	if err := beat.Run(Name, "", beater.New()); err != nil {
		os.Exit(1)
	}
}
