package repository

import (
	"database/sql"
	"fmt"
	"io/ioutil"
	"path/filepath"
	"recovery-adviser-api/config"
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
	queryPath := filepath.Join(config.ConfigData.SQLQueryPath, dbType, "job_queries.sql")
	return &JobRepository{db: db, queryPath: queryPath}, nil
}

// loadQueryは、指定されたクエリキーに対応するSQLクエリをファイルから読み込む
func (r *JobRepository) loadQuery(queryKey string) (string, error) {
	data, err := ioutil.ReadFile(r.queryPath)
	if err != nil {
		return "", fmt.Errorf("failed to read query file: %v", err)
	}

	queries := strings.Split(string(data), "--")
	for _, query := range queries {
		query = strings.TrimSpace(query)
		if strings.HasPrefix(query, queryKey) {
			return strings.TrimSpace(query[len(queryKey):]), nil
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
	if err := r.db.QueryRow(query, seppenbuban).Scan(
		&jobStatus.LatestProcessOrder,
		&jobStatus.LatestRegisterTimestamp,
		&jobStatus.LatestHost,
		&jobStatus.NeedsInvestigation,
		&jobStatus.NeedsDetailedReview,
		&jobStatus.JobNotCompletedCorrectly,
		&jobStatus.ErrorOccurredDuringJob,
	); err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, fmt.Errorf("query execution failed: %v", err)
	}
	return &jobStatus, nil
}

// GetJobQueueは、指定されたプロセスオーダーと部品番号に対応するジョブキューを取得する
func (r *JobRepository) GetJobQueue(processOrder, seppenbuban string) (*domain.JobQueue, error) {
	if processOrder != "" && seppenbuban != "" {
		return nil, fmt.Errorf("both processOrder and seppenbuban are specified")
	}
	if processOrder == "" && seppenbuban == "" {
		return nil, fmt.Errorf("neither processOrder nor seppenbuban are specified")
	}

	var query string
	var err error
	var jobQueue domain.JobQueue

	if processOrder != "" {
		query, err = r.loadQuery("GetJobQueueByProcessOrder")
		if err != nil {
			return nil, fmt.Errorf("failed to load query: %v", err)
		}
		err = r.db.QueryRow(query, processOrder).Scan(
			&jobQueue.ProcessOrder,
			&jobQueue.Status,
			&jobQueue.Host,
			&jobQueue.RegisterTimestamp,
			&jobQueue.Parameter,
		)
	} else {
		query, err = r.loadQuery("GetJobQueueBySeppenbuban")
		if err != nil {
			return nil, fmt.Errorf("failed to load query: %v", err)
		}
		err = r.db.QueryRow(query, seppenbuban).Scan(
			&jobQueue.ProcessOrder,
			&jobQueue.Status,
			&jobQueue.Host,
			&jobQueue.RegisterTimestamp,
			&jobQueue.Parameter,
		)
	}

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, fmt.Errorf("query execution failed: %v", err)
	}

	return &jobQueue, nil
}

// UpdateJobQueueは、指定されたプロセスオーダーのジョブキューを更新する
func (r *JobRepository) UpdateJobQueue(processOrder string, jobQueue domain.JobQueue) error {
	query, err := r.loadQuery("UpdateJobQueue")
	if err != nil {
		return err
	}

	if _, err := r.db.Exec(query, jobQueue.Status, jobQueue.Host, processOrder); err != nil {
		return fmt.Errorf("failed to update job queue: %v", err)
	}
	return nil
}

// GetJobLockは、指定されたプロセスオーダーに対応するジョブロックを取得する
func (r *JobRepository) GetJobLock(processOrder string) (*domain.JobLock, error) {
	query, err := r.loadQuery("GetJobLock")
	if err != nil {
		return nil, err
	}

	var jobLock domain.JobLock
	if err := r.db.QueryRow(query, processOrder).Scan(&jobLock.ProcessOrder, &jobLock.LockTimestamp); err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, fmt.Errorf("query execution failed: %v", err)
	}
	return &jobLock, nil
}

// DeleteJobLockは、指定されたプロセスオーダーのジョブロックを削除する
func (r *JobRepository) DeleteJobLock(processOrder string) error {
	query, err := r.loadQuery("DeleteJobLock")
	if err != nil {
		return err
	}

	if _, err := r.db.Exec(query, processOrder); err != nil {
		return fmt.Errorf("failed to delete job lock: %v", err)
	}
	return nil
}
