/*
 * @Description: Portfolio Service 单元测试
 * @Author: Anheyu
 * @Date: 2025-02-24
 */
package portfolio

import (
	"context"
	"errors"
	"testing"

	"github.com/anzhiyu-c/anheyu-app/pkg/domain/model"
)

// mockPortfolioRepository 模拟 Portfolio Repository
type mockPortfolioRepository struct {
	createFunc  func(ctx context.Context, req *model.CreatePortfolioRequest) (*model.Portfolio, error)
	getByIDFunc func(ctx context.Context, publicID string) (*model.Portfolio, error)
	updateFunc  func(ctx context.Context, publicID string, req *model.UpdatePortfolioRequest) (*model.Portfolio, error)
	deleteFunc  func(ctx context.Context, publicID string) error
	listFunc    func(ctx context.Context, options *model.PortfolioListOptions) ([]*model.Portfolio, int, error)
	getStatsFunc func(ctx context.Context) (*model.PortfolioStatsResponse, error)
	updateSortFunc func(ctx context.Context, sorts map[string]int) error
}

func (m *mockPortfolioRepository) Create(ctx context.Context, req *model.CreatePortfolioRequest) (*model.Portfolio, error) {
	if m.createFunc != nil {
		return m.createFunc(ctx, req)
	}
	return &model.Portfolio{ID: "test-id"}, nil
}

func (m *mockPortfolioRepository) GetByID(ctx context.Context, publicID string) (*model.Portfolio, error) {
	if m.getByIDFunc != nil {
		return m.getByIDFunc(ctx, publicID)
	}
	return &model.Portfolio{ID: publicID}, nil
}

func (m *mockPortfolioRepository) Update(ctx context.Context, publicID string, req *model.UpdatePortfolioRequest) (*model.Portfolio, error) {
	if m.updateFunc != nil {
		return m.updateFunc(ctx, publicID, req)
	}
	return &model.Portfolio{ID: publicID}, nil
}

func (m *mockPortfolioRepository) Delete(ctx context.Context, publicID string) error {
	if m.deleteFunc != nil {
		return m.deleteFunc(ctx, publicID)
	}
	return nil
}

func (m *mockPortfolioRepository) List(ctx context.Context, options *model.PortfolioListOptions) ([]*model.Portfolio, int, error) {
	if m.listFunc != nil {
		return m.listFunc(ctx, options)
	}
	return []*model.Portfolio{}, 0, nil
}

func (m *mockPortfolioRepository) GetStats(ctx context.Context) (*model.PortfolioStatsResponse, error) {
	if m.getStatsFunc != nil {
		return m.getStatsFunc(ctx)
	}
	return &model.PortfolioStatsResponse{}, nil
}

func (m *mockPortfolioRepository) UpdateSort(ctx context.Context, sorts map[string]int) error {
	if m.updateSortFunc != nil {
		return m.updateSortFunc(ctx, sorts)
	}
	return nil
}

// TestPortfolioService_Create 测试创建作品
func TestPortfolioService_Create(t *testing.T) {
	ctx := context.Background()

	tests := []struct {
		name    string
		req     *model.CreatePortfolioRequest
		setup   func(*mockPortfolioRepository)
		wantErr bool
	}{
		{
			name: "成功创建作品 - 带完整参数",
			req: &model.CreatePortfolioRequest{
				Title:       "测试项目",
				Description: "测试描述",
				ProjectType: "frontend",
				Status:      "completed",
				Tier:        "featured",
			},
			setup: func(m *mockPortfolioRepository) {
				m.createFunc = func(ctx context.Context, req *model.CreatePortfolioRequest) (*model.Portfolio, error) {
					return &model.Portfolio{
						ID:          "test-id",
						Title:       req.Title,
						Description: req.Description,
						ProjectType: req.ProjectType,
						Status:      req.Status,
						Tier:        req.Tier,
					}, nil
				}
			},
			wantErr: false,
		},
		{
			name: "成功创建作品 - 使用默认值",
			req: &model.CreatePortfolioRequest{
				Title:       "测试项目",
				Description: "测试描述",
				ProjectType: "",
				Status:      "",
				Tier:        "",
			},
			setup: func(m *mockPortfolioRepository) {
				m.createFunc = func(ctx context.Context, req *model.CreatePortfolioRequest) (*model.Portfolio, error) {
					// 验证默认值已被设置
					if req.Status != "developing" {
						t.Errorf("期望默认状态为 developing，实际为 %s", req.Status)
					}
					if req.ProjectType != "other" {
						t.Errorf("期望默认项目类型为 other，实际为 %s", req.ProjectType)
					}
					if req.Tier != "normal" {
						t.Errorf("期望默认层级为 normal，实际为 %s", req.Tier)
					}
					return &model.Portfolio{ID: "test-id"}, nil
				}
			},
			wantErr: false,
		},
		{
			name: "创建失败 - 仓库错误",
			req: &model.CreatePortfolioRequest{
				Title:       "测试项目",
				Description: "测试描述",
			},
			setup: func(m *mockPortfolioRepository) {
				m.createFunc = func(ctx context.Context, req *model.CreatePortfolioRequest) (*model.Portfolio, error) {
					return nil, errors.New("database error")
				}
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mock := &mockPortfolioRepository{}
			if tt.setup != nil {
				tt.setup(mock)
			}
			svc := NewService(mock)

			got, err := svc.Create(ctx, tt.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("Create() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !tt.wantErr && got == nil {
				t.Error("Create() 返回 nil，期望非 nil")
			}
		})
	}
}

// TestPortfolioService_List 测试获取作品列表
func TestPortfolioService_List(t *testing.T) {
	ctx := context.Background()

	// 模拟数据
	mockPortfolios := []*model.Portfolio{
		{ID: "1", Title: "项目 A", Tier: "featured"},
		{ID: "2", Title: "项目 B", Tier: "recommended"},
		{ID: "3", Title: "项目 C", Tier: "normal"},
	}

	tests := []struct {
		name       string
		options    *model.PortfolioListOptions
		setup      func(*mockPortfolioRepository)
		wantCount  int
		wantTotal  int
		wantErr    bool
		verifyTier string // 如果非空，验证所有返回项的 tier
	}{
		{
			name: "成功获取列表 - 全部",
			options: &model.PortfolioListOptions{
				Page:     1,
				PageSize: 10,
			},
			setup: func(m *mockPortfolioRepository) {
				m.listFunc = func(ctx context.Context, options *model.PortfolioListOptions) ([]*model.Portfolio, int, error) {
					return mockPortfolios, len(mockPortfolios), nil
				}
			},
			wantCount: 3,
			wantTotal: 3,
			wantErr:   false,
		},
		{
			name: "按 tier 筛选 - featured",
			options: &model.PortfolioListOptions{
				Page:     1,
				PageSize: 10,
				Tier:     "featured",
			},
			setup: func(m *mockPortfolioRepository) {
				m.listFunc = func(ctx context.Context, options *model.PortfolioListOptions) ([]*model.Portfolio, int, error) {
					// 验证 tier 参数被正确传递
					if options.Tier != "featured" {
						t.Errorf("期望 Tier = featured，实际为 %s", options.Tier)
					}
					// 返回筛选结果
					filtered := []*model.Portfolio{{ID: "1", Title: "项目 A", Tier: "featured"}}
					return filtered, 1, nil
				}
			},
			wantCount:  1,
			wantTotal:  1,
			wantErr:    false,
			verifyTier: "featured",
		},
		{
			name: "按 tier 筛选 - recommended",
			options: &model.PortfolioListOptions{
				Page:     1,
				PageSize: 10,
				Tier:     "recommended",
			},
			setup: func(m *mockPortfolioRepository) {
				m.listFunc = func(ctx context.Context, options *model.PortfolioListOptions) ([]*model.Portfolio, int, error) {
					if options.Tier != "recommended" {
						t.Errorf("期望 Tier = recommended，实际为 %s", options.Tier)
					}
					filtered := []*model.Portfolio{{ID: "2", Title: "项目 B", Tier: "recommended"}}
					return filtered, 1, nil
				}
			},
			wantCount:  1,
			wantTotal:  1,
			wantErr:    false,
			verifyTier: "recommended",
		},
		{
			name: "按 tier 筛选 - normal",
			options: &model.PortfolioListOptions{
				Page:     1,
				PageSize: 10,
				Tier:     "normal",
			},
			setup: func(m *mockPortfolioRepository) {
				m.listFunc = func(ctx context.Context, options *model.PortfolioListOptions) ([]*model.Portfolio, int, error) {
					if options.Tier != "normal" {
						t.Errorf("期望 Tier = normal，实际为 %s", options.Tier)
					}
					filtered := []*model.Portfolio{{ID: "3", Title: "项目 C", Tier: "normal"}}
					return filtered, 1, nil
				}
			},
			wantCount:  1,
			wantTotal:  1,
			wantErr:    false,
			verifyTier: "normal",
		},
		{
			name: "默认分页参数 - Page < 1",
			options: &model.PortfolioListOptions{
				Page:     0,
				PageSize: 10,
			},
			setup: func(m *mockPortfolioRepository) {
				m.listFunc = func(ctx context.Context, options *model.PortfolioListOptions) ([]*model.Portfolio, int, error) {
					if options.Page != 1 {
						t.Errorf("期望默认 Page = 1，实际为 %d", options.Page)
					}
					return mockPortfolios, len(mockPortfolios), nil
				}
			},
			wantCount: 3,
			wantTotal: 3,
			wantErr:   false,
		},
		{
			name: "默认分页参数 - PageSize < 1",
			options: &model.PortfolioListOptions{
				Page:     1,
				PageSize: 0,
			},
			setup: func(m *mockPortfolioRepository) {
				m.listFunc = func(ctx context.Context, options *model.PortfolioListOptions) ([]*model.Portfolio, int, error) {
					if options.PageSize != 12 {
						t.Errorf("期望默认 PageSize = 12，实际为 %d", options.PageSize)
					}
					return mockPortfolios, len(mockPortfolios), nil
				}
			},
			wantCount: 3,
			wantTotal: 3,
			wantErr:   false,
		},
		{
			name: "列表查询失败",
			options: &model.PortfolioListOptions{
				Page:     1,
				PageSize: 10,
			},
			setup: func(m *mockPortfolioRepository) {
				m.listFunc = func(ctx context.Context, options *model.PortfolioListOptions) ([]*model.Portfolio, int, error) {
					return nil, 0, errors.New("database error")
				}
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mock := &mockPortfolioRepository{}
			if tt.setup != nil {
				tt.setup(mock)
			}
			svc := NewService(mock)

			got, err := svc.List(ctx, tt.options)
			if (err != nil) != tt.wantErr {
				t.Errorf("List() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !tt.wantErr {
				if got == nil {
					t.Error("List() 返回 nil，期望非 nil")
					return
				}
				if len(got.List) != tt.wantCount {
					t.Errorf("List() 返回 %d 项，期望 %d 项", len(got.List), tt.wantCount)
				}
				if got.Total != tt.wantTotal {
					t.Errorf("List() Total = %d，期望 %d", got.Total, tt.wantTotal)
				}
				// 验证 tier 筛选
				if tt.verifyTier != "" {
					for _, p := range got.List {
						if p.Tier != tt.verifyTier {
							t.Errorf("List() 返回项的 Tier = %s，期望 %s", p.Tier, tt.verifyTier)
						}
					}
				}
			}
		})
	}
}

// TestPortfolioService_ListByTier 专门测试 tier 筛选功能
func TestPortfolioService_ListByTier(t *testing.T) {
	ctx := context.Background()

	// 模拟数据 - 包含所有 tier 类型
	allPortfolios := []*model.Portfolio{
		{ID: "1", Title: "精选项目 A", Tier: "featured", SortOrder: 1},
		{ID: "2", Title: "精选项目 B", Tier: "featured", SortOrder: 2},
		{ID: "3", Title: "推荐项目 A", Tier: "recommended", SortOrder: 1},
		{ID: "4", Title: "推荐项目 B", Tier: "recommended", SortOrder: 2},
		{ID: "5", Title: "普通项目 A", Tier: "normal", SortOrder: 1},
		{ID: "6", Title: "普通项目 B", Tier: "normal", SortOrder: 2},
		{ID: "7", Title: "无层级项目", Tier: "", SortOrder: 3}, // 空 tier
	}

	tests := []struct {
		name        string
		tier        string
		expectCount int
	}{
		{
			name:        "筛选 featured 层级",
			tier:        "featured",
			expectCount: 2,
		},
		{
			name:        "筛选 recommended 层级",
			tier:        "recommended",
			expectCount: 2,
		},
		{
			name:        "筛选 normal 层级",
			tier:        "normal",
			expectCount: 2,
		},
		{
			name:        "空 tier 筛选",
			tier:        "",
			expectCount: 7, // 返回所有
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mock := &mockPortfolioRepository{
				listFunc: func(ctx context.Context, options *model.PortfolioListOptions) ([]*model.Portfolio, int, error) {
					// 验证 tier 参数
					if options.Tier != tt.tier {
						t.Errorf("期望 Tier = %s，实际为 %s", tt.tier, options.Tier)
					}

					// 如果 tier 为空，返回全部
					if options.Tier == "" {
						return allPortfolios, len(allPortfolios), nil
					}

					// 按 tier 筛选
					var filtered []*model.Portfolio
					for _, p := range allPortfolios {
						if p.Tier == options.Tier {
							filtered = append(filtered, p)
						}
					}
					return filtered, len(filtered), nil
				},
			}

			svc := NewService(mock)

			resp, err := svc.List(ctx, &model.PortfolioListOptions{
				Page:     1,
				PageSize: 100,
				Tier:     tt.tier,
			})

			if err != nil {
				t.Errorf("List() error = %v", err)
				return
			}

			if resp == nil {
				t.Error("List() 返回 nil")
				return
			}

			if len(resp.List) != tt.expectCount {
				t.Errorf("List() 返回 %d 项，期望 %d 项", len(resp.List), tt.expectCount)
			}

			// 验证所有返回项的 tier 都正确
			if tt.tier != "" {
				for _, p := range resp.List {
					if p.Tier != tt.tier {
						t.Errorf("返回项 Tier = %s，期望 %s (ID: %s, Title: %s)", p.Tier, tt.tier, p.ID, p.Title)
					}
				}
			}
		})
	}
}

// TestPortfolioService_GetByID 测试获取单个作品
func TestPortfolioService_GetByID(t *testing.T) {
	ctx := context.Background()

	tests := []struct {
		name    string
		id      string
		setup   func(*mockPortfolioRepository)
		wantErr bool
	}{
		{
			name: "成功获取作品",
			id:   "test-id",
			setup: func(m *mockPortfolioRepository) {
				m.getByIDFunc = func(ctx context.Context, publicID string) (*model.Portfolio, error) {
					return &model.Portfolio{
						ID:      publicID,
						Title:   "测试项目",
						Overview: "# 测试概览\n这是一个测试项目",
						Challenge: "## 挑战\n这是一个挑战",
						Solution:  "### 解决方案\n这是解决方案",
					}, nil
				}
			},
			wantErr: false,
		},
		{
			name: "作品不存在",
			id:   "non-existent",
			setup: func(m *mockPortfolioRepository) {
				m.getByIDFunc = func(ctx context.Context, publicID string) (*model.Portfolio, error) {
					return nil, errors.New("not found")
				}
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mock := &mockPortfolioRepository{}
			if tt.setup != nil {
				tt.setup(mock)
			}
			svc := NewService(mock)

			got, err := svc.GetByID(ctx, tt.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetByID() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !tt.wantErr && got == nil {
				t.Error("GetByID() 返回 nil，期望非 nil")
			}
		})
	}
}

// TestPortfolioService_Update 测试更新作品
func TestPortfolioService_Update(t *testing.T) {
	ctx := context.Background()

	newTitle := "更新后的标题"
	featuredTier := "featured"

	tests := []struct {
		name    string
		id      string
		req     *model.UpdatePortfolioRequest
		setup   func(*mockPortfolioRepository)
		wantErr bool
	}{
		{
			name: "成功更新作品",
			id:   "test-id",
			req: &model.UpdatePortfolioRequest{
				Title: &newTitle,
			},
			setup: func(m *mockPortfolioRepository) {
				m.updateFunc = func(ctx context.Context, publicID string, req *model.UpdatePortfolioRequest) (*model.Portfolio, error) {
					return &model.Portfolio{
						ID:    publicID,
						Title: *req.Title,
					}, nil
				}
			},
			wantErr: false,
		},
		{
			name: "更新 tier",
			id:   "test-id",
			req: &model.UpdatePortfolioRequest{
				Tier: &featuredTier,
			},
			setup: func(m *mockPortfolioRepository) {
				m.updateFunc = func(ctx context.Context, publicID string, req *model.UpdatePortfolioRequest) (*model.Portfolio, error) {
					return &model.Portfolio{
						ID:   publicID,
						Tier: *req.Tier,
					}, nil
				}
			},
			wantErr: false,
		},
		{
			name: "更新失败 - 作品不存在",
			id:   "non-existent",
			req: &model.UpdatePortfolioRequest{
				Title: &newTitle,
			},
			setup: func(m *mockPortfolioRepository) {
				m.updateFunc = func(ctx context.Context, publicID string, req *model.UpdatePortfolioRequest) (*model.Portfolio, error) {
					return nil, errors.New("not found")
				}
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mock := &mockPortfolioRepository{}
			if tt.setup != nil {
				tt.setup(mock)
			}
			svc := NewService(mock)

			got, err := svc.Update(ctx, tt.id, tt.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("Update() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !tt.wantErr && got == nil {
				t.Error("Update() 返回 nil，期望非 nil")
			}
		})
	}
}

// TestPortfolioService_Delete 测试删除作品
func TestPortfolioService_Delete(t *testing.T) {
	ctx := context.Background()

	tests := []struct {
		name    string
		id      string
		setup   func(*mockPortfolioRepository)
		wantErr bool
	}{
		{
			name: "成功删除作品",
			id:   "test-id",
			setup: func(m *mockPortfolioRepository) {
				m.deleteFunc = func(ctx context.Context, publicID string) error {
					return nil
				}
			},
			wantErr: false,
		},
		{
			name: "删除失败 - 作品不存在",
			id:   "non-existent",
			setup: func(m *mockPortfolioRepository) {
				m.deleteFunc = func(ctx context.Context, publicID string) error {
					return errors.New("not found")
				}
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mock := &mockPortfolioRepository{}
			if tt.setup != nil {
				tt.setup(mock)
			}
			svc := NewService(mock)

			err := svc.Delete(ctx, tt.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("Delete() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

// TestPortfolioService_GetStats 测试获取统计数据
func TestPortfolioService_GetStats(t *testing.T) {
	ctx := context.Background()

	tests := []struct {
		name    string
		setup   func(*mockPortfolioRepository)
		wantErr bool
	}{
		{
			name: "成功获取统计数据",
			setup: func(m *mockPortfolioRepository) {
				m.getStatsFunc = func(ctx context.Context) (*model.PortfolioStatsResponse, error) {
					return &model.PortfolioStatsResponse{
						Total: 10,
						ByType: map[string]int{
							"frontend": 5,
							"backend":  3,
							"fullstack": 2,
						},
					}, nil
				}
			},
			wantErr: false,
		},
		{
			name: "获取统计数据失败",
			setup: func(m *mockPortfolioRepository) {
				m.getStatsFunc = func(ctx context.Context) (*model.PortfolioStatsResponse, error) {
					return nil, errors.New("database error")
				}
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mock := &mockPortfolioRepository{}
			if tt.setup != nil {
				tt.setup(mock)
			}
			svc := NewService(mock)

			got, err := svc.GetStats(ctx)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetStats() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !tt.wantErr && got == nil {
				t.Error("GetStats() 返回 nil，期望非 nil")
			}
		})
	}
}

// TestPortfolioService_UpdateSort 测试批量更新排序
func TestPortfolioService_UpdateSort(t *testing.T) {
	ctx := context.Background()

	sorts := map[string]int{
		"id1": 1,
		"id2": 2,
		"id3": 3,
	}

	tests := []struct {
		name    string
		sorts   map[string]int
		setup   func(*mockPortfolioRepository)
		wantErr bool
	}{
		{
			name:  "成功更新排序",
			sorts: sorts,
			setup: func(m *mockPortfolioRepository) {
				m.updateSortFunc = func(ctx context.Context, sorts map[string]int) error {
					return nil
				}
			},
			wantErr: false,
		},
		{
			name:  "更新排序失败",
			sorts: sorts,
			setup: func(m *mockPortfolioRepository) {
				m.updateSortFunc = func(ctx context.Context, sorts map[string]int) error {
					return errors.New("database error")
				}
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mock := &mockPortfolioRepository{}
			if tt.setup != nil {
				tt.setup(mock)
			}
			svc := NewService(mock)

			err := svc.UpdateSort(ctx, tt.sorts)
			if (err != nil) != tt.wantErr {
				t.Errorf("UpdateSort() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

// TestConvertPortfolioList 测试列表转换函数
func TestConvertPortfolioList(t *testing.T) {
	input := []*model.Portfolio{
		{ID: "1", Title: "项目 A"},
		{ID: "2", Title: "项目 B"},
		{ID: "3", Title: "项目 C"},
	}

	result := convertPortfolioList(input)

	if len(result) != len(input) {
		t.Errorf("convertPortfolioList() 返回 %d 项，期望 %d 项", len(result), len(input))
	}

	for i, item := range result {
		if item.ID != input[i].ID {
			t.Errorf("convertPortfolioList()[%d].ID = %s，期望 %s", i, item.ID, input[i].ID)
		}
		if item.Title != input[i].Title {
			t.Errorf("convertPortfolioList()[%d].Title = %s，期望 %s", i, item.Title, input[i].Title)
		}
	}
}
