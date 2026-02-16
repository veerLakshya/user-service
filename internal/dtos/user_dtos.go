package dtos

type SignUpDto struct {
	Password  string `json:"password" validate:"required,min=8"` //validation
	Email     string `json:"email" validate:"required,email"`
	ClientId  string `json:"client_id"`
	VerifyUrl string `json:"verify_url" validate:"required,url"`
	Domain    string `json:"domain" validate:"required"`
}
