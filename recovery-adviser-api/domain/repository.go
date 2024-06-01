package domain

// PartRepositoryインターフェースの定義
type PartRepository interface {
	GetPartInfo(seppenbuban string) (*PartInfo, error)
}

// JobRepositoryインターフェースの定義
type JobRepository interface {
	GetRecoveryJobStatus(seppenbuban string) (*JobStatus, error)
	GetJobQueue(processOrder, seppenbuban string) (*JobQueue, error)
	UpdateJobQueue(processOrder string, jobQueue JobQueue) error
	GetJobLock(processOrder string) (*JobLock, error)
	DeleteJobLock(processOrder string) error
}
