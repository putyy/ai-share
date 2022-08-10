package form

type CommonForm struct {
	Page           int    `form:"page,default=1"`
	Keyword        string `form:"keyword"`
	LastId         int    `form:"last_id,default=0"`
	CreatedAtStart string `form:"created_at_start" time_format:"2006-01-02" time_utc:"1"`
	CreatedAtEnd   string `form:"created_at_end" time_format:"2006-01-02" time_utc:"1"`
}
