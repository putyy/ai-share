package models

type OrderAiDou struct {
	ID    int `json:"id" gorm:"primary_key"`
	Oid   int `json:"oid,omitempty"`
	AiDou int `json:"ai_dou,omitempty"`
}

func (OrderAiDou) TableName() string {
	return "a_order_aidou"
}
