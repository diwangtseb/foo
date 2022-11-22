package pkg

import (
	"os"

	"github.com/go-kratos/kratos/v2/log"
	"golang.org/x/exp/slog"
)

var _ log.Logger = (*SLoger)(nil)

type SLoger struct {
	log *slog.Logger
}

func NewSlog() *SLoger {
	logger := slog.New(slog.NewJSONHandler(os.Stderr))
	slog.SetDefault(logger)
	return &SLoger{
		log: logger,
	}
}

func (s *SLoger) Log(level log.Level, keyvals ...interface{}) error {
	switch level {
	case log.LevelDebug:
		s.log.Debug("debug", keyvals)
	case log.LevelError:
		s.log.Error("error", nil, keyvals)
	case log.LevelInfo:
		s.log.Info("info", keyvals...)
	case log.LevelWarn:
		s.log.Warn("warn", keyvals...)
	}
	return nil
}
