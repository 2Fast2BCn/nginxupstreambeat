package collector

import (
	"fmt"
	"encoding/json"
)

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

func main() {
    fmt.Printf("hello, world\n")
    byt := []byte(`{"servers": {"total": 4,"generation": 1,"server": [{"index": 0, "upstream": "websphere", "name": "10.60.42.8:9080", "status": "up", "rise": 366, "fall": 0, "type": "http", "port": 0}, {"index": 1, "upstream": "websphere", "name": "10.60.42.8:9081", "status": "up", "rise": 2560, "fall": 0, "type": "http", "port": 0}, {"index": 2, "upstream": "websphere", "name": "10.60.42.9:9080", "status": "up", "rise": 1861, "fall": 0, "type": "http", "port": 0}, {"index": 3, "upstream": "websphere", "name": "10.60.42.9:9081", "status": "up", "rise": 764, "fall": 0, "type": "http", "port": 0}]}}`)
	var dat map[string]interface{}
	if err := json.Unmarshal(byt, &dat); err != nil {
		panic(err)
	}
    fmt.Println(dat)
}