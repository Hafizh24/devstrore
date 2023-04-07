package schema

type GetCategoryResp struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

type CreateCategoryReq struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}

type UpdateCategoryReq struct {
	ID          int    `json:"id"`
	Name        string `json:"name" binding:"required"`
	Description string `json:"description" binding:"required"`
}
