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

type PartRepository struct {
	db        *sql.DB
	queryPath string
}

func NewPartRepository(db *sql.DB, dbType string) (*PartRepository, error) {
	queryPath := filepath.Join(config.ConfigData.SQLQueryPath,dbType, "part_queries.sql")
	return &PartRepository{db: db, queryPath: queryPath}, nil
}

func (r *PartRepository) loadQuery(queryKey string) (string, error) {
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

func (r *PartRepository) GetPartInfo(seppenbuban string) (*domain.PartInfo, error) {
	query, err := r.loadQuery("GetPartInfo")
	if err != nil {
		return nil, err
	}

	var partInfo domain.PartInfo
	err = r.db.QueryRow(query, seppenbuban).Scan(
		&partInfo.KBuban, &partInfo.Revision, &partInfo.KRevision,
	)
	if err != nil {
		return nil, err
	}
	return &partInfo, nil
}
