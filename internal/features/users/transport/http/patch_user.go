package users_transport_http

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/nurkenspashev92/todoapp/internal/core/domain"
	core_logger "github.com/nurkenspashev92/todoapp/internal/core/logger"
	core_http_request "github.com/nurkenspashev92/todoapp/internal/core/transport/http/request"
	core_http_response "github.com/nurkenspashev92/todoapp/internal/core/transport/http/response"
	core_http_types "github.com/nurkenspashev92/todoapp/internal/core/transport/http/types"
)

type PatchUserRequest struct {
	FullName    core_http_types.Nullable[string] `json:"full_name"`
	PhoneNumber core_http_types.Nullable[string] `json:"phone_number"`
}

func (r *PatchUserRequest) Validate() error {
	if r.FullName.Set {
		if r.FullName.Value == nil {
			return fmt.Errorf("FullNmae can't be NULL")
		}
	}

	if r.PhoneNumber.Set && r.PhoneNumber.Value != nil {
		phoneNumberLen := len([]rune(*r.PhoneNumber.Value))
		if phoneNumberLen < 10 || phoneNumberLen > 15 {
			return fmt.Errorf("PhoneNumber can't must be between 10 and 15")
		}

		if !strings.HasPrefix(*r.PhoneNumber.Value, "+") {
			return fmt.Errorf("PhoneNumber must startswith +")
		}
	}

	return nil
}

type UpdatedUserResponse UserDTOResponse

func (h *UsersHTTPHandler) PatchUserById(rw http.ResponseWriter, r *http.Request) {
	var request PatchUserRequest

	ctx := r.Context()
	log := core_logger.FromContext(ctx)
	responseHandler := core_http_response.NewHTTPResponseHander(log, rw)

	userId, err := getUserIdQueryParam(r)
	if err != nil {
		responseHandler.ErrorResponse(
			err,
			"faild to get user_id in service",
		)
		return
	}

	if err := core_http_request.DecodeAndValidate(r, &request); err != nil {
		responseHandler.ErrorResponse(
			err,
			"failed to decode and validate HTTP request",
		)
		return
	}

	userPatch := userPatchFromRequest(request)

	userDomain, err := h.userService.PatchUserById(ctx, userId, userPatch)
	if err != nil {
		responseHandler.ErrorResponse(
			err,
			"failed to get user in service",
		)
		return
	}

	response := UpdatedUserResponse(userDTOFromDomain(userDomain))

	responseHandler.JSONResponse(response, http.StatusOK)
}

func userPatchFromRequest(request PatchUserRequest) domain.UserPatch {
	return domain.NewUserPatch(
		request.FullName.ToDomain(),
		request.PhoneNumber.ToDomain(),
	)
}
