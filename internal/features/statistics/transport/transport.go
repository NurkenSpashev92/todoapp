package statistics_transport_http

import (
	"net/http"

	core_http_server "github.com/nurkenspashev92/todoapp/internal/core/transport/http/server"
	statistics_service "github.com/nurkenspashev92/todoapp/internal/features/statistics/service"
)

type StatisticsHTTPHandler struct {
	statisticsService statistics_service.StatisticsServiceInterface
}

func NewStatisticsHTTPHandler(statisticsService statistics_service.StatisticsServiceInterface) *StatisticsHTTPHandler {
	return &StatisticsHTTPHandler{
		statisticsService: statisticsService,
	}
}

func (h *StatisticsHTTPHandler) Routes() []core_http_server.Route {
	return []core_http_server.Route{
		{
			Method:  http.MethodGet,
			Path:    "/statistics",
			Handler: h.GetStatistics,
		},
	}
}
