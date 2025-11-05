package client

import (
	"fmt"
	"io"

	"github.com/gookit/slog"
)

var (
	LabelField = "label"
)

var logger *slog.SugaredLogger

func init() {
	logger = slog.Std()
	if Cfg.OutputType == 0 {
		logger.Output = io.Discard
	}
	logger.AddHandler(NewGprcHandler())
}

func Label(val string) *slog.Record {
	return slog.Std().WithField(LabelField, val)
}

func Std() *slog.SugaredLogger {
	return logger
}

func StrCaller(r *slog.Record) string {
	return fmt.Sprintf("file:%s, line:%d, func:%s", r.Caller.File, r.Caller.Line, r.Caller.Func.Name())
}

func StrLabel(r *slog.Record) string {
	label, _ := r.Fields[LabelField].(string)
	return label
}
