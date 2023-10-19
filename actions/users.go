package actions

import (
	"net/http"
	"pinjam_buku/models"
	"pinjam_buku/services"

	"github.com/gobuffalo/buffalo"
)

func UserStore(ctx buffalo.Context) error {
	type response struct {
		Response
		Data *models.User `json:"data"`
	}

	userRaw := &models.User{}
	if err := ctx.Bind(userRaw); err != nil {
		return ctx.Render(http.StatusBadRequest, r.JSON(ResponseErr{
			Success: false,
			Message: "Failed create user",
		}))
	}

	user, errVal := services.UserServiceStore(userRaw)

	if errVal.HasAny() {
		return ctx.Render(http.StatusBadRequest, r.JSON(ResponseErr{
			Success: false,
			Message: "Invalid request",
			Errors:  errVal.Error(),
		}))
	}

	models.DB.Store.Transaction()
	errors := models.DB.Create(user)

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
		Data: user,
	}))
}

func UserGet(ctx buffalo.Context) error {
	type response struct {
		Response
		Data []models.User `json:"data"`
	}

	users := []models.User{}

	err := models.DB.All(&users)

	return ctx.Render(http.StatusOK, r.JSON(response{
		Response: Response{
			Success: err != nil,
			Message: "Success getting data.",
		},
		Data: users,
	}))
}
