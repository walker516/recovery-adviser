package repository

import (
	"database/sql"
	"fmt"
	"io/ioutil"
	"recovery-adviser-api/domain"
	"strings"
)

// JobRepository構造体の定義
type JobRepository struct {
	db        *sql.DB
	queryPath string
}

// NewJobRepositoryの新規作成関数
func NewJobRepository(db *sql.DB, dbType string) (*JobRepository, error) {
	queryPath := fmt.Sprintf("infrastructure/sql/%s/job_queries.sql", dbType)
	return &JobRepository{db: db, queryPath: queryPath}, nil
}

// loadQueryは、指定されたクエリキーに対応するSQLクエリをファイルから読み込む
func (r *JobRepository) loadQuery(queryKey string) (string, error) {
	data, err := ioutil.ReadFile(r.queryPath)
	if err != nil {
		return "", fmt.Errorf("failed to read query file: %v", err)
	}

	queries := strings.Split(string(data), ";")
	for _, query := range queries {
		if strings.HasPrefix(query, "-- "+queryKey) {
			return strings.TrimSpace(query[len(queryKey)+3:]), nil
		}
	}

	return "", fmt.Errorf("query key %s not found in file", queryKey)
}

// GetRecoveryJobStatusは、指定された部品番号に対応するリカバリージョブのステータスを取得する
func (r *JobRepository) GetRecoveryJobStatus(seppenbuban string) (*domain.JobStatus, error) {
	query, err := r.loadQuery("GetRecoveryJobStatus")
	if err != nil {
		return nil, err
	}

	var jobStatus domain.JobStatus
	err = r.db.QueryRow(query, seppenbuban).Scan(
		&jobStatus.LatestProcessOrder,
		&jobStatus.LatestRegisterTimestamp,
		&jobStatus.LatestHost,
		&jobStatus.NeedsInvestigation,
		&jobStatus.NeedsDetailedReview,
		&jobStatus.JobNotCompletedCorrectly,
		&jobStatus.ErrorOccurredDuringJob,
	)
	if err != nil {
		return nil, err
	}
	return &jobStatus, nil
}

// GetJobQueueは、指定されたプロセスオーダーと部品番号に対応するジョブキューを取得する
func (r *JobRepository) GetJobQueue(processOrder, seppenbuban string) (*domain.JobQueue, error) {
	query, err := r.loadQuery("GetJobQueueByProcessOrder")
	if err != nil {
		return nil, err
	}

	var jobQueue domain.JobQueue
	err = r.db.QueryRow(query, processOrder).Scan(
		&jobQueue.ProcessOrder,
		&jobQueue.Status,
		&jobQueue.Host,
		&jobQueue.RegisterTimestamp,
		&jobQueue.Parameter,
	)
	if err != nil && err != sql.ErrNoRows {
		return nil, err
	}
	if err == sql.ErrNoRows && seppenbuban != "" {
		query, err = r.loadQuery("GetJobQueueBySeppenbuban")
		if err != nil {
			return nil, err
		}
		err = r.db.QueryRow(query, seppenbuban).Scan(
			&jobQueue.ProcessOrder,
			&jobQueue.Status,
			&jobQueue.Host,
			&jobQueue.RegisterTimestamp,
			&jobQueue.Parameter,
		)
		if err != nil {
			return nil, err
		}
	}
	return &jobQueue, nil
}

// UpdateJobQueueは、指定されたプロセスオーダーのジョブキューを更新する
func (r *JobRepository) UpdateJobQueue(processOrder string, jobQueue domain.JobQueue) error {
	query, err := r.loadQuery("UpdateJobQueue")
	if err != nil {
		return err
	}

	_, err = r.db.Exec(query, jobQueue.Status, jobQueue.Host, processOrder)
	return err
}

// GetJobLockは、指定されたプロセスオーダーに対応するジョブロックを取得する
func (r *JobRepository) GetJobLock(processOrder string) (*domain.JobLock, error) {
	query, err := r.loadQuery("GetJobLock")
	if err != nil {
		return nil, err
	}

	var jobLock domain.JobLock
	err = r.db.QueryRow(query, processOrder).Scan(&jobLock.ProcessOrder, &jobLock.LockTimestamp)
	if err != nil {
		return nil, err
	}
	return &jobLock, nil
}

// DeleteJobLockは、指定されたプロセスオーダーのジョブロックを削除する
func (r *JobRepository) DeleteJobLock(processOrder string) error {
	query, err := r.loadQuery("DeleteJobLock")
	if err != nil {
		return err
	}

	_, err = r.db.Exec(query, processOrder)
	return err
}
