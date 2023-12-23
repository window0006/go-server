package bindings

type FamilyListQuery struct {
	// 就算是 query，要使用 bind 并使用 tag 校验，也要用 form 标记
	Page int    `json:"page" form:"page" binding:"required,min=1"`
	Size int    `json:"size" form:"size" binding:"required,min=1,max=100"`
	Name string `json:"name" form:"name" binding:"max=32"`
}

type CreateFamilyBody struct {
	Name      string `json:"name" from:"name" binding:"required,max=32"`
	Patriarch string `json:"patriarch" from:"patriarch" binding:"required"`
	Members   string `json:"members" from:"members" binding:"required"`
}
