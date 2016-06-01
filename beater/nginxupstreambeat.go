package beater

import (
	"fmt"
	"net/url"
	"time"

	"github.com/elastic/beats/libbeat/beat"
	"github.com/elastic/beats/libbeat/logp"
	"github.com/elastic/beats/libbeat/publisher"
	"github.com/elastic/2Fast2BCn/nginxupstreambeat/collector"
)

type Nginxupstreambeat struct {
	period time.Duration

	TbConfig ConfigSettings
	events   publisher.Client

	url  url.URL

	done chan struct{}
}

func (tb *Nginxupstreambeat) Config(b *beat.Beat) error {

	nginxupstreambeatSection := "nginxupstreambeat"

	rawNnginxupstreambeatConfig, err := b.RawConfig.Child(nginxupstreambeatSection, -1)
	if err != nil {
		logp.Err("Error reading configuration file: %v", err)
		return err
	}

	tb.TbConfig.Nnginxupstreambeat = defaultConfig
	err = rawNnginxupstreambeatConfig.Unpack(&tb.TbConfig.Nnginxupstreambeat)
	if err != nil {
		logp.Err("Error reading configuration file: %v", err)
		return err
	}

	nginxupstreambeatConfig := tb.TbConfig.Nnginxupstreambeat
	tb.period = nginxupstreambeatConfig.Period
	tb.url = nginxupstreambeatConfig.Url

	logp.Debug("nginxupstreambeat", "Init nginxupstreambeat")
	logp.Debug("nginxupstreambeat", "Period %v", tb.period)
	logp.Debug("nginxupstreambeat", "Url %t", tb.url)

	return nil
}

func (t *Nnginxupstreambeat) Setup(b *beat.Beat) error {
	t.events = b.Publisher.Connect()
	t.done = make(chan struct{})
	return nil
}

func (t *Nnginxupstreambeat) Run(b *beat.Beat) error {
	t.isAlive = true
	t.initProcStats()
	var err error
	for t.isAlive {
		time.Sleep(t.period)
		err = collector.newUpstreamCollector()
		if err != nil {
			logp.Err("Error reading status: %v", err)
		}
	return err
}

func (tb *Nnginxupstreambeat) Cleanup(b *beat.Beat) error {
	return nil
}

func (t *Nnginxupstreambeat) Stop() {
	logp.Info("Send stop signal to nginxupstreambeat main loop")
	t.isAlive = false
	close(t.done)
	t.events.Close()
}
