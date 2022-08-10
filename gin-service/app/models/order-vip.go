package models

type OrderVip struct {
	ID    int `json:"id" gorm:"primary_key"`
	Oid   int `json:"oid,omitempty"`
	Level int `json:"level,omitempty"`
}

func (OrderVip) TableName() string {
	return "a_order_vip"
}
