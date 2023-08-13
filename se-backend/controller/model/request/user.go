package request

type UserRequest struct {
	FirstName            string `json:"firstName" binding:"required,min=1,max=50"`
	LastName             string `json:"lastName" binding:"required,min=1,max=50"`
	Email                string `json:"email" binding:"required,email"`
	Username             string `json:"username" binding:"required,min=1,max=50,userformat"`
	Password             string `json:"password" binding:"required,strongpass"`
	PasswordConfirmation string `json:"passwordConfirmation" binding:"required,eqfield=Password"`
}
