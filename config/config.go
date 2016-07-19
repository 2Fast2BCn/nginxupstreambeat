// Config is put into a different package to prevent cyclic imports in case
// it is needed in several locations

package config

import (
	"time"
	"net/url"
)

type Config struct {
	Period  time.Duration `config:"period"`
	Url     *url.URL      `config:"url"`
}

var DefaultConfig = Config{
	Period: 1 * time.Second,
	Url: "http://127.0.0.1/nginx_upstream_status"
}
