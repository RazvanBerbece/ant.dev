package startup

import (
	"log/slog"
)

type Environment struct {
	Port                       string
	useLocalStorageForComments string
}

func NewEnvironment(port string, useLocalStorageForComments string) Environment {
	return Environment{
		Port:                       port,
		useLocalStorageForComments: useLocalStorageForComments,
	}
}

func (e Environment) UsesLocalStorageForComments() bool {
	return e.useLocalStorageForComments == "1"
}

func (e Environment) LogStartupStatus(logger slog.Logger) {
	logger.Info("Webserver up and running",
		slog.String("Port", e.Port),
		slog.String("UseLocalStorageForComments", e.useLocalStorageForComments),
	)
}
