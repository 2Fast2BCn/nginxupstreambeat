package collector

import (
	"fmt"
	"net/http"
	"net/url"
	"io/ioutil"
	"encoding/json"

	"github.com/elastic/beats/libbeat/logp"
)

// UpstreamCollector is a Collector that collects Nginx upstream status page.
type UpstreamCollector struct {
	http     *http.Client
	requests int
}

type Response struct {
	Servers	Servers    `json:"servers"`
}

type Servers struct {
	Total   int         `json:"total"`
	Generation []string `json:"generation"`
	Server	[]Server    `json:"server"`
}

type Server struct {
	Index   int         `json:"index"`
	Upstream string     `json:"upstream"`
	Name	 string     `json:"name"`
	Status	 string     `json:"status"`
	Rise	 int        `json:"rise"`
	Fall	 int        `json:"fall"`
	Type	 string     `json:"type"`
	Port	 int        `json:"port"`
}

// NewStubCollector constructs a new UpstreamCollector.
func NewUpstreamCollector() Collector {
	return &UpstreamCollector{
		http:     HTTPClient(),
		requests: 0,
	}
}

// Collect Nginx upstream status from given url.
func (c *UpstreamCollector) Collect(u url.URL) (map[string]interface{}, error) {
	res, err := c.http.Get(u.String())
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	if res.StatusCode != 200 {
		return nil, fmt.Errorf("HTTP%s", res.Status)
	}

	// Nginx upstream status sample:
	// {"servers": {
	//	"total": 4,
	//	"generation": 1,
	//	"server": [
	//		{"index": 0, "upstream": "websphere", "name": "10.60.42.8:9080", "status": "up", "rise": 366, "fall": 0, "type": "http", "port": 0},
	//		{"index": 1, "upstream": "websphere", "name": "10.60.42.8:9081", "status": "up", "rise": 2560, "fall": 0, "type": "http", "port": 0},
	//		{"index": 2, "upstream": "websphere", "name": "10.60.42.9:9080", "status": "up", "rise": 1861, "fall": 0, "type": "http", "port": 0},
	//		{"index": 3, "upstream": "websphere", "name": "10.60.42.9:9081", "status": "up", "rise": 764, "fall": 0, "type": "http", "port": 0}
	//	]
	// }}
	//
	var dat map[string]interface{}
	body, err := ioutil.ReadAll(res.Body)
	if err := json.Unmarshal([]byte(body), &dat); err != nil {
		panic(err)
	}
	
	if err != nil {
		logp.Err("Error closing: %v", err)
	}
	fmt.Println(dat)
	return dat, nil
}