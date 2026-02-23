/*
 * @Description: Portfolio Service 接口定义
 * @Author: Anheyu
 * @Date: 2025-02-23
 */
package portfolio

import (
	"context"

	"github.com/anzhiyu-c/anheyu-app/pkg/domain/model"
)

// Service 作品服务接口
type Service interface {
	// Create 创建作品
	Create(ctx context.Context, req *model.CreatePortfolioRequest) (*model.Portfolio, error)

	// GetByID 根据ID获取作品
	GetByID(ctx context.Context, publicID string) (*model.Portfolio, error)

	// Update 更新作品
	Update(ctx context.Context, publicID string, req *model.UpdatePortfolioRequest) (*model.Portfolio, error)

	// Delete 删除作品
	Delete(ctx context.Context, publicID string) error

	// List 获取作品列表
	List(ctx context.Context, options *model.PortfolioListOptions) (*model.PortfolioListResponse, error)

	// GetStats 获取统计数据
	GetStats(ctx context.Context) (*model.PortfolioStatsResponse, error)

	// UpdateSort 批量更新排序
	UpdateSort(ctx context.Context, sorts map[string]int) error
}
