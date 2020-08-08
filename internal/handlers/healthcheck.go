package handlers

import (
	"xsolla-sdk-backend/internal/app"
	"xsolla-sdk-backend/internal/server/restapi/operations/healthcheck"

	"github.com/go-openapi/runtime/middleware"
)

const (
	OkValue     = "OK"
	FailedValue = "FAILED"
)

type HealthcheckHandler struct {
	a *app.Application
}

func NewHealthcheckHandler(a *app.Application) HealthcheckHandler {
	return HealthcheckHandler{
		a: a,
	}
}

func (h *HealthcheckHandler) CheckHealth(params healthcheck.HealthcheckParams) middleware.Responder {
	return healthcheck.NewHealthcheckOK()
}
