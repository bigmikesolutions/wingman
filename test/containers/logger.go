package containers

import (
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

type logger struct {
	logger zerolog.Logger
}

func (c logger) Printf(format string, v ...interface{}) {
	log.Logger.Debug().Msgf(format, v...)
}
