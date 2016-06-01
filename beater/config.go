package beater

import (
	"errors"
	"time"
)

type ConfigSettings struct {
	Nginxupstreambeat NginxupstreamConfig `config:"nginxupstreambeat"`
}

type NginxupstreamConfig struct {
	Period time.Duration     `config:"period"  validate:"min=1ms"`
	Url  string          `config:"url"`
}

var (
	defaultConfig = NginxupstreamConfig{
		Period: 10 * time.Second,
		Url:  string "http://127.0.0.1/nginx_upstream_status",
	}
)

func (c *NginxupstreamConfig) Validate() error {
	return nil
}
