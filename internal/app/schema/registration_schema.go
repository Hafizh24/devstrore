package schema

type RegisterReq struct {
	Email    string `validate:"required,email" json:"email"`
	Password string `validate:"required,min=8,alphanum" json:"password"`
	Username string `validate:"required,alphanum" json:"username"`
}
