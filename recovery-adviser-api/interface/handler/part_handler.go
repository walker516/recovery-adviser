package handler

import (
	"log"
	"net/http"
	"recovery-adviser-api/usecase"

	"github.com/labstack/echo/v4"
)

// PartHandler構造体の定義
type PartHandler struct {
	PartUseCase usecase.PartUseCase
}

// PartHandlerの新規作成関数
func NewPartHandler(pu usecase.PartUseCase) *PartHandler {
	return &PartHandler{PartUseCase: pu}
}

// 部品情報を取得するハンドラ関数
func (ph *PartHandler) GetPartInfo(c echo.Context) error {
	seppenbuban := c.Param("seppenbuban")
	usrID := c.QueryParam("usr_id")
	log.Printf("GetPartInfo called by user: %s", usrID)

	partInfo, err := ph.PartUseCase.GetPartInfo(seppenbuban, usrID)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Failed to get part information: "+err.Error())
	}

	if partInfo == nil {
		return echo.NewHTTPError(http.StatusNotFound, "No part information found")
	}

	return c.JSON(http.StatusOK, partInfo)
}
