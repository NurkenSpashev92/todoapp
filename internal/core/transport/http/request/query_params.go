package core_http_request

import (
	"fmt"
	"net/http"
	"strconv"
	"time"

	core_errors "github.com/nurkenspashev92/todoapp/internal/core/errors"
)

func GetIntQueryParam(r *http.Request, key string) (*int, error) {
	param := r.URL.Query().Get(key)
	if param == "" {
		return nil, nil
	}

	val, err := strconv.Atoi(param)
	if err != nil {
		return nil, fmt.Errorf(
			"param=`%s` by key=`%s` not a valid integer: %v: %w",
			param,
			key,
			err,
			core_errors.ErrInvlidArgument,
		)
	}

	return &val, nil
}

func GetDateQueryParam(r *http.Request, key string) (*time.Time, error) {
	param := r.URL.Query().Get(key)
	if param == "" {
		return nil, nil
	}

	val, err := time.Parse(time.DateOnly, param)
	if err != nil {
		return nil, fmt.Errorf(
			"param=`%s` by key=`%s` not a valid date (YYYY-MM-DD): %v: %w",
			param,
			key,
			err,
			core_errors.ErrInvlidArgument,
		)
	}

	return &val, nil
}
