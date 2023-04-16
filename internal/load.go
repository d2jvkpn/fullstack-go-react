package internal

import (
	// "fmt"
	"net/http"
	"strings"

	"fullstack-go-react/internal/handlers"
	"fullstack-go-react/internal/settings"

	"github.com/d2jvkpn/go-web/pkg/wrap"
	"github.com/d2jvkpn/x-ai/pkg/chatgpt"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func Load(config string, release bool) (err error) {
	var (
		engine *gin.Engine
		router *gin.RouterGroup
	)

	//
	if settings.GPTCli, err = chatgpt.NewClient(config, "chatgpt"); err != nil {
		return err
	}

	//
	if settings.Tls, err = settings.NewTlsConfig(config, "tls"); err != nil {
		return err
	}

	//
	level := zap.DebugLevel
	if release {
		level = zap.InfoLevel
	}
	settings.Logger = wrap.NewLogger("logs/fullstack-go-react.log", level, 256, nil)
	settings.SetupLoggers()

	//
	if release {
		gin.SetMode(gin.ReleaseMode)
		engine = gin.New()
		engine.Use(gin.Recovery())
	} else {
		engine = gin.Default()
	}
	engine.Use(cors(settings.GetCorsOrigins()))

	router = &engine.RouterGroup
	router.Static("/site", "./build")

	handlers.RouteOpen(router)
	handlers.RouteAuth(router) // TODO: auth

	_Server.Handler = engine

	return nil
}

func cors(origin string) gin.HandlerFunc {
	allowHeaders := strings.Join([]string{"Content-Type", "Authorization"}, ", ")

	exposeHeaders := strings.Join([]string{
		"Access-Control-Allow-Origin",
		"Access-Control-Allow-Headers",
		"Content-Type",
		"Content-Length",
	}, ", ")

	return func(ctx *gin.Context) {
		ctx.Header("Access-Control-Allow-Origin", origin)

		ctx.Header("Access-Control-Allow-Headers", allowHeaders)
		// Content-Type, Authorization, X-CSRF-Token
		ctx.Header("Access-Control-Expose-Headers", exposeHeaders)
		ctx.Header("Access-Control-Allow-Credentials", "true")
		ctx.Header("Access-Control-Allow-Methods", "GET, POST, OPTIONS, HEAD")

		if ctx.Request.Method == "OPTIONS" {
			ctx.AbortWithStatus(http.StatusNoContent)
			return
		}
		ctx.Next()
	}
}
