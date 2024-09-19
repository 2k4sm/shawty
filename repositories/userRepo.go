package repositories

import (
	"errors"
	"fmt"

	"github.com/2k4sm/shawty/models"
	"gorm.io/gorm"
)

type UserRepoInterface interface {
	FindUserById(id int) (*models.User, error)
	FindUserByEmail(email string) (*models.User, error)
	FindUserByUname(uname string) (*models.User, error)
	CreateUser(user *models.User) (*models.User, error)
	UpdateUserPass(email string, newPass string) (*models.User, error)
	DeleteUserById(id int) error
}

type UserRepo struct {
	DB *gorm.DB
}

func (u *UserRepo) NewUserRepo(db *gorm.DB) UserRepoInterface {
	return &UserRepo{
		DB: db,
	}
}

func (u *UserRepo) FindUserById(id int) (*models.User, error) {
	var userToFind models.User

	txn := u.DB.Find(&userToFind, id)

	if userToFind.ID == 0 || txn.Error != nil {
		return nil, fmt.Errorf("user not found : %v", txn.Error)
	}

	return &userToFind, nil
}

func (u *UserRepo) FindUserByEmail(email string) (*models.User, error) {
	var userToFind models.User

	txn := u.DB.Find(&userToFind, "email = ?", email)

	if userToFind.ID == 0 || txn.Error != nil {
		return nil, fmt.Errorf("user not found : %v", txn.Error)
	}

	return &userToFind, nil
}

func (u *UserRepo) FindUserByUname(name string) (*models.User, error) {
	var userToFind models.User

	txn := u.DB.Find(&userToFind, "name = ?", name)

	if userToFind.ID == 0 || txn.Error != nil {
		return nil, fmt.Errorf("user not found : %v", txn.Error)
	}

	return &userToFind, nil
}

func (u *UserRepo) CreateUser(user *models.User) (*models.User, error) {
	var newUser models.User

	u.DB.Find(&newUser, "email = ?", user.Email)

	if newUser.ID != 0 {
		return nil, errors.New("user already exists")
	}

	txn := u.DB.Create(user)

	if txn.Error != nil {
		return nil, fmt.Errorf("error creating user: %v", txn.Error)
	}

	return user, nil
}

func (u *UserRepo) UpdateUserPass(email string, newPass string) (*models.User, error) {
	var userToUpdate models.User

	txn := u.DB.Find(&userToUpdate, "email = ?", email)

	if userToUpdate.ID == 0 || txn.Error != nil {
		return nil, fmt.Errorf("user not found for updating : %v", txn.Error)
	}

	userToUpdate.Password = newPass

	txn = u.DB.Save(&userToUpdate)

	if txn.Error != nil {
		return nil, fmt.Errorf("error updating user: %v", txn.Error)
	}

	return &userToUpdate, nil
}

func (u *UserRepo) DeleteUserById(id int) error {
	var userToDel models.User

	txn := u.DB.Find(&userToDel, id)

	if userToDel.ID == 0 || txn.Error != nil {
		return fmt.Errorf("user not found for deleting: %v", txn.Error)
	}

	txn = u.DB.Delete(&userToDel)

	if txn.Error != nil {
		return fmt.Errorf("error deleting user with id %d : %v", id, txn.Error)
	}

	return nil
}
