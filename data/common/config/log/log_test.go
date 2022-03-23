package log

import (
	"fmt"
	"go.uber.org/zap"
	"testing"
	"time"
)

func TestNewLogger(t *testing.T) {
	NewLogger()
	lg := ZapLogger
	type s struct {
		K string
		A int
	}
	ss := &s{
		K: "namemmm",
		A: 123,
	}
	for i := 0; i < 1; i++ {
		time.Sleep(time.Second)
		lg.Error(ss)
		lg.Debug(fmt.Sprint("debug log ", 1), zap.Int("line", 47))
		lg.Info(fmt.Sprint("Info log ", 2), zap.Any("level", "1231231231"))
		lg.Warn(fmt.Sprint("warn log ", 3), zap.String("level", `{"a":"4","b":"5"}`))
		lg.Error(fmt.Sprint("err log ", 4), zap.String("level", `{"a":"7","b":"8"}`))

	}
}
