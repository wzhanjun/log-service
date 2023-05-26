package client

import (
	"testing"
	"time"
)

func TestLog(t *testing.T) {
	for i := 0; i < 3; i++ {
		// slog.Info("test", i)
		Label("test").Info("test", i)
		time.Sleep(time.Second)
	}
}
