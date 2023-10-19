package actions

import (
	"net/http"
	"pinjam_buku/models"
	"pinjam_buku/repository"
	"pinjam_buku/services"

	"github.com/gobuffalo/buffalo"
)

type AuthResource struct{}

type AuthStruct struct {
	Token string `json:"token"`
}

func (authResource AuthResource) Login(ctx buffalo.Context) error {
	type response struct {
		Response
		Data *AuthStruct `json:"data"`
	}

	inputBody := &services.InputAuthLogin{}

	if err := ctx.Bind(inputBody); err != nil {
		return ctx.Render(http.StatusBadRequest, r.JSON(ResponseErr{
			Success: false,
			Message: "Invalid Request.",
			Errors:  err.Error(),
		}))
	}

	input, validationError := services.UserServiceLogin(inputBody)

	if validationError != nil {
		return ctx.Render(http.StatusBadRequest, r.JSON(ResponseErr{
			Success: false,
			Message: "Invalid Request.",
			Errors:  validationError.Error(),
		}))
	}

	token, err := repository.Login(input)

	if err != nil {
		return ctx.Render(http.StatusBadRequest, r.JSON(ResponseErr{
			Success: false,
			Message: "Invalid Request.",
			Errors:  err.Error(),
		}))
	}

	return ctx.Render(http.StatusCreated, r.JSON(response{
		Response: Response{
			Success: true,
			Message: "Welcome back!",
		},
		Data: &AuthStruct{
			Token: token,
		},
	}))
}

func (authResource AuthResource) Register(ctx buffalo.Context) error {
	type response struct {
		Response
		Data *AuthStruct `json:"data"`
	}

	userRaw := &models.User{}
	if err := ctx.Bind(userRaw); err != nil {
		return ctx.Render(http.StatusBadRequest, r.JSON(ResponseErr{
			Success: false,
			Message: "Failed create user",
		}))
	}

	user, errVal := services.UserServiceRegister(userRaw)

	if errVal.HasAny() {
		return ctx.Render(http.StatusBadRequest, r.JSON(ResponseErr{
			Success: false,
			Message: "Invalid request",
			Errors:  errVal.Error(),
		}))
	}

	models.DB.Store.Transaction()
	token, errors := repository.Register(user)

	if errors != nil {
		models.DB.Store.Rollback()
		return ctx.Render(http.StatusInternalServerError, r.JSON(ResponseErr{
			Success: false,
			Message: "Internal Error",
			Errors:  errors.Error(),
		}))
	}

	models.DB.Store.Commit()
	return ctx.Render(http.StatusCreated, r.JSON(response{
		Response: Response{
			Success: true,
			Message: "Success create data user.",
		},
		Data: &AuthStruct{
			Token: token,
		},
	}))
}
