package logger

import (
	"github.com/vladimish/soo/pkg/configurator"
	"go.uber.org/zap"
)

var L = &zap.Logger{}

// Initialize needs to be called once on program start.
func Initialize() {
	var err error
	switch configurator.Cfg.Logging.LogLevel {
	case configurator.DEBUG:
		L, err = zap.NewDevelopment()
		if err != nil {
			panic(err)
		}
	case configurator.RELEASE:
		L, err = zap.NewProduction()
		if err != nil {
			panic(err)
		}
	}
}
