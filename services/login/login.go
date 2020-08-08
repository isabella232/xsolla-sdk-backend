package servicelogin

import (
	"fmt"

	domainlogin "xsolla-sdk-backend/domain/login"
	"xsolla-sdk-backend/internal/app"
	"xsolla-sdk-backend/internal/server/models"
)

type LoginService struct {
	a *app.Application
}

func NewLoginService(a *app.Application) LoginService {
	return LoginService{
		a: a,
	}
}

func (l *LoginService) Login(email fmt.Stringer) (*models.AccessToken, error) {
	login := domainlogin.NewLoginDomain(l.a)
	userItem, err := login.LoginUser(email)
	if err != nil {
		return nil, fmt.Errorf("failed login user. Error: %s", err)
	}

	payload := models.AccessToken{
		AccessToken: userItem.AccessToken,
	}
	return &payload, nil
}
