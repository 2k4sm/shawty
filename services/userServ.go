package services

import (
	"github.com/2k4sm/shawty/dto"
	"github.com/2k4sm/shawty/models"
	"github.com/2k4sm/shawty/repositories"
	"golang.org/x/crypto/bcrypt"
)

type UserServInterface interface {
	Login(*dto.UserLogin) (*models.User, error)
	SignUp(*dto.UserSignup) (*models.User, error)
	UpdatePass(*dto.UpdateUserPass) (*models.User, error)
	DeleteUser(id int) error
}

type UserServ struct {
	repo repositories.UserRepoInterface
}

func (s *UserServ) NewUserServ(userRepo repositories.UserRepoInterface) UserServInterface {
	return &UserServ{
		repo: userRepo,
	}
}

func (s *UserServ) Login(user *dto.UserLogin) (*models.User, error) {

	return s.repo.FindUserByEmail(user.Email)
}

func (s *UserServ) SignUp(user *dto.UserSignup) (*models.User, error) {

	password, _ := bcrypt.GenerateFromPassword([]byte(user.Password), 14)
	newUserModel := models.User{}

	newUserModel.Name = user.Name
	newUserModel.Email = user.Email
	newUserModel.Password = string(password)

	return s.repo.CreateUser(&newUserModel)
}

func (s *UserServ) UpdatePass(user *dto.UpdateUserPass) (*models.User, error) {
	return s.repo.UpdateUserPass(user.Email, user.Password)
}

func (s *UserServ) DeleteUser(id int) error {
	return s.repo.DeleteUserById(id)
}
