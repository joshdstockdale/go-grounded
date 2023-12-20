package handler

import (
	"joshdstockdale/grounded/model"
	"joshdstockdale/grounded/view/page"

	"github.com/labstack/echo/v4"
)

type PageHandler struct{}

func (h PageHandler) HandlePageShow(c echo.Context) error {
	user := model.User{Email: "josh@me.com"}
	return render(c, page.Index(user))
} 