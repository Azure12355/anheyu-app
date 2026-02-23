/*
 * @Description: Portfolio 公开接口 Handler
 * @Author: Anheyu
 * @Date: 2025-02-23
 */
package portfolio

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/anzhiyu-c/anheyu-app/pkg/domain/model"
	portfoliosvc "github.com/anzhiyu-c/anheyu-app/pkg/service/portfolio"
)

// PublicHandler 公开接口处理器
type PublicHandler struct {
	svc portfoliosvc.Service
}

// NewPublicHandler 创建公开接口处理器
func NewPublicHandler(svc portfoliosvc.Service) *PublicHandler {
	return &PublicHandler{svc: svc}
}

// List 获取作品列表
// @Summary      获取作品列表
// @Description  获取作品列表，支持分页、类型筛选、状态筛选和关键词搜索
// @Tags         作品展示-公开接口
// @Accept       json
// @Produce      json
// @Param        page query int false "页码" default(1)
// @Param        page_size query int false "每页数量" default(12)
// @Param        project_type query string false "项目类型" Enums(frontend, vibecoding, fullstack, miniprogram, app, other)
// @Param        status query string false "项目状态" Enums(developing, completed, archived)
// @Param        keyword query string false "搜索关键词"
// @Param        featured query bool false "是否只看精选"
// @Success      200 {object} response.Response{data=model.PortfolioListResponse}
// @Router       /api/public/portfolio/list [get]
func (h *PublicHandler) List(c *gin.Context) {
	var options model.PortfolioListOptions
	if err := c.ShouldBindQuery(&options); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    400,
			"message": "请求参数无效",
			"data":    nil,
		})
		return
	}

	result, err := h.svc.List(c.Request.Context(), &options)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":    500,
			"message": "获取列表失败",
			"data":    nil,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": "success",
		"data":    result,
	})
}

// GetByID 获取作品详情
// @Summary      获取作品详情
// @Description  根据公共ID获取作品详细信息
// @Tags         作品展示-公开接口
// @Accept       json
// @Produce      json
// @Param        id path string true "作品公共ID"
// @Success      200 {object} response.Response{data=model.Portfolio}
// @Failure      404 {object} response.Response
// @Router       /api/public/portfolio/{id} [get]
func (h *PublicHandler) GetByID(c *gin.Context) {
	id := c.Param("id")

	portfolio, err := h.svc.GetByID(c.Request.Context(), id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"code":    404,
			"message": "作品不存在",
			"data":    nil,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": "success",
		"data":    portfolio,
	})
}

// GetStats 获取统计数据
// @Summary      获取作品统计
// @Description  获取作品总数、按类型统计、按状态统计和技术栈排行
// @Tags         作品展示-公开接口
// @Accept       json
// @Produce      json
// @Success      200 {object} response.Response{data=model.PortfolioStatsResponse}
// @Router       /api/public/portfolio/stats [get]
func (h *PublicHandler) GetStats(c *gin.Context) {
	stats, err := h.svc.GetStats(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":    500,
			"message": "获取统计数据失败",
			"data":    nil,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": "success",
		"data":    stats,
	})
}
