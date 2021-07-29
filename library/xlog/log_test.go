package xlog

import (
	"testing"
	"time"

	"github.com/sirupsen/logrus"
)

func TestNewLogger(t *testing.T) {
	c := Config{
		Path:         "runtime/logs/",
		FilePrefix:   "log%Y%m%d%H%M%S",
		LevelMode:    "release",
		RotationTime: 10,
	}

	logging := Init(&c, "test")
	for {
		logging.WithFields(logrus.Fields{
			"animal": "walrus",
			"size":   10,
		}).Warn("A group of walrus emerges from the ocean")
		time.Sleep(time.Duration(2) * time.Second)
	}
}
