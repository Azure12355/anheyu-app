package repository

import (
	"context"

	"github.com/anzhiyu-c/anheyu-app/pkg/domain/model"
)

// PortfolioRepository 作品仓储接口
type PortfolioRepository interface {
	// Create 创建作品
	Create(ctx context.Context, req *model.CreatePortfolioRequest) (*model.Portfolio, error)

	// GetByID 根据公共ID获取作品
	GetByID(ctx context.Context, publicID string) (*model.Portfolio, error)

	// Update 更新作品
	Update(ctx context.Context, publicID string, req *model.UpdatePortfolioRequest) (*model.Portfolio, error)

	// Delete 删除作品（软删除）
	Delete(ctx context.Context, publicID string) error

	// List 获取作品列表
	List(ctx context.Context, options *model.PortfolioListOptions) ([]*model.Portfolio, int, error)

	// GetStats 获取统计数据
	GetStats(ctx context.Context) (*model.PortfolioStatsResponse, error)

	// UpdateSort 批量更新排序
	UpdateSort(ctx context.Context, sorts map[string]int) error
}
