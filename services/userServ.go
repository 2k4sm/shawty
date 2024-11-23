package services

import (
	"os"
	"time"

	"github.com/2k4sm/shawty/dto"
	"github.com/2k4sm/shawty/models"
	"github.com/2k4sm/shawty/repositories"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

type UserServInterface interface {
	Login(*dto.UserAuth) (string, error)
	SignUp(*dto.UserAuth) (string, error)
	// UpdatePass(*dto.UpdateUserPass) (*models.User, error)
	// DeleteUser(id int) error
}

type UserServ struct {
	repo repositories.UserRepoInterface
}

func NewUserServ(userRepo repositories.UserRepoInterface) UserServInterface {
	return &UserServ{
		repo: userRepo,
	}
}

func (s UserServ) Login(user *dto.UserAuth) (string, error) {
	existingUser, err := s.repo.FindUserByEmail(user.Email)

	if err != nil {
		return "", err
	}

	err = bcrypt.CompareHashAndPassword([]byte(existingUser.Password),[]byte(user.Password))

	if err != nil {
		return "", err
	}

	jwtCreator := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"expiry" : time.Now().Add(time.Hour),
		"userId" : existingUser.ID,
	})

	token, err := jwtCreator.SignedString([]byte(os.Getenv("JWT_SIGN_KEY")))

	if err != nil {
		return "", err
	}

	return token, nil
}

func (s UserServ) SignUp(user *dto.UserAuth) (string, error) {
	cryptPass, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)

	if err != nil {
		return "", err
	}
	
	newUser := models.User{
		Email : user.Email,
		Name : user.Name,
		Password: string(cryptPass),
	}

	createdUser, err := s.repo.CreateUser(&newUser)

	if err != nil {
		return "", err
	}

	jwtCreator := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"expiry" : time.Now().Add(time.Hour),
		"userId" : createdUser.ID,
	})

	token, err := jwtCreator.SignedString([]byte(os.Getenv("JWT_SIGN_KEY")))

	if err != nil {
		return "",err
	}

	return token, nil
}

// func (s UserServ) UpdatePass(user *dto.UpdateUserPass) (*models.User, error) {
// 	return s.repo.UpdateUserPass(user.Email, user.Password)
// }

// func (s UserServ) DeleteUser(id int) error {
// 	return s.repo.DeleteUserById(id)
// }
