package handler

import (
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
)

type SysdateHandler struct{}

func NewSysdateHandler() *SysdateHandler {
	return &SysdateHandler{}
}

func (sh *SysdateHandler) GetSysdate(c echo.Context) error {
	currentTime := time.Now().Format(time.RFC3339)
	return c.JSON(http.StatusOK, map[string]string{"sysdate": currentTime})
}