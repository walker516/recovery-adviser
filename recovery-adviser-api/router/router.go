package router

import (
	"database/sql"
	"recovery-adviser-api/config"
	"recovery-adviser-api/domain"
	"recovery-adviser-api/infrastructure/repository"
	"recovery-adviser-api/interface/handler"
	"recovery-adviser-api/usecase"

	"github.com/labstack/echo/v4"
)

// NewRouter initializes the router
func NewRouter(db *sql.DB) *echo.Echo {
	e := echo.New()

	var partRepo domain.PartRepository
	var jobRepo domain.JobRepository

	dbType := config.ConfigData.Database.Type

	partRepo, _ = repository.NewPartRepository(db, dbType)
	jobRepo, _ = repository.NewJobRepository(db, dbType)

	partUseCase := usecase.NewPartUseCase(partRepo)
	jobUseCase := usecase.NewJobUseCase(jobRepo)

	sysdateHandler := handler.NewSysdateHandler()
	partHandler := handler.NewPartHandler(partUseCase)
	jobHandler := handler.NewJobHandler(jobUseCase)

	e.GET("/sysdate", sysdateHandler.GetSysdate)
	e.GET("/part/:seppenbuban", partHandler.GetPartInfo)
	e.GET("/recovery-job-status/:seppenbuban", jobHandler.GetRecoveryJobStatus)
	e.GET("/job-queue/:process_order", jobHandler.GetJobQueue)
	e.PUT("/job-queue/:process_order", jobHandler.UpdateJobQueue)
	e.GET("/job-lock/:process_order", jobHandler.GetJobLock)
	e.DELETE("/job-lock/:process_order", jobHandler.DeleteJobLock)

	return e
}
