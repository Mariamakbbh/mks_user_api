package userLogin

import (
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type RepoImpl struct {
	DB *gorm.DB
}

// Create new user
func (r *RepoImpl) AddNewUser(a *UserDetails) error {

	var err error

	// insert new db entry
	if result := r.DB.Create(&a); result.Error != nil {
		return fiber.NewError(fiber.StatusNotFound, result.Error.Error())
	}

	return err
}

// Authorise User Email
func (r *RepoImpl) EmailExist(a *UserDetails) error {

	// find user in db by email
	findEmailResult := r.DB.Where("email = ?", a.Email).First(&a)

	if findEmailResult.Error != nil {
		return fiber.NewError(fiber.StatusNotFound, findEmailResult.Error.Error())
	}

	return nil

}

// Authorise User Password
func (r *RepoImpl) PasswordExist(a *UserDetails) error {
	// find user in db by password
	findPasswordResult := r.DB.Where("password = ?", a.Password).First(&a)

	if findPasswordResult.Error != nil {
		return fiber.NewError(fiber.StatusNotFound, findPasswordResult.Error.Error())
	}

	return nil
}

// // update user last logged in time
// addLastLoggedIN := r.DB.Model(&a).Update("last_logged_in", a.LastLoggedIn)
// if addLastLoggedIN != nil {
// 	return fiber.NewError(fiber.StatusNotFound, addLastLoggedIN.Error.Error())
// }

// return err

// //get User Password (Forgot Password)
// func (r *RepoImpl) ChangePassword(a *UserDetails) error {

// 	var err error

// 	// find user in db by email
// 	if result := r.DB.First(&a, "email = ?", a.Email); result.Error != nil {
// 		return fiber.NewError(fiber.StatusNotFound, result.Error.Error())
// 	}

// 	// update user last logged in time
// 	addLastLoggedIN := r.DB.Model(&a).Update("last_logged_in", a.LastLoggedIn)
// 	if addLastLoggedIN != nil {
// 		return fiber.NewError(fiber.StatusNotFound, addLastLoggedIN.Error.Error())
// 	}

// 	return err
// }

//list all User Details

func NewRepo(DB *gorm.DB) *RepoImpl {
	return &RepoImpl{
		DB: DB,
	}
}
