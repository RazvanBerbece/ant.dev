package startup

import (
	"log/slog"
)

type Environment struct {
	Port                       string
	env                        string
	useLocalStorageForComments string
}

func NewEnvironment(env string, port string, useLocalStorageForComments string) Environment {
	return Environment{
		Port:                       port,
		env:                        env,
		useLocalStorageForComments: useLocalStorageForComments,
	}
}

func (e Environment) UsesLocalStorageForComments() bool {
	return e.useLocalStorageForComments == "1"
}

func (e Environment) IsProduction() bool {
	return e.env == "production"
}

func (e Environment) LogStartupStatus(logger slog.Logger) {
	logger.Info("Webserver up and running\n",
		slog.String("Environment", e.env),
		slog.String("Port", e.Port),
		slog.String("UseLocalStorageForComments", e.useLocalStorageForComments),
	)
}
