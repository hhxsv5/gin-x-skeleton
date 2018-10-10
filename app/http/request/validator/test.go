package validator

type TestUpdate struct {
	Name   string `form:"name" json:"name" binding:"required"`
	Number uint64 `form:"number" json:"number" binding:"required"`
	Type   uint16 `form:"type" json:"type" binding:"required"`
}

type TestList struct {
	Ids     []uint64
	Keyword string `form:"keyword"`
	Status  uint16 `form:"status"`
}

type TestDetail struct {
	Id string `form:"id"`
}
