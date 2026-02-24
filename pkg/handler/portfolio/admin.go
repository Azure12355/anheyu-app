/*
 * @Description: Portfolio 管理接口 Handler
 * @Author: Anheyu
 * @Date: 2025-02-23
 */
package portfolio

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/anzhiyu-c/anheyu-app/ent"
	"github.com/anzhiyu-c/anheyu-app/pkg/domain/model"
	portfoliosvc "github.com/anzhiyu-c/anheyu-app/pkg/service/portfolio"
)

// AdminHandler 管理接口处理器
type AdminHandler struct {
	svc portfoliosvc.Service
}

// NewAdminHandler 创建管理接口处理器
func NewAdminHandler(svc portfoliosvc.Service) *AdminHandler {
	return &AdminHandler{svc: svc}
}

// Create 创建作品
// @Summary      创建作品
// @Description  创建新的作品展示项目
// @Tags         作品展示-管理接口
// @Security     BearerAuth
// @Accept       json
// @Produce      json
// @Param        portfolio body model.CreatePortfolioRequest true "作品信息"
// @Success      200 {object} response.Response{data=model.Portfolio}
// @Router       /api/portfolio [post]
func (h *AdminHandler) Create(c *gin.Context) {
	var req model.CreatePortfolioRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    400,
			"message": "请求参数无效",
			"data":    nil,
		})
		return
	}

	portfolio, err := h.svc.Create(c.Request.Context(), &req)
	if err != nil {
		// 判断是否为约束冲突错误
		if ent.IsConstraintError(err) {
			c.JSON(http.StatusConflict, gin.H{
				"code":    409,
				"message": "数据冲突，可能存在重复记录",
				"data":    nil,
			})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":    500,
			"message": "创建失败: " + err.Error(),
			"data":    nil,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": "创建成功",
		"data":    portfolio,
	})
}

// Update 更新作品
// @Summary      更新作品
// @Description  更新作品信息
// @Tags         作品展示-管理接口
// @Security     BearerAuth
// @Accept       json
// @Produce      json
// @Param        id path string true "作品公共ID"
// @Param        portfolio body model.UpdatePortfolioRequest true "作品信息"
// @Success      200 {object} response.Response{data=model.Portfolio}
// @Router       /api/portfolio/{id} [put]
func (h *AdminHandler) Update(c *gin.Context) {
	id := c.Param("id")
	var req model.UpdatePortfolioRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    400,
			"message": "请求参数无效",
			"data":    nil,
		})
		return
	}

	portfolio, err := h.svc.Update(c.Request.Context(), id, &req)
	if err != nil {
		// 判断是否为 NotFound 错误
		if ent.IsNotFound(err) {
			c.JSON(http.StatusNotFound, gin.H{
				"code":    404,
				"message": "作品不存在",
				"data":    nil,
			})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":    500,
			"message": "更新失败: " + err.Error(),
			"data":    nil,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": "更新成功",
		"data":    portfolio,
	})
}

// Delete 删除作品
// @Summary      删除作品
// @Description  删除作品（软删除）
// @Tags         作品展示-管理接口
// @Security     BearerAuth
// @Accept       json
// @Produce      json
// @Param        id path string true "作品公共ID"
// @Success      200 {object} response.Response
// @Router       /api/portfolio/{id} [delete]
func (h *AdminHandler) Delete(c *gin.Context) {
	id := c.Param("id")

	err := h.svc.Delete(c.Request.Context(), id)
	if err != nil {
		// 判断是否为 NotFound 错误
		if ent.IsNotFound(err) {
			c.JSON(http.StatusNotFound, gin.H{
				"code":    404,
				"message": "作品不存在",
				"data":    nil,
			})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":    500,
			"message": "删除失败: " + err.Error(),
			"data":    nil,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": "删除成功",
		"data":    nil,
	})
}

// UpdateSort 批量更新排序
// @Summary      批量更新排序
// @Description  批量更新作品的排序权重
// @Tags         作品展示-管理接口
// @Security     BearerAuth
// @Accept       json
// @Produce      json
// @Param        sorts body map[string]int true "排序映射 {public_id: sort_order}"
// @Success      200 {object} response.Response
// @Router       /api/portfolio/sort [put]
func (h *AdminHandler) UpdateSort(c *gin.Context) {
	var sorts map[string]int
	if err := c.ShouldBindJSON(&sorts); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    400,
			"message": "请求参数无效",
			"data":    nil,
		})
		return
	}

	// 限制批量更新的数量，防止恶意请求
	if len(sorts) > 100 {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    400,
			"message": "单次更新数量不能超过100个",
			"data":    nil,
		})
		return
	}

	err := h.svc.UpdateSort(c.Request.Context(), sorts)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":    500,
			"message": "更新排序失败: " + err.Error(),
			"data":    nil,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": "更新成功",
		"data":    nil,
	})
}
