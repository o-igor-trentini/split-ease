package response

type UserResponse struct {
	ID        uint64 `json:"id"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Email     string `json:"email"`
	Username  string `json:"username"`
	Verified  bool   `json:"verified"`
}
