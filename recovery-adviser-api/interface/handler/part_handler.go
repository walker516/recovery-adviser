package handler

import (
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

	partInfo, err := ph.PartUseCase.GetPartInfo(seppenbuban)
	if err != nil {
		return c.String(http.StatusInternalServerError, "Database query error: "+err.Error())
	}

	if partInfo == nil {
		return c.String(http.StatusNotFound, "No part information found")
	}

	return c.JSON(http.StatusOK, partInfo)
}
