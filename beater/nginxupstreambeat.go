package beater

import (
	"fmt"
	"time"
	"net/url"

	"github.com/elastic/beats/libbeat/beat"
	"github.com/elastic/beats/libbeat/common"
	"github.com/elastic/beats/libbeat/logp"
	"github.com/elastic/beats/libbeat/publisher"

	"github.com/2Fast2BCn/nginxupstreambeat/config"
	"github.com/2Fast2BCn/nginxupstreambeat/collector"
)

type Nginxupstreambeat struct {
	done       chan struct{}
	config     config.Config
	client     publisher.Client
}

// Creates beater
func New(b *beat.Beat, cfg *common.Config) (beat.Beater, error) {
	config := config.DefaultConfig
	if err := cfg.Unpack(&config); err != nil {
		return nil, fmt.Errorf("Error reading config file: %v", err)
	}

	bt := &Nginxupstreambeat{
		done: make(chan struct{}),
		config: config,
	}
	return bt, nil
}

func (bt *Nginxupstreambeat) Run(b *beat.Beat) error {
	logp.Info("nginxupstreambeat is running! Hit CTRL-C to stop it.")

	bt.client = b.Publisher.Connect()
	ticker := time.NewTicker(bt.config.Period)
	counter := 1
	for {
		select {
		case <-bt.done:
			return nil
		case <-ticker.C:
		}

		var c collector.Collector
		c = collector.NewUpstreamCollector()
		u, err := url.Parse(bt.config.Url)
		if err != nil {
				logp.Err("Fail to parse Nginx upstream status url: %v", err)
		}
		s, err := c.Collect(*u)
		if err != nil {
				logp.Err("Fail to read Nginx upstream status: %v", err)
		}
		//logp.Println(s)

		event := common.MapStr{
				"@timestamp": common.Time(time.Now()),
				"type":       b.Name,
				"nginx_upstream_status":     s,
		}
		bt.client.PublishEvent(event)
		logp.Info("Event sent")
	}
}

func (bt *Nginxupstreambeat) Stop() {
	bt.client.Close()
	close(bt.done)
}
