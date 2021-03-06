package config

import (
	"os"
	"strings"
	"time"

	"github.com/caarlos0/env"
	"github.com/dalot/go-skeleton/pkg/constants"
	"github.com/rs/zerolog"
)

// Config - environment variables are parsed to this struct
type Config struct {
	AppName    string `env:"APP_NAME" envDefault:"boilerplate"`
	ServerPort int    `env:"PORT" envDefault:"8000"`
	Env        string `env:"ENV" envDefault:"env not set"`
	LogLevel   string `env:"LOG_LEVEL" envDefault:"debug"`
	LogOutput  string `env:"LOG_OUTPUT" envDefault:"console"`

	IdleTimeout time.Duration `env:"IDLE_TIMEOUT" envDefault:"5s"`
	// WriteTimeout maximum time the server will handle a request before timing out writes of the response.
	// It must be bigger than RequestTimeout
	WriteTimeout time.Duration `env:"WRITE_TIMEOUT" envDefault:"4s"`
	// RequestTimeout the timeout for the incoming request set on the request handler
	RequestTimeout    time.Duration `env:"REQUEST_TIMEOUT" envDefault:"2s"`
	ReadHeaderTimeout time.Duration `env:"READ_HEADER_TIMEOUT" envDefault:"1s"`

	// ShutdownTimeout the time the sever will wait server.Shutdown to return
	ShutdownTimeout time.Duration `env:"SHUTDOWN_TIMEOUT" envDefault:"6s"`
}

// Parse environment variables, returns (guess what?) an error if an error occurs
func Parse() (Config, error) {
	confs := Config{}
	err := env.Parse(&confs)
	return confs, err
}

// Logger returns a initialised zerolog.Logger
func (c Config) Logger() zerolog.Logger {
	logLevelOk := true
	logLevel, err := zerolog.ParseLevel(c.LogLevel)
	if err != nil {
		logLevel = zerolog.InfoLevel
		logLevelOk = false
	}

	zerolog.SetGlobalLevel(logLevel)
	zerolog.TimestampFieldName = constants.LogKeyTimestamp

	host, _ := os.Hostname()
	logger := zerolog.New(os.Stdout).
		With().
		Timestamp().
		Str(constants.LogKeyApp, c.AppName).
		Str(constants.LogKeyHost, host).
		Str(constants.LogKeyEnv, c.Env).
		Logger()

	if strings.ToUpper(c.LogOutput) == "CONSOLE" {
		logger = logger.Output(zerolog.ConsoleWriter{Out: os.Stderr})
	}

	if !logLevelOk {
		logger.Warn().Err(err).Msgf("%s is not a valid zerolog log level, defaulting to info", c.LogLevel)
	}

	return logger
}
