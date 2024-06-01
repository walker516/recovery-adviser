package domain

// PartInfoの構造体定義
type PartInfo struct {
	KBuban    string `json:"kbuban"`
	Revision  string `json:"revision"`
	KRevision string `json:"krevision"`
}
