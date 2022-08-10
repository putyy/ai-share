package models

type Order struct {
	ID           int      `json:"id" gorm:"primary_key"`
	Uid          int      `json:"uid,omitempty"`
	OrderNo      string   `json:"order_no,omitempty"`
	Scene        int      `json:"scene,omitempty"`
	TotalPrice   int      `json:"total_price,omitempty"`
	ActualAmount int      `json:"actual_amount,omitempty"`
	PayType      int      `json:"pay_type,omitempty"`
	PayStatus    int      `json:"pay_status,omitempty"`
	Remark       string   `json:"remark"`
	AdminUid     int      `json:"admin_uid,omitempty"`
	CreatedAt    JsonTime `json:"created_at,omitempty"`
	UpdatedAt    JsonTime `json:"updated_at,omitempty"`
	DeletedAt    JsonTime `json:"deleted_at,omitempty"`
}

func (Order) TableName() string {
	return "a_order"
}
