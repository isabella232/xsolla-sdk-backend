package handlers

import (
	"xsolla-sdk-backend/internal/app"
	"xsolla-sdk-backend/internal/server/restapi/operations/login"
	servicelogin "xsolla-sdk-backend/services/login"

	"github.com/go-openapi/runtime/middleware"
)

type LoginHandler struct {
	a *app.Application
}

func NewLoginHandler(a *app.Application) LoginHandler {
	return LoginHandler{
		a: a,
	}
}

func (l *LoginHandler) Login(params login.LoginParams) middleware.Responder {
	email := *params.Body.Email

	loginService := servicelogin.NewLoginService(l.a)

	result, err := loginService.Login(email)

	if err != nil {
		panic(err)
	}
	return login.NewLoginOK().WithPayload(result)
}
