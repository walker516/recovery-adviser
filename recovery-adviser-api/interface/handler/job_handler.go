package handler

import (
	"log"
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
	usrID := c.QueryParam("usr_id")
	log.Printf("GetRecoveryJobStatus called by user: %s", usrID)

	jobStatus, err := jh.JobUseCase.GetRecoveryJobStatus(seppenbuban, usrID)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Failed to get recovery job status: "+err.Error())
	}

	return c.JSON(http.StatusOK, jobStatus)
}

// ジョブキューを取得するハンドラ関数
func (jh *JobHandler) GetJobQueue(c echo.Context) error {
	processOrder := c.Param("process_order")
	seppenbuban := c.QueryParam("seppenbuban")
	usrID := c.QueryParam("usr_id")
	log.Printf("GetJobQueue called by user: %s", usrID)

	jobQueue, err := jh.JobUseCase.GetJobQueue(processOrder, seppenbuban, usrID)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Failed to get job queue: "+err.Error())
	}
	if jobQueue == nil {
		if processOrder != "" && seppenbuban != "" {
			return echo.NewHTTPError(http.StatusBadRequest, "Both process order and seppenbuban are specified. Please specify only one.")
		} else if processOrder == "" && seppenbuban == "" {
			return echo.NewHTTPError(http.StatusBadRequest, "Neither process order nor seppenbuban are specified. Please specify one.")
		}
		return echo.NewHTTPError(http.StatusNotFound, "No job queue information found")
	}

	return c.JSON(http.StatusOK, jobQueue)
}

// ジョブキューを更新するハンドラ関数
func (jh *JobHandler) UpdateJobQueue(c echo.Context) error {
	processOrder := c.Param("process_order")
	usrID := c.QueryParam("usr_id")
	log.Printf("UpdateJobQueue called by user: %s", usrID)

	var jobQueue domain.JobQueue
	if err := c.Bind(&jobQueue); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid request payload: "+err.Error())
	}

	if err := jh.JobUseCase.UpdateJobQueue(processOrder, jobQueue, usrID); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Failed to update job queue: "+err.Error())
	}

	return c.JSON(http.StatusOK, map[string]string{"message": "Job queue updated successfully"})
}

// ジョブロックを取得するハンドラ関数
func (jh *JobHandler) GetJobLock(c echo.Context) error {
	processOrder := c.Param("process_order")
	usrID := c.QueryParam("usr_id")
	log.Printf("GetJobLock called by user: %s", usrID)

	jobLock, err := jh.JobUseCase.GetJobLock(processOrder, usrID)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Failed to get job lock: "+err.Error())
	}

	if jobLock == nil {
		return echo.NewHTTPError(http.StatusNotFound, "No job lock information found for process order: "+processOrder)
	}

	return c.JSON(http.StatusOK, jobLock)
}

// ジョブロックを削除するハンドラ関数
func (jh *JobHandler) DeleteJobLock(c echo.Context) error {
	processOrder := c.Param("process_order")
	usrID := c.QueryParam("usr_id")
	log.Printf("DeleteJobLock called by user: %s", usrID)

	if err := jh.JobUseCase.DeleteJobLock(processOrder, usrID); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Failed to delete job lock: "+err.Error())
	}

	return c.JSON(http.StatusOK, map[string]string{"message": "Job lock deleted successfully"})
}
