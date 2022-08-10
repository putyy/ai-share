package models

type WalletLog struct {
	ID        int      `gorm:"primary_key" json:"id"`
	Uid       int      `json:"uid,omitempty"`
	Source    int      `json:"source,omitempty"`
	Balance   int      `json:"balance"`
	Oid       int      `json:"oid,omitempty"`
	Content   string   `json:"content"`
	Direction int      `json:"direction"`
	CreatedAt JsonTime `json:"created_at,omitempty"`
}

func (WalletLog) TableName() string {
	return "a_wallet_log"
}
