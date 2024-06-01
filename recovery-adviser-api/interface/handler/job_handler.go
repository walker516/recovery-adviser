package handler

import (
	"net/http"
	"recovery-adviser-api/domain"
	"recovery-adviser-api/usecase"

	"github.com/labstack/echo/v4"
)

// JobHandler構造体の定義
type JobHandler struct {
	JobUseCase usecase.JobUseCase
}

// JobHandlerの新規作成関数
func NewJobHandler(ju usecase.JobUseCase) *JobHandler {
	return &JobHandler{JobUseCase: ju}
}

// リカバリージョブのステータスを取得するハンドラ関数
func (jh *JobHandler) GetRecoveryJobStatus(c echo.Context) error {
	seppenbuban := c.Param("seppenbuban")

	jobStatus, err := jh.JobUseCase.GetRecoveryJobStatus(seppenbuban)
	if err != nil {
		return c.String(http.StatusInternalServerError, "Database query error")
	}

	return c.JSON(http.StatusOK, jobStatus)
}

// ジョブキューを取得するハンドラ関数
func (jh *JobHandler) GetJobQueue(c echo.Context) error {
	processOrder := c.Param("process_order")
	seppenbuban := c.QueryParam("seppenbuban")

	jobQueue, err := jh.JobUseCase.GetJobQueue(processOrder, seppenbuban)
	if err != nil {
		return c.String(http.StatusInternalServerError, "No job queue information found")
	}

	return c.JSON(http.StatusOK, jobQueue)
}

// ジョブキューを更新するハンドラ関数
func (jh *JobHandler) UpdateJobQueue(c echo.Context) error {
	processOrder := c.Param("process_order")

	var jobQueue domain.JobQueue
	if err := c.Bind(&jobQueue); err != nil {
		return c.String(http.StatusBadRequest, "Invalid request payload")
	}

	err := jh.JobUseCase.UpdateJobQueue(processOrder, jobQueue)
	if err != nil {
		return c.String(http.StatusInternalServerError, "Database update error")
	}

	return c.JSON(http.StatusOK, map[string]string{"message": "Job queue updated successfully"})
}

// ジョブロックを取得するハンドラ関数
func (jh *JobHandler) GetJobLock(c echo.Context) error {
	processOrder := c.Param("process_order")

	jobLock, err := jh.JobUseCase.GetJobLock(processOrder)
	if err != nil {
		return c.String(http.StatusInternalServerError, "No job lock information found")
	}

	return c.JSON(http.StatusOK, jobLock)
}

// ジョブロックを削除するハンドラ関数
func (jh *JobHandler) DeleteJobLock(c echo.Context) error {
	processOrder := c.Param("process_order")

	err := jh.JobUseCase.DeleteJobLock(processOrder)
	if err != nil {
		return c.String(http.StatusInternalServerError, "Database delete error")
	}

	return c.JSON(http.StatusOK, map[string]string{"message": "Job lock deleted successfully"})
}
