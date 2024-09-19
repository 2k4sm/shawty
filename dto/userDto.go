package dto

type User struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

type UserSignup struct {
	Name     string `json:"name" validate:"required"`
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}

type UserLogin struct {
	Email string `json:"email" validate:"required, email"`
	Name  string `json:"name" validate:"required"`
}

type UpdateUserPass struct {
	ID          int    `json:"id" validate:"required"`
	Email       string `json:"email" validate:"required,email"`
	Password    string `json:"password" validate:"required"`
	OldPassword string `json:"oldPassword" validate:"required"`
}
