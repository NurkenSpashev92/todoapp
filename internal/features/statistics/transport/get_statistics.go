package statistics_transport_http

import (
	"fmt"
	"net/http"
	"time"

	"github.com/nurkenspashev92/todoapp/internal/core/domain"
	core_logger "github.com/nurkenspashev92/todoapp/internal/core/logger"
	core_http_request "github.com/nurkenspashev92/todoapp/internal/core/transport/http/request"
	core_http_response "github.com/nurkenspashev92/todoapp/internal/core/transport/http/response"
)

const (
	userIdQueryParamsKey = "user_id"
	fromQueryParamsKey   = "from"
	toQueryParamsKey     = "to"
)

type GetStatisticsResponse struct {
	TasksCreated               int      `json:"tasks_created"`
	TasksCompleted             int      `json:"tasks_completed"`
	TasksCompletedRate         *float64 `json:"tasks_completed_rate"`
	TasksAverageCompletionTime *string  `json:"tasks_average_completion_time"`
}

func (h *StatisticsHTTPHandler) GetStatistics(rw http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	log := core_logger.FromContext(ctx)
	responseHandler := core_http_response.NewHTTPResponseHander(log, rw)

	userId, from, to, err := getStatisticsQueryParams(r)
	if err != nil {
		responseHandler.ErrorResponse(
			err,
			"failed get 'user_id'/'from'/'to' query params",
		)
		return
	}

	statisticsDomain, err := h.statisticsService.GetStatistics(ctx, userId, from, to)
	if err != nil {
		responseHandler.ErrorResponse(
			err,
			"failed to get statistics in service",
		)
		return
	}

	response := statisticsDTOFromDomain(statisticsDomain)

	responseHandler.JSONResponse(response, http.StatusOK)
}

func getStatisticsQueryParams(r *http.Request) (*int, *time.Time, *time.Time, error) {
	userId, err := core_http_request.GetIntQueryParam(r, userIdQueryParamsKey)
	if err != nil {
		return nil, nil, nil, fmt.Errorf("get 'user_id' query param: %w", err)
	}

	from, err := core_http_request.GetDateQueryParam(r, fromQueryParamsKey)
	if err != nil {
		return nil, nil, nil, fmt.Errorf("get 'from' query param: %w", err)
	}

	to, err := core_http_request.GetDateQueryParam(r, toQueryParamsKey)
	if err != nil {
		return nil, nil, nil, fmt.Errorf("get 'to' query param: %w", err)
	}

	return userId, from, to, nil
}

func statisticsDTOFromDomain(statistics domain.Statistics) GetStatisticsResponse {
	var averageCompletionTime *string
	if statistics.TasksAverageCompletionTime != nil {
		duration := statistics.TasksAverageCompletionTime.String()
		averageCompletionTime = &duration
	}

	return GetStatisticsResponse{
		TasksCreated:               statistics.TasksCreated,
		TasksCompleted:             statistics.TasksCompleted,
		TasksCompletedRate:         statistics.TasksCompletedRate,
		TasksAverageCompletionTime: averageCompletionTime,
	}
}
