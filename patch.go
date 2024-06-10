package go_websockify

import (
	"context"
	"fmt"
	"log/slog"
)

func Start(ip string, port int) {
	if err := server.Shutdown(context.Background()); err != nil {
		slog.Warn("server.Shutdown", slog.Any("error", err))
	}

	ctx, stopHTTP = context.WithCancel(context.Background())

	config.bindAddr = "0.0.0.0:6080"
	config.remoteAddr = fmt.Sprintf("%s:%d", ip, port)
	StartHTTP()
}
