package statsdconnect

import (
	"fmt"
	"time"

	"github.com/alexcesaro/statsd"
)

const FlushPeriod = 500 * time.Millisecond

var StatsdClient statsd.Client

func InitClient(host string, port int, prefix string) error {
	addr := fmt.Sprintf("%s:%d", host, port)

	client, err := statsd.New(
		statsd.Address(addr),
		statsd.FlushPeriod(FlushPeriod),
		statsd.Prefix(prefix))
	if err != nil {
		return err
	}

	StatsdClient = *client
	return nil
}
