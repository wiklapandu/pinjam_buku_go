package repository

import (
	"pinjam_buku/models"
	"pinjam_buku/services"

	"github.com/gobuffalo/envy"
	"github.com/gobuffalo/validate/v3"
	"github.com/golang-jwt/jwt/v5"
)

type JWTClaimsUser struct {
	Email    string `json:"email"`
	Username string `json:"username"`
	*jwt.RegisteredClaims
}

func Login(input *services.InputAuthLogin) (string, *validate.Errors) {
	envSign := envy.Get("JWT_SECRET", "H3110w0rld")
	errorVal := validate.NewErrors()

	user := models.User{}
	query := models.DB.Where("username = ? OR email = ?", input.Username, input.Username)

	error := query.First(&user)

	if error != nil {
		errorVal.Add("fail", "Account is not register")
		return "", errorVal
	}

	if !models.PasswordMatch(user.Password, input.Password) {
		errorVal.Add("fail", "Invalid credentials")
		return "", errorVal
	}

	if errorVal.HasAny() {
		return "", errorVal
	}

	signatures := []byte(envSign)

	claims := JWTClaimsUser{
		Username: user.Username,
		Email:    user.Email,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	parse, err := token.SignedString(signatures)

	if err != nil {
		errorVal.Add("fail", "Error parse token.")
		return "Error parse token.", errorVal
	}

	return parse, nil
}

func Register(user *models.User) (string, *validate.Errors) {
	envSign := envy.Get("JWT_SECRET", "H3110w0rld")
	errorVal := validate.NewErrors()
	error := models.DB.Create(user)

	if error != nil {
		errorVal.Add("fail", "Failed register.")
		return "", errorVal
	}

	signatures := []byte(envSign)

	claims := JWTClaimsUser{
		Username: user.Username,
		Email:    user.Email,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	parse, err := token.SignedString(signatures)

	if err != nil {
		errorVal.Add("fail", "Error parse token.")
		return "Error parse token.", errorVal
	}

	return parse, nil
}
