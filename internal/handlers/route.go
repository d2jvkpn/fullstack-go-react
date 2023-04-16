package handlers

import (
	"net/http"

	"fullstack-go-react/internal/settings"

	"github.com/gin-gonic/gin"
)

func RouteOpen(router *gin.RouterGroup, handers ...gin.HandlerFunc) {
	router.GET("/healthz", func(ctx *gin.Context) {
		ctx.AbortWithStatus(http.StatusOK)
	})

	open := router.Group("/api/v1/open", handers...)

	open.GET("/version", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"code": 0, "msg": "ok", "data": gin.H{"version": settings.GetVersion()},
		})
	})
}

func RouteAuth(router *gin.RouterGroup, handers ...gin.HandlerFunc) {
	_ = router.Group("/api/v1/auth", handers...)
}
