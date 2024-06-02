package usecase

import (
	"log"
	"recovery-adviser-api/domain"
)

// PartUseCaseインターフェースの定義
type PartUseCase interface {
	GetPartInfo(seppenbuban, usrID string) (*domain.PartInfo, error)
}

// partUseCase構造体の定義
type partUseCase struct {
	partRepo domain.PartRepository
}

// NewPartUseCaseの新規作成関数
func NewPartUseCase(pr domain.PartRepository) PartUseCase {
	return &partUseCase{partRepo: pr}
}

// 部品情報を取得するユースケース関数
func (pu *partUseCase) GetPartInfo(seppenbuban, usrID string) (*domain.PartInfo, error) {
	log.Printf("GetPartInfo called by user: %s for seppenbuban: %s", usrID, seppenbuban)
	return pu.partRepo.GetPartInfo(seppenbuban)
}
