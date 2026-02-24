/*
 * @Description: Portfolio Repository Ent 实现
 * @Author: Anheyu
 * @Date: 2025-02-23
 */
package ent

import (
	"context"
	"fmt"
	"sort"

	"github.com/anzhiyu-c/anheyu-app/ent"
	"github.com/anzhiyu-c/anheyu-app/ent/portfolio"
	"github.com/anzhiyu-c/anheyu-app/ent/portfoliotechnology"
	"github.com/anzhiyu-c/anheyu-app/pkg/domain/model"
	"github.com/anzhiyu-c/anheyu-app/pkg/domain/repository"
	"github.com/anzhiyu-c/anheyu-app/pkg/idgen"
)

type portfolioRepo struct {
	client *ent.Client
}

// NewPortfolioRepo 创建作品仓储
func NewPortfolioRepo(client *ent.Client) repository.PortfolioRepository {
	return &portfolioRepo{client: client}
}

// toModel 转换为领域模型
func (r *portfolioRepo) toModel(p *ent.Portfolio) *model.Portfolio {
	publicID, _ := idgen.GeneratePublicID(p.ID, idgen.EntityTypePortfolio)

	// 加载技术栈关联数据
	technologies := make([]string, 0, len(p.Edges.Technologies))
	for _, tech := range p.Edges.Technologies {
		technologies = append(technologies, tech.Technology)
	}

	return &model.Portfolio{
		ID:            publicID,
		Title:         p.Title,
		Description:   p.Description,
		CoverURL:      p.CoverURL,
		ProjectType:   string(p.ProjectType),
		Status:        string(p.Status),
		Tier:          string(p.Tier),
		Technologies:  technologies,
		DemoURL:       p.DemoURL,
		GithubURL:     p.GithubURL,
		Featured:      p.Featured,
		SortOrder:     int(p.SortOrder),
		Overview:      p.Overview,
		Role:          p.Role,
		Duration:      p.Duration,
		Client:        p.Client,
		Challenge:     p.Challenge,
		Solution:      p.Solution,
		GalleryImages: p.GalleryImages,
		CreatedAt:     p.CreatedAt,
		UpdatedAt:     p.UpdatedAt,
	}
}

// Create 创建作品
func (r *portfolioRepo) Create(ctx context.Context, req *model.CreatePortfolioRequest) (*model.Portfolio, error) {
	// 使用事务
	tx, err := r.client.Tx(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to begin transaction: %w", err)
	}

	// 确保事务在发生错误时回滚
	defer func() {
		if err != nil {
			tx.Rollback()
		}
	}()

	// 创建作品
	p, err := tx.Portfolio.Create().
		SetTitle(req.Title).
		SetDescription(req.Description).
		SetCoverURL(req.CoverURL).
		SetProjectType(portfolio.ProjectType(req.ProjectType)).
		SetStatus(portfolio.Status(req.Status)).
		SetTier(portfolio.Tier(req.Tier)).
		SetDemoURL(req.DemoURL).
		SetGithubURL(req.GithubURL).
		SetFeatured(req.Featured).
		SetSortOrder(req.SortOrder).
		SetOverview(req.Overview).
		SetRole(req.Role).
		SetDuration(req.Duration).
		SetClient(req.Client).
		SetChallenge(req.Challenge).
		SetSolution(req.Solution).
		SetGalleryImages(req.GalleryImages).
		Save(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to create portfolio: %w", err)
	}

	// 创建技术栈关联（去重处理）
	techSet := make(map[string]bool)
	for _, tech := range req.Technologies {
		// 跳过空字符串
		if tech == "" {
			continue
		}
		// 去重：同一作品内避免重复技术
		if techSet[tech] {
			continue
		}
		techSet[tech] = true

		_, err = tx.PortfolioTechnology.Create().
			SetPortfolioID(p.ID).
			SetTechnology(tech).
			Save(ctx)
		if err != nil {
			return nil, fmt.Errorf("failed to create technology '%s': %w", tech, err)
		}
	}

	if err := tx.Commit(); err != nil {
		return nil, fmt.Errorf("failed to commit transaction: %w", err)
	}

	// 重新查询带关联的作品
	p, err = r.client.Portfolio.Query().
		Where(portfolio.ID(p.ID)).
		WithTechnologies(func(q *ent.PortfolioTechnologyQuery) {
			// 按技术名称排序，保证结果稳定
			q.Order(ent.Asc(portfoliotechnology.FieldTechnology))
		}).
		Only(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to query portfolio: %w", err)
	}

	return r.toModel(p), nil
}

// GetByID 根据公共ID获取作品
func (r *portfolioRepo) GetByID(ctx context.Context, publicID string) (*model.Portfolio, error) {
	id, _, err := idgen.DecodePublicID(publicID)
	if err != nil {
		return nil, fmt.Errorf("invalid public ID: %w", err)
	}

	p, err := r.client.Portfolio.Query().
		Where(portfolio.ID(id)).
		WithTechnologies(func(q *ent.PortfolioTechnologyQuery) {
			// 按技术名称排序，保证结果稳定
			q.Order(ent.Asc(portfoliotechnology.FieldTechnology))
		}).
		Only(ctx)
	if err != nil {
		return nil, fmt.Errorf("portfolio not found: %w", err)
	}

	return r.toModel(p), nil
}

// Update 更新作品
func (r *portfolioRepo) Update(ctx context.Context, publicID string, req *model.UpdatePortfolioRequest) (*model.Portfolio, error) {
	id, _, err := idgen.DecodePublicID(publicID)
	if err != nil {
		return nil, fmt.Errorf("invalid public ID: %w", err)
	}

	// 使用事务
	tx, err := r.client.Tx(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to begin transaction: %w", err)
	}

	// 确保事务在发生错误时回滚
	defer func() {
		if err != nil {
			tx.Rollback()
		}
	}()

	// 构建更新器 - 使用 UpdateOneID 以便获取更新后的对象
	update := tx.Portfolio.UpdateOneID(id)

	if req.Title != nil {
		update.SetTitle(*req.Title)
	}
	if req.Description != nil {
		update.SetDescription(*req.Description)
	}
	if req.CoverURL != nil {
		update.SetCoverURL(*req.CoverURL)
	}
	if req.ProjectType != nil {
		update.SetProjectType(portfolio.ProjectType(*req.ProjectType))
	}
	if req.Status != nil {
		update.SetStatus(portfolio.Status(*req.Status))
	}
	if req.Tier != nil {
		update.SetTier(portfolio.Tier(*req.Tier))
	}
	if req.DemoURL != nil {
		update.SetDemoURL(*req.DemoURL)
	}
	if req.GithubURL != nil {
		update.SetGithubURL(*req.GithubURL)
	}
	if req.Featured != nil {
		update.SetFeatured(*req.Featured)
	}
	if req.SortOrder != nil {
		update.SetSortOrder(*req.SortOrder)
	}
	if req.Overview != nil {
		update.SetOverview(*req.Overview)
	}
	if req.Role != nil {
		update.SetRole(*req.Role)
	}
	if req.Duration != nil {
		update.SetDuration(*req.Duration)
	}
	if req.Client != nil {
		update.SetClient(*req.Client)
	}
	if req.Challenge != nil {
		update.SetChallenge(*req.Challenge)
	}
	if req.Solution != nil {
		update.SetSolution(*req.Solution)
	}
	if req.GalleryImages != nil {
		update.SetGalleryImages(*req.GalleryImages)
	}

	p, err := update.Save(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to update portfolio: %w", err)
	}

	// 更新技术栈
	if req.Technologies != nil {
		// 先查询当前的技术栈关联
		currentTechs, err := tx.Portfolio.Query().
			Where(portfolio.ID(id)).
			QueryTechnologies().
			All(ctx)
		if err != nil {
			return nil, fmt.Errorf("failed to query current technologies: %w", err)
		}

		// 删除旧的关联
		for _, tech := range currentTechs {
			err = tx.PortfolioTechnology.DeleteOne(tech).Exec(ctx)
			if err != nil {
				return nil, fmt.Errorf("failed to delete old technology: %w", err)
			}
		}

		// 创建新的关联（去重处理）
		techSet := make(map[string]bool)
		for _, techName := range *req.Technologies {
			// 跳过空字符串
			if techName == "" {
				continue
			}
			// 去重：同一作品内避免重复技术
			if techSet[techName] {
				continue
			}
			techSet[techName] = true

			_, err = tx.PortfolioTechnology.Create().
				SetPortfolioID(p.ID).
				SetTechnology(techName).
				Save(ctx)
			if err != nil {
				return nil, fmt.Errorf("failed to create technology '%s': %w", techName, err)
			}
		}
	}

	if err := tx.Commit(); err != nil {
		return nil, fmt.Errorf("failed to commit transaction: %w", err)
	}

	// 重新查询
	p, err = r.client.Portfolio.Query().
		Where(portfolio.ID(p.ID)).
		WithTechnologies(func(q *ent.PortfolioTechnologyQuery) {
			// 按技术名称排序，保证结果稳定
			q.Order(ent.Asc(portfoliotechnology.FieldTechnology))
		}).
		Only(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to query portfolio: %w", err)
	}

	return r.toModel(p), nil
}

// Delete 删除作品（软删除）
func (r *portfolioRepo) Delete(ctx context.Context, publicID string) error {
	id, _, err := idgen.DecodePublicID(publicID)
	if err != nil {
		return fmt.Errorf("invalid public ID: %w", err)
	}

	_, err = r.client.Portfolio.Delete().
		Where(portfolio.ID(id)).
		Exec(ctx)
	if err != nil {
		return fmt.Errorf("failed to delete portfolio: %w", err)
	}

	return nil
}

// List 获取作品列表
func (r *portfolioRepo) List(ctx context.Context, options *model.PortfolioListOptions) ([]*model.Portfolio, int, error) {
	// 先构建筛选查询（不加载 technologies，用于计数）
	countQuery := r.client.Portfolio.Query()

	// 筛选条件
	if options.ProjectType != "" {
		countQuery = countQuery.Where(portfolio.ProjectTypeEQ(portfolio.ProjectType(options.ProjectType)))
	}
	if options.Status != "" {
		countQuery = countQuery.Where(portfolio.StatusEQ(portfolio.Status(options.Status)))
	}
	if options.Keyword != "" {
		countQuery = countQuery.Where(portfolio.TitleContains(options.Keyword))
	}
	if options.Featured != nil {
		countQuery = countQuery.Where(portfolio.Featured(*options.Featured))
	}
	if options.Tier != "" {
		countQuery = countQuery.Where(portfolio.TierEQ(portfolio.Tier(options.Tier)))
	}

	// 获取总数
	total, err := countQuery.Count(ctx)
	if err != nil {
		return nil, 0, fmt.Errorf("failed to count portfolios: %w", err)
	}

	// 分页和排序
	page := options.Page
	if page < 1 {
		page = 1
	}
	pageSize := options.PageSize
	if pageSize < 1 {
		pageSize = 12
	}
	// 限制最大每页数量，防止恶意请求
	if pageSize > 100 {
		pageSize = 100
	}

	// 构建数据查询（加载 technologies）
	query := r.client.Portfolio.Query().WithTechnologies(func(q *ent.PortfolioTechnologyQuery) {
		// 按技术名称排序，保证结果稳定
		q.Order(ent.Asc(portfoliotechnology.FieldTechnology))
	})

	// 应用相同的筛选条件
	if options.ProjectType != "" {
		query = query.Where(portfolio.ProjectTypeEQ(portfolio.ProjectType(options.ProjectType)))
	}
	if options.Status != "" {
		query = query.Where(portfolio.StatusEQ(portfolio.Status(options.Status)))
	}
	if options.Keyword != "" {
		query = query.Where(portfolio.TitleContains(options.Keyword))
	}
	if options.Featured != nil {
		query = query.Where(portfolio.Featured(*options.Featured))
	}
	if options.Tier != "" {
		query = query.Where(portfolio.TierEQ(portfolio.Tier(options.Tier)))
	}

	portfolios, err := query.
		Order(ent.Asc(portfolio.FieldSortOrder)).
		Order(ent.Desc(portfolio.FieldCreatedAt)).
		Offset((page - 1) * pageSize).
		Limit(pageSize).
		All(ctx)
	if err != nil {
		return nil, 0, fmt.Errorf("failed to query portfolios: %w", err)
	}

	result := make([]*model.Portfolio, len(portfolios))
	for i, p := range portfolios {
		result[i] = r.toModel(p)
	}

	return result, total, nil
}

// GetStats 获取统计数据
func (r *portfolioRepo) GetStats(ctx context.Context) (*model.PortfolioStatsResponse, error) {
	portfolios, err := r.client.Portfolio.Query().
		WithTechnologies().
		All(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to query portfolios: %w", err)
	}

	stats := &model.PortfolioStatsResponse{
		Total:           len(portfolios),
		ByType:          make(map[string]int),
		ByStatus:        make(map[string]int),
		TopTechnologies: make([]model.TechnologyCount, 0),
	}

	techCount := make(map[string]int)

	for _, p := range portfolios {
		stats.ByType[string(p.ProjectType)]++
		stats.ByStatus[string(p.Status)]++

		for _, tech := range p.Edges.Technologies {
			techCount[tech.Technology]++
		}
	}

	// 技术栈 TOP 10 - 按使用次数降序排列
	for tech, count := range techCount {
		stats.TopTechnologies = append(stats.TopTechnologies, model.TechnologyCount{
			Name:  tech,
			Count: count,
		})
	}

	// 按使用次数降序排序
	sort.Slice(stats.TopTechnologies, func(i, j int) bool {
		if stats.TopTechnologies[i].Count == stats.TopTechnologies[j].Count {
			// 次数相同时按名称排序，保证结果稳定
			return stats.TopTechnologies[i].Name < stats.TopTechnologies[j].Name
		}
		return stats.TopTechnologies[i].Count > stats.TopTechnologies[j].Count
	})

	// 只返回 TOP 10
	if len(stats.TopTechnologies) > 10 {
		stats.TopTechnologies = stats.TopTechnologies[:10]
	}

	return stats, nil
}

// UpdateSort 批量更新排序
func (r *portfolioRepo) UpdateSort(ctx context.Context, sorts map[string]int) error {
	for publicID, sortOrder := range sorts {
		id, _, err := idgen.DecodePublicID(publicID)
		if err != nil {
			return fmt.Errorf("invalid public ID %s: %w", publicID, err)
		}

		_, err = r.client.Portfolio.UpdateOneID(id).
			SetSortOrder(sortOrder).
			Save(ctx)
		if err != nil {
			return fmt.Errorf("failed to update sort order for %s: %w", publicID, err)
		}
	}
	return nil
}
