package request

type UserRequest struct {
	FirstName            string `json:"firstName" binding:"required,min=1,max=20"`
	LastName             string `json:"lastName" binding:"required,min=1,max=20"`
	Email                string `json:"email" binding:"required,email"`
	Username             string `json:"username" binding:"required,min=1,max=20"`
	Password             string `json:"password" binding:"required,strongpass"`
	PasswordConfirmation string `json:"passwordConfirmation" binding:"required,eqfield=Password"`
}
