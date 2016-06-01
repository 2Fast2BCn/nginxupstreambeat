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
	beatConfig *config.Config
	done       chan struct{}
	period     time.Duration
	url        *url.URL
	client     publisher.Client
}

// Creates beater
func New() *Nginxupstreambeat {
	return &Nginxupstreambeat{
		done: make(chan struct{}),
	}
}

func (bt *Nginxupstreambeat) Config(b *beat.Beat) error {
	// Load beater beatConfig
	err := b.RawConfig.Unpack(&bt.beatConfig)
	if err != nil {
		return fmt.Errorf("Error reading config file: %v", err)
	}

	logp.Debug("nginxupstreambeat", "Init nginxupstreambeat")
	logp.Debug("nginxupstreambeat", "Period %v", bt.period)
	logp.Debug("nginxupstreambeat", "Url %v", bt.url)

	return nil
}

func (bt *Nginxupstreambeat) Setup(b *beat.Beat) error {
	// Setting default period if not set
	if bt.beatConfig.Nginxupstreambeat.Period == "" {
		bt.beatConfig.Nginxupstreambeat.Period = "1s"
	}

	bt.client = b.Publisher.Connect()

	var err error
	bt.period, err = time.ParseDuration(bt.beatConfig.Nginxupstreambeat.Period)
	if err != nil {
		return err
	}
	//add parse URL ???

	return nil
}

func (bt *Nginxupstreambeat) Run(b *beat.Beat) error {
	logp.Info("nginxupstreambeat is running! Hit CTRL-C to stop it.")

	ticker := time.NewTicker(bt.period)
	counter := 1
	for {
		select {
		case <-bt.done:
			return nil
		case <-ticker.C:
		}

		var c collector.Collector
		c = collector.NewUpstreamCollector()
		s, err := c.Collect(*bt.url)
		if err != nil {
			logp.Err("Fail to read Nginx upstream status: %v", err)
		}
		
		event := common.MapStr{
			"@timestamp": common.Time(time.Now()),
			"type":       b.Name,
			"counter":    counter,
		}
		bt.client.PublishEvent(event)
		logp.Info("Event sent")
		counter++
	}
}

func (bt *Nginxupstreambeat) Cleanup(b *beat.Beat) error {
	return nil
}

func (bt *Nginxupstreambeat) Stop() {
	close(bt.done)
}
