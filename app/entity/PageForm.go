package entity

type PageForm struct {
	Page    int32 `gorm:"column:page;" json:"page" form:"page"`
	PageNum int32 `gorm:"column:page_num" json:"page_num" form:"page_num"`
}
