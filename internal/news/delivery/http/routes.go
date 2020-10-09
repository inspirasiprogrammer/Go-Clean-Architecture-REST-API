package http

import (
	"github.com/AleksK1NG/api-mc/config"
	"github.com/AleksK1NG/api-mc/internal/auth"
	"github.com/AleksK1NG/api-mc/internal/middleware"
	"github.com/AleksK1NG/api-mc/internal/news"
	"github.com/AleksK1NG/api-mc/pkg/logger"
	"github.com/labstack/echo"
)

// Map news routes
func MapNewsRoutes(newsGroup *echo.Group, h news.Handlers, authUC auth.UseCase, cfg *config.Config, logger *logger.Logger) {
	newsGroup.POST("/create", h.Create(), middleware.AuthJWTMiddleware(authUC, cfg, logger))
	newsGroup.PUT("/:news_id", h.Update(), middleware.AuthJWTMiddleware(authUC, cfg, logger))
	newsGroup.DELETE("/:news_id", h.Delete(), middleware.AuthJWTMiddleware(authUC, cfg, logger))
	newsGroup.GET("/:news_id", h.GetByID())
	newsGroup.GET("/search", h.SearchByTitle())
	newsGroup.GET("", h.GetNews())
}
