package domain

// JobStatusの構造体定義
type JobStatus struct {
	LatestProcessOrder       string `json:"latest_process_order"`
	LatestRegisterTimestamp  string `json:"latest_register_timestamp"`
	LatestHost               string `json:"latest_host"`
	NeedsInvestigation       int    `json:"needs_investigation"`
	NeedsDetailedReview      int    `json:"needs_detailed_review"`
	JobNotCompletedCorrectly int    `json:"job_not_completed_correctly"`
	ErrorOccurredDuringJob   int    `json:"error_occurred_during_job"`
}

// JobQueueの構造体定義
type JobQueue struct {
	ProcessOrder     string `json:"process_order"`
	Status           int `json:"status"`
	Host             string `json:"host"`
	RegisterTimestamp string `json:"register_timestamp"`
	Parameter        string `json:"parameter"`
}

// JobLockの構造体定義
type JobLock struct {
	ProcessOrder  string `json:"process_order"`
	LockTimestamp string `json:"lock_timestamp"`
}
