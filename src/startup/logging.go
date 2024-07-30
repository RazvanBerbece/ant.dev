package startup

import "log/slog"

func NewLayerLogger(logHandler slog.Handler, layer string) slog.Logger {
	return *slog.New(logHandler).With(
		"layer", layer,
	)
}

// Creates a new dependency (i.e repository, service, etc.) logger for the given output write handler,
// populating the log by default with the dependency layer and dependency name.
//
// For example, this could be used to create service-specific loggers
// which populate their output logs by default with the service name.
func NewDependencyLogger(logHandler slog.Handler, dependencyLayer string, dependencyName string) slog.Logger {

	if dependencyName == "" {
		return *slog.New(logHandler).With(
			"layer", dependencyLayer,
		)
	}

	return *slog.New(logHandler).With(
		"layer", dependencyLayer,
		"scope", dependencyName,
	)
}
