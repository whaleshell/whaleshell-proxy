package proxy

import (
	"context"

	"github.com/glaciforge/slogx"
)

// InitLogging installs a corporate slogx logger as slog default for the proxy process.
// Call from sandbox sidecar entrypoints before ListenAndServe.
func InitLogging(ctx context.Context, service string) *slogx.Logger {
	if service == "" {
		service = "osg-proxy"
	}
	log := slogx.SetupDefault(
		slogx.WithCorporateMasking(),
		slogx.WithFormat(slogx.FormatJSON),
		slogx.WithStackOnError(true),
		slogx.WithTraceContext(true),
	)
	log = log.With("service", service)
	go log.WatchLevelEnv(ctx, "OSG_LOG_LEVEL", 0)
	return log
}
