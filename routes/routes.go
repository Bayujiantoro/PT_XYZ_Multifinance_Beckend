package routes

import "github.com/labstack/echo/v4"

func RouterInit(e *echo.Group) {
	AuthRoutes(e)
	UserRouter(e)
	TransactionRouter(e)
	PinjamanRoutes(e)
}