package handlers

import (
	"net/http"
	s "sfvn_test/server"

	"github.com/labstack/echo/v4"
)

type AnonymHandler struct {
	server *s.Server
}

func NewAnonymHandler(server *s.Server) *AnonymHandler {
	return &AnonymHandler{server: server}
}

func (h *AnonymHandler) HealthCheck(c echo.Context) error {

	return s.SuccessResponse(c, http.StatusOK, "I'm fine! let's go", nil)
}
