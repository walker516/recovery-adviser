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

// PartRepository構造体の定義
type PartRepository struct {
	db        *sql.DB
	queryPath string
}

// NewPartRepositoryの新規作成関数
func NewPartRepository(db *sql.DB, dbType string) (*PartRepository, error) {
	queryPath := filepath.Join(config.ConfigData.SQLQueryPath, dbType, "part_queries.sql")
	return &PartRepository{db: db, queryPath: queryPath}, nil
}

// loadQueryは、指定されたクエリキーに対応するSQLクエリをファイルから読み込む
func (r *PartRepository) loadQuery(queryKey string) (string, error) {
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

// GetPartInfoは、指定された部品番号に対応する部品情報を取得する
func (r *PartRepository) GetPartInfo(seppenbuban string) (*domain.PartInfo, error) {
	query, err := r.loadQuery("GetPartInfo")
	if err != nil {
		return nil, err
	}

	var partInfo domain.PartInfo
	if err := r.db.QueryRow(query, seppenbuban).Scan(
		&partInfo.KBuban, &partInfo.Revision, &partInfo.KRevision,
	); err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, fmt.Errorf("query execution failed: %v", err)
	}
	return &partInfo, nil
}
