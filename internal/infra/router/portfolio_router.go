/*
 * @Description: Portfolio 路由注册
 * @Author: Anheyu
 * @Date: 2025-02-23
 */
package router

import (
	"github.com/gin-gonic/gin"
	"github.com/anzhiyu-c/anheyu-app/pkg/handler/portfolio"
)

// RegisterPortfolioRoutes 注册作品展示路由
func RegisterPortfolioRoutes(r *gin.Engine, publicHandler *portfolio.PublicHandler, adminHandler *portfolio.AdminHandler) {
	// 公开接口
	public := r.Group("/api/public/portfolio")
	{
		public.GET("/list", publicHandler.List)
		public.GET("/stats", publicHandler.GetStats)
		public.GET("/:id", publicHandler.GetByID)
	}

	// 管理接口（需要认证）
	admin := r.Group("/api/portfolio")
	// admin.Use(middleware.Auth()) // TODO: 添加认证中间件
	{
		admin.POST("", adminHandler.Create)
		admin.PUT("/:id", adminHandler.Update)
		admin.DELETE("/:id", adminHandler.Delete)
		admin.PUT("/sort", adminHandler.UpdateSort)
	}
}
