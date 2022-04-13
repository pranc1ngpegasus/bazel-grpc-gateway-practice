package logger

import (
	"os"
	"sync"

	"github.com/rs/zerolog"
)

var (
	once   sync.Once
	logger zerolog.Logger
)

func New() zerolog.Logger {
	once.Do(func() {
		logger = zerolog.New(os.Stderr).
			With().
			Caller().
			Stack().
			Timestamp().
			Logger()
	})

	return logger
}
