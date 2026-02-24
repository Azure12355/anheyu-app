/*
 * @Description: Portfolio Service 实现
 * @Author: Anheyu
 * @Date: 2025-02-23
 */
package portfolio

import (
	"context"
	"fmt"

	"github.com/anzhiyu-c/anheyu-app/internal/pkg/parser"
	"github.com/anzhiyu-c/anheyu-app/pkg/domain/model"
	"github.com/anzhiyu-c/anheyu-app/pkg/domain/repository"
)

type serviceImpl struct {
	repo repository.PortfolioRepository
}

// NewService 创建作品服务
func NewService(repo repository.PortfolioRepository) Service {
	return &serviceImpl{repo: repo}
}

// Create 创建作品
func (s *serviceImpl) Create(ctx context.Context, req *model.CreatePortfolioRequest) (*model.Portfolio, error) {
	// 设置默认值
	if req.Status == "" {
		req.Status = "developing"
	}
	if req.ProjectType == "" {
		req.ProjectType = "other"
	}
	if req.Technologies == nil {
		req.Technologies = []string{}
	}

	portfolio, err := s.repo.Create(ctx, req)
	if err != nil {
		return nil, fmt.Errorf("failed to create portfolio: %w", err)
	}

	return portfolio, nil
}

// GetByID 根据ID获取作品
func (s *serviceImpl) GetByID(ctx context.Context, publicID string) (*model.Portfolio, error) {
	portfolio, err := s.repo.GetByID(ctx, publicID)
	if err != nil {
		return nil, fmt.Errorf("failed to get portfolio: %w", err)
	}

	// 渲染 Markdown 内容为 HTML
	if portfolio.Overview != "" {
		overviewHTML, err := parser.MarkdownToHTML(portfolio.Overview)
		if err == nil {
			portfolio.OverviewHTML = overviewHTML
		}
		// 如果渲染失败，保留原始 markdown 文本在前端处理
	}
	if portfolio.Challenge != "" {
		challengeHTML, err := parser.MarkdownToHTML(portfolio.Challenge)
		if err == nil {
			portfolio.ChallengeHTML = challengeHTML
		}
	}
	if portfolio.Solution != "" {
		solutionHTML, err := parser.MarkdownToHTML(portfolio.Solution)
		if err == nil {
			portfolio.SolutionHTML = solutionHTML
		}
	}

	return portfolio, nil
}

// Update 更新作品
func (s *serviceImpl) Update(ctx context.Context, publicID string, req *model.UpdatePortfolioRequest) (*model.Portfolio, error) {
	portfolio, err := s.repo.Update(ctx, publicID, req)
	if err != nil {
		return nil, fmt.Errorf("failed to update portfolio: %w", err)
	}

	return portfolio, nil
}

// Delete 删除作品
func (s *serviceImpl) Delete(ctx context.Context, publicID string) error {
	err := s.repo.Delete(ctx, publicID)
	if err != nil {
		return fmt.Errorf("failed to delete portfolio: %w", err)
	}

	return nil
}

// List 获取作品列表
func (s *serviceImpl) List(ctx context.Context, options *model.PortfolioListOptions) (*model.PortfolioListResponse, error) {
	// 设置默认分页
	if options.Page < 1 {
		options.Page = 1
	}
	if options.PageSize < 1 {
		options.PageSize = 12
	}

	portfolios, total, err := s.repo.List(ctx, options)
	if err != nil {
		return nil, fmt.Errorf("failed to list portfolios: %w", err)
	}

	return &model.PortfolioListResponse{
		List:  convertPortfolioList(portfolios),
		Total: total,
	}, nil
}

// GetStats 获取统计数据
func (s *serviceImpl) GetStats(ctx context.Context) (*model.PortfolioStatsResponse, error) {
	stats, err := s.repo.GetStats(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to get stats: %w", err)
	}

	return stats, nil
}

// UpdateSort 批量更新排序
func (s *serviceImpl) UpdateSort(ctx context.Context, sorts map[string]int) error {
	err := s.repo.UpdateSort(ctx, sorts)
	if err != nil {
		return fmt.Errorf("failed to update sort: %w", err)
	}

	return nil
}

// convertPortfolioList 转换作品列表
func convertPortfolioList(portfolios []*model.Portfolio) []model.Portfolio {
	result := make([]model.Portfolio, len(portfolios))
	for i, p := range portfolios {
		result[i] = *p
	}
	return result
}
