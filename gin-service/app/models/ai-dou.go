package models

// 爱享豆
type AiDou struct {
	Uid        int `gorm:"uid" json:"uid,omitempty"`
	AiDou      int `json:"ai_dou"`
	TotalAiDou int `json:"total_ai_dou"`
}

func (AiDou) TableName() string {
	return "a_ai_dou"
}
