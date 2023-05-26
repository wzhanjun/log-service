package client

import (
	"github.com/gookit/slog"
	hh "github.com/wzhanjun/log-service/client/handler"
)

func init() {
	slog.AddHandler(hh.NewGprcHandler())
}
