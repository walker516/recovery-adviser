package usecase

import (
	"recovery-adviser-api/domain"
)

// JobUseCaseインターフェースの定義
type JobUseCase interface {
	GetRecoveryJobStatus(seppenbuban string) (*domain.JobStatus, error)
	GetJobQueue(processOrder, seppenbuban string) (*domain.JobQueue, error)
	UpdateJobQueue(processOrder string, jobQueue domain.JobQueue) error
	GetJobLock(processOrder string) (*domain.JobLock, error)
	DeleteJobLock(processOrder string) error
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
func (ju *jobUseCase) GetRecoveryJobStatus(seppenbuban string) (*domain.JobStatus, error) {
	return ju.jobRepo.GetRecoveryJobStatus(seppenbuban)
}

// ジョブキューを取得するユースケース関数
func (ju *jobUseCase) GetJobQueue(processOrder, seppenbuban string) (*domain.JobQueue, error) {
	return ju.jobRepo.GetJobQueue(processOrder, seppenbuban)
}

// ジョブキューを更新するユースケース関数
func (ju *jobUseCase) UpdateJobQueue(processOrder string, jobQueue domain.JobQueue) error {
	return ju.jobRepo.UpdateJobQueue(processOrder, jobQueue)
}

// ジョブロックを取得するユースケース関数
func (ju *jobUseCase) GetJobLock(processOrder string) (*domain.JobLock, error) {
	return ju.jobRepo.GetJobLock(processOrder)
}

// ジョブロックを削除するユースケース関数
func (ju *jobUseCase) DeleteJobLock(processOrder string) error {
	return ju.jobRepo.DeleteJobLock(processOrder)
}
