package handlers

import (
	// "net/http"

	"fmt"
	"net/http"
	s "sfvn_test/server"

	dtos "sfvn_test/dtos/requests"

	_userService "sfvn_test/services/user"

	"github.com/labstack/echo/v4"
)

type userHandler struct {
	server      *s.Server
	userService _userService.ServiceWrapper
}

func NewUserHandler(server *s.Server) *userHandler {
	return &userHandler{
		server:      server,
		userService: _userService.ProviderService(server.DB, server.Config, server.Redis),
	}
}

func (u *userHandler) GetHistories(c echo.Context) error {
	uid := fmt.Sprintf("%v", c.Get("uid"))
	var queryInput = new(dtos.GetHistories)
	if err := c.Bind(queryInput); err != nil {
		return s.ErrorResponse(c, 400, "bad_request", "bad_request")
	}

	if err := queryInput.Validate(); err != nil {
		return s.ErrorResponse(c, 400, err.Error(), "bad_request")
	}

	rs, err := u.userService.GetHistoriesOfSymbol(c.Request().Context(), queryInput, uid)
	if err != nil {
		return s.ErrorResponse(c, 400, err.Error(), "internal_error")
	}

	return s.SuccessResponse(c, http.StatusOK, "", rs)
}
