package client

import (
	"fmt"

	"github.com/gookit/slog"
)

var (
	LabelField = "label"
)

func init() {
	slog.AddHandler(NewGprcHandler())
}

func Label(val string) *slog.Record {
	return slog.Std().WithField(LabelField, val)
}

func Std() *slog.SugaredLogger {
	return slog.Std()
}

func StrCaller(r *slog.Record) string {
	return fmt.Sprintf("file:%s, line:%d, func:%s", r.Caller.File, r.Caller.Line, r.Caller.Func.Name())
}

func StrLabel(r *slog.Record) string {
	label, _ := r.Fields[LabelField].(string)
	return label
}
