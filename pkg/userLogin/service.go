package userLogin

import (
	"time"
)

// ServiceImpl implements repo for API.
type ServiceImpl struct {
	repo Repo
}

// Create new user
func (s ServiceImpl) AddNewUser(a *UserRequestBody) error {

	currentTime := time.Now()

	var newUser UserDetails
	newUser.Email = a.Email
	newUser.Password = a.Password

	newUser.CreatedAt = &currentTime

	err := s.repo.AddNewUser(&newUser)

	return err
}

func (s ServiceImpl) AuthEmail(a *UserRequestBody) error {

	var user UserDetails
	user.Email = a.Email

	err := s.repo.EmailExist(&user)

	if err != nil {
		//user does not exist
		return err
	}

	return nil
}

func (s ServiceImpl) AuthPassword(a *UserRequestBody) error {

	var user UserDetails
	user.Password = a.Password

	err := s.repo.PasswordExist(&user)

	if err != nil {
		//Password does not exist
		return err
	}

	return nil
}

// func (s ServiceImpl) ChangePassword(a *UserRequestBody) error {

// 	currentTime := time.Now()

// 	var user UserDetails
// 	user.Email = a.Email
// 	user.Password = a.Password

// 	user.LastLoggedIn = &currentTime

// 	err := s.repo.AddNewUser(&user)

// 	return err
// }

func NewService(repo Repo) *ServiceImpl {
	return &ServiceImpl{
		repo: repo,
	}
}
