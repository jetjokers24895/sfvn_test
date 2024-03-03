package server

import (
	"sfvn_test/config"
	"sfvn_test/db"

	"github.com/labstack/echo/v4"
	"github.com/redis/rueidis"
	"gorm.io/gorm"
)

type Server struct {
	Echo   *echo.Echo
	DB     *gorm.DB
	Redis  rueidis.Client
	Config *config.Config
}

func NewServer(cfg *config.Config) *Server {
	return &Server{
		Echo:   echo.New(),
		DB:     db.InitPostgres(cfg),
		Redis:  db.InitRedisClient(cfg),
		Config: cfg,
	}
}

func (server *Server) Start(addr string) error {
	return server.Echo.Start(":" + addr)
}

type Data struct {
	Code      int         `json:"code"`
	ErrorCode string      `json:"error_code"`
	Message   string      `json:"message"`
	Data      interface{} `json:"data"`
}

func Response(c echo.Context, statusCode int, data interface{}) error {
	// nolint // context.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	// nolint // context.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE")
	// nolint // context.Writer.Header().Set("Access-Control-Allow-Headers", "Authorization")
	return c.JSON(statusCode, data)
}

func SuccessResponse(c echo.Context, statusCode int, message string, data interface{}) error {
	return Response(c, statusCode, Data{
		Code:    statusCode,
		Message: message,
		Data:    data,
	})
}

func ErrorResponse(c echo.Context, statusCode int, message string, errorCode string) error {
	if errorCode == "" {
		errorCode = "INTERNAL_SERVER_ERROR"
	}

	if message == "" {
		//TODO: get message from error code
		message = "Internal Server Error"
	}

	return Response(c, statusCode, Data{
		Code:      statusCode,
		Message:   message,
		ErrorCode: errorCode,
	})
}
