package client

import (
	"testing"
	"time"

	"github.com/gookit/slog"
)

func TestLog(t *testing.T) {
	for i := 0; i < 100; i++ {
		slog.Info("test", i)
		time.Sleep(time.Second)
	}
}
