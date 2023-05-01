package schema

type GetProductResp struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Currency    string `json:"currency"`
	Price       int    `json:"price"`
	TotalStock  int    `json:"total_stock"`
	IsActive    bool   `json:"is_active"`
	CategoryID  int    `json:"category_id"`
}

type CreateProductReq struct {
	Name        string `validate:"required" json:"name"`
	Description string `validate:"required" json:"description"`
	Currency    string `validate:"required" json:"currency"`
	Price       int    `validate:"required" json:"price"`
	TotalStock  int    `validate:"required" json:"total_stock"`
	IsActive    bool   `validate:"required" json:"is_active"`
	CategoryID  int    `validate:"required" json:"category_id"`
}
type UpdateProductReq struct {
	Name        string `validate:"required" json:"name"`
	Description string `validate:"required" json:"description"`
	Currency    string `validate:"required" json:"currency"`
	Price       int    `validate:"required" json:"price"`
	TotalStock  int    `validate:"required" json:"total_stock"`
	IsActive    bool   `validate:"required" json:"is_active"`
	CategoryID  int    `validate:"required" json:"category_id"`
}

type GetDetailResp struct {
	ID          int      `json:"id"`
	Name        string   `json:"name"`
	Description string   `json:"description"`
	Currency    string   `json:"currency"`
	Price       int      `json:"price"`
	TotalStock  int      `json:"total_stock"`
	IsActive    bool     `json:"is_active"`
	Category    Category `json:"category"`
}

type Category struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}
