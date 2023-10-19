package services

import (
	"pinjam_buku/models"

	"github.com/gobuffalo/validate/v3"
)

type InputAuthLogin struct {
	Username string `json:"username" form:"username"`
	Password string `json:"password" form:"password"`
}

func UserServiceStore(user *models.User) (*models.User, *validate.Errors) {
	errors := validate.NewErrors()

	if user.Email == "" {
		errors.Add("email", "Email is required.")
	}

	if user.Username == "" {
		errors.Add("username", "Username is required.")
	}

	if errors.HasAny() {
		return user, errors
	}

	theUser := models.User{}
	query := models.DB.Where("username = ? OR email = ?", user.Username, user.Email)

	error := query.First(&theUser)

	if error == nil {
		errors.Add("fail", "Username or Email is not available.")
		return user, errors
	}

	password := models.PasswordHash(user.Password)
	user.Password = password
	return user, nil
}

func UserServiceRegister(user *models.User) (*models.User, *validate.Errors) {
	errors := validate.NewErrors()

	if user.Email == "" {
		errors.Add("email", "Email is required.")
	}

	if user.Username == "" {
		errors.Add("username", "Username is required.")
	}

	if errors.HasAny() {
		return user, errors
	}

	theUser := models.User{}
	query := models.DB.Where("username = ? OR email = ?", user.Username, user.Email)

	error := query.First(&theUser)

	if error == nil {
		errors.Add("fail", "Username or Email is not available.")
		return user, errors
	}

	password := models.PasswordHash(user.Password)
	user.Password = password
	return user, nil
}

func UserServiceLogin(input *InputAuthLogin) (*InputAuthLogin, *validate.Errors) {
	errors := validate.NewErrors()

	if input.Username == "" {
		errors.Add("username", "username is required.")
	}

	if input.Password == "" {
		errors.Add("password", "password is required")
	}

	if errors.HasAny() {
		return input, errors
	}

	return input, nil
}
