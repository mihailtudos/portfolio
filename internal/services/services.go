package services

import "log/slog"

type Services struct {
	Assets AssetService
}

func NewServices(logger *slog.Logger) Services {
	return Services{
		Assets: AssetService{
			logger: logger,
		},
	}
}
