package models

type AiDouLog struct {
	ID        int      `gorm:"id" json:"id"`
	Uid       int      `json:"uid"`
	Sid       int      `json:"sid"`
	Source    int      `json:"source,omitempty"`
	AiDou     int      `json:"ai_dou,omitempty"`
	Content   string   `json:"content"`
	Direction int      `json:"direction"`
	CreatedAt JsonTime `json:"created_at,omitempty"`
}

func (AiDouLog) TableName() string {
	return "a_ai_dou_log"
}
