package utils

import (
	"github.com/gin-gonic/gin"

	"github.com/aaantiii/lostapp/backend/env"
)

type FrontendRoute string

const (
	FERouteLoginFailed FrontendRoute = "/auth/login/failed"
	FERouteIndex       FrontendRoute = ""
)

func (route FrontendRoute) Path() string {
	return string(route)
}

func (route FrontendRoute) URL() string {
	return env.FRONTEND_URL.Value() + route.Path()
}

func (route FrontendRoute) RedirectWithStatus(c *gin.Context, status int) {
	c.Redirect(status, route.URL())
}
