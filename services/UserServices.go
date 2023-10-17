package services

import (
	"pinjam_buku/models"

	"github.com/gobuffalo/validate/v3"
)

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

	password := models.PasswordHash(user.Password)
	user.Password = password
	return user, nil
}
