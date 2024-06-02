package usecase

import (
	"log"
	"recovery-adviser-api/domain"
)

// JobUseCaseインターフェースの定義
type JobUseCase interface {
	GetRecoveryJobStatus(seppenbuban, usrID string) (*domain.JobStatus, error)
	GetJobQueue(processOrder, seppenbuban, usrID string) (*domain.JobQueue, error)
	UpdateJobQueue(processOrder string, jobQueue domain.JobQueue, usrID string) error
	GetJobLock(processOrder, usrID string) (*domain.JobLock, error)
	DeleteJobLock(processOrder, usrID string) error
}

// jobUseCase構造体の定義
type jobUseCase struct {
	jobRepo domain.JobRepository
}

// NewJobUseCaseの新規作成関数
func NewJobUseCase(jr domain.JobRepository) JobUseCase {
	return &jobUseCase{jobRepo: jr}
}

// リカバリージョブのステータスを取得するユースケース関数
func (ju *jobUseCase) GetRecoveryJobStatus(seppenbuban, usrID string) (*domain.JobStatus, error) {
	log.Printf("GetRecoveryJobStatus called by user: %s for seppenbuban: %s", usrID, seppenbuban)
	return ju.jobRepo.GetRecoveryJobStatus(seppenbuban)
}

// ジョブキューを取得するユースケース関数
func (ju *jobUseCase) GetJobQueue(processOrder, seppenbuban, usrID string) (*domain.JobQueue, error) {
	log.Printf("GetJobQueue called by user: %s for processOrder: %s, seppenbuban: %s", usrID, processOrder, seppenbuban)
	return ju.jobRepo.GetJobQueue(processOrder, seppenbuban)
}

// ジョブキューを更新するユースケース関数
func (ju *jobUseCase) UpdateJobQueue(processOrder string, jobQueue domain.JobQueue, usrID string) error {
	log.Printf("UpdateJobQueue called by user: %s for processOrder: %s with jobQueue: %+v", usrID, processOrder, jobQueue)
	return ju.jobRepo.UpdateJobQueue(processOrder, jobQueue)
}

// ジョブロックを取得するユースケース関数
func (ju *jobUseCase) GetJobLock(processOrder, usrID string) (*domain.JobLock, error) {
	log.Printf("GetJobLock called by user: %s for processOrder: %s", usrID, processOrder)
	return ju.jobRepo.GetJobLock(processOrder)
}

// ジョブロックを削除するユースケース関数
func (ju *jobUseCase) DeleteJobLock(processOrder, usrID string) error {
	log.Printf("DeleteJobLock called by user: %s for processOrder: %s", usrID, processOrder)
	return ju.jobRepo.DeleteJobLock(processOrder)
}
