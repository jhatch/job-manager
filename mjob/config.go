package mjob

import (
	"time"

	"github.com/jeffrom/job-manager/mjob/client"
)

type ConsumerConfig struct {
	DequeueOpts     client.DequeueOpts `json:"dequeue_opts"`
	Concurrency     int                `json:"concurrency"`
	ShutdownTimeout time.Duration      `json:"shutdown_timeout"`
}

var defaultConsumerConfig ConsumerConfig = ConsumerConfig{
	Concurrency:     1,
	ShutdownTimeout: 15 * time.Second,
}
