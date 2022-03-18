package logger

import (
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

var ServiceInfo = serviceInfo{
	serviceName: "default",
}

type serviceInfo struct {
	serviceName string
}

func (s *serviceInfo) Setup(name string) {
	s.serviceName = name
}

func (s *serviceInfo) GetName() string {
	return s.serviceName
}

func New() *zerolog.Logger {
	logger := log.With().Str("name", ServiceInfo.GetName()).Logger()
	return &logger
}

func Setup(serviceName string, logLevel string) {
	ServiceInfo.Setup(serviceName)
	level, err := zerolog.ParseLevel(logLevel)
	if err != nil {
		level = zerolog.InfoLevel
	}
	zerolog.SetGlobalLevel(level)
}

func Info(msg string) {
	New().Info().Msg(msg)
}

func Debug(msg string) {
	New().Debug().Msg(msg)
}
