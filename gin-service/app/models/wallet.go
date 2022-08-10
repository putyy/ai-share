package models

type Wallet struct {
	Uid          int `gorm:"primary_key" json:"uid,omitempty"`
	Balance      int `json:"balance"`
	TotalBalance int `json:"total_balance"`
}

func (Wallet) TableName() string {
	return "a_wallet"
}
