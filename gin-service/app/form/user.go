package form

type UserApiForm struct {
	Uid      int
	UserName string `form:"user_name" binding:"required"`
	NickName string `form:"nick_name" binding:"required"`
	HeadImg  string `form:"head_img" binding:"required"`
}

type LogApiForm struct {
	CommonForm
	Direction int `form:"type" binding:"required"`
}

type FansListApiForm struct {
	CommonForm
	Vip int `form:"vip" binding:"required"`
}

type OpenVipAdminForm struct {
	Uid       int    `form:"uid" binding:"required"`
	Level     int    `form:"level" binding:"required"`
	ShowPrice int    `form:"show_price,default=0"`
	Price     int    `form:"price,default=0"`
	Remark    string `form:"remark"`
	PayType   int    `form:"pay_type,default=3" binding:"required"`
}

type RechargeAiDouAdminForm struct {
	Uid     int    `form:"uid" binding:"required"`
	Price   int    `form:"price" binding:"required"`
	Remark  string `form:"remark"`
	PayType int    `form:"pay_type,default=3" binding:"required"`
	AiDou   int    `form:"ai_dou" binding:"required"`
}
