package usecase

import (
	"recovery-adviser-api/domain"
)

// PartUseCaseインターフェースの定義
type PartUseCase interface {
	GetPartInfo(seppenbuban string) (*domain.PartInfo, error)
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
func (pu *partUseCase) GetPartInfo(seppenbuban string) (*domain.PartInfo, error) {
	return pu.partRepo.GetPartInfo(seppenbuban)
}
