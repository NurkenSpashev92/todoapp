package core_http_request

import (
	"fmt"
	"net/http"
	"strconv"

	core_errors "github.com/nurkenspashev92/todoapp/internal/core/errors"
)

func GetIntPathValue(r *http.Request, key string) (int, error) {
	pathValue := r.PathValue(key)
	if pathValue == "" {
		return 0, fmt.Errorf(
			"no key=`%s` in path values: %w",
			key,
			core_errors.ErrInvlidArgument,
		)
	}

	value, err := strconv.Atoi(pathValue)
	if err != nil {
		return 0, fmt.Errorf(
			"path value=`%s` by key=`%s` not valid integer: $%: %w",
			value,
			key,
			err,
			core_errors.ErrInvlidArgument,
		)
	}

	return value, nil
}
