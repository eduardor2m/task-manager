package request

type UserDTO struct {
	Username string `json:"username" example:"eduardor2m"`
	Email    string `json:"email" example:"eduardo@gmail.com"`
	Password string `json:"password" example:"123456"`
}
