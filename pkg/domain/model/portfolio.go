/*
 * @Description: Portfolio 领域模型
 * @Author: 安知鱼
 * @Date: 2025-02-23 18:54:10
 * @LastEditTime: 2025-02-23 18:54:10
 * @LastEditors: 安知鱼
 */
package model

import "time"

// --- 核心领域对象 (Domain Object) ---

// Portfolio 是项目作品集的核心领域模型
type Portfolio struct {
	ID            string    `json:"id"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
	Title         string    `json:"title"`          // 项目标题
	Description   string    `json:"description"`    // 项目描述
	CoverURL      string    `json:"cover_url"`      // 封面图 URL
	ProjectType   string    `json:"project_type"`   // 项目类型：frontend/vibecoding/fullstack/backend/miniprogram/app/uiux/devops/game/3d-model/illustration/other
	Status        string    `json:"status"`         // 状态：developing/completed/archived
	Technologies  []string  `json:"technologies"`   // 技术栈列表
	DemoURL       string    `json:"demo_url"`       // 演示地址
	GithubURL     string    `json:"github_url"`     // GitHub 仓库地址
	Featured      bool      `json:"featured"`       // 是否精选/推荐
	SortOrder     int       `json:"sort_order"`     // 排序权重
	Overview      string    `json:"overview"`       // 项目概览（Markdown 格式）
	OverviewHTML  string    `json:"overview_html,omitempty"`  // 项目概览（HTML 渲染版本）
	Role          string    `json:"role"`           // 作者在项目中的角色
	Duration      string    `json:"duration"`       // 项目持续时间
	Client        string    `json:"client"`         // 客户名称（如有）
	Challenge     string    `json:"challenge"`      // 项目挑战描述（Markdown 格式）
	ChallengeHTML string    `json:"challenge_html,omitempty"` // 项目挑战描述（HTML 渲染版本）
	Solution      string    `json:"solution"`       // 解决方案描述（Markdown 格式）
	SolutionHTML  string    `json:"solution_html,omitempty"`  // 解决方案描述（HTML 渲染版本）
	GalleryImages []string  `json:"gallery_images"` // 项目展示图片列表
}

// --- API 数据传输对象 (Data Transfer Objects) ---

// CreatePortfolioRequest 定义了创建项目作品集的请求体
type CreatePortfolioRequest struct {
	Title         string   `json:"title" binding:"required"`                                                                     // 项目标题
	Description   string   `json:"description" binding:"required"`                                                              // 项目描述
	CoverURL      string   `json:"cover_url"`                                                                                   // 封面图 URL
	ProjectType   string   `json:"project_type" binding:"required,oneof=frontend vibecoding fullstack backend miniprogram app uiux devops game 3d-model illustration other"`    // 项目类型
	Status        string   `json:"status" binding:"omitempty,oneof=developing completed archived"`                              // 状态
	Technologies  []string `json:"technologies"`                                                                                // 技术栈列表
	DemoURL       string   `json:"demo_url"`                                                                                    // 演示地址
	GithubURL     string   `json:"github_url"`                                                                                  // GitHub 仓库地址
	Featured      bool     `json:"featured"`                                                                                    // 是否精选/推荐
	SortOrder     int      `json:"sort_order"`                                                                                  // 排序权重
	Overview      string   `json:"overview"`                                                                                    // 项目概览（富文本）
	Role          string   `json:"role"`                                                                                        // 作者在项目中的角色
	Duration      string   `json:"duration"`                                                                                    // 项目持续时间
	Client        string   `json:"client"`                                                                                      // 客户名称
	Challenge     string   `json:"challenge"`                                                                                   // 项目挑战描述
	Solution      string   `json:"solution"`                                                                                    // 解决方案描述
	GalleryImages []string `json:"gallery_images"`                                                                              // 项目展示图片列表
}

// UpdatePortfolioRequest 定义了更新项目作品集的请求体
// 使用指针类型表示可选更新字段
type UpdatePortfolioRequest struct {
	Title         *string   `json:"title"`
	Description   *string   `json:"description"`
	CoverURL      *string   `json:"cover_url"`
	ProjectType   *string   `json:"project_type" binding:"omitempty,oneof=frontend vibecoding fullstack backend miniprogram app uiux devops game 3d-model illustration other"`
	Status        *string   `json:"status" binding:"omitempty,oneof=developing completed archived"`
	Technologies  *[]string `json:"technologies"`
	DemoURL       *string   `json:"demo_url"`
	GithubURL     *string   `json:"github_url"`
	Featured      *bool     `json:"featured"`
	SortOrder     *int      `json:"sort_order"`
	Overview      *string   `json:"overview"`
	Role          *string   `json:"role"`
	Duration      *string   `json:"duration"`
	Client        *string   `json:"client"`
	Challenge     *string   `json:"challenge"`
	Solution      *string   `json:"solution"`
	GalleryImages *[]string `json:"gallery_images"`
}

// PortfolioResponse 定义了项目作品集信息的标准 API 响应结构
type PortfolioResponse struct {
	ID            string    `json:"id"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
	Title         string    `json:"title"`
	Description   string    `json:"description"`
	CoverURL      string    `json:"cover_url"`
	ProjectType   string    `json:"project_type"`
	Status        string    `json:"status"`
	Technologies  []string  `json:"technologies"`
	DemoURL       string    `json:"demo_url"`
	GithubURL     string    `json:"github_url"`
	Featured      bool      `json:"featured"`
	SortOrder     int       `json:"sort_order"`
	Overview      string    `json:"overview"`
	OverviewHTML  string    `json:"overview_html,omitempty"`   // 概述的 HTML 渲染版本
	Role          string    `json:"role"`
	Duration      string    `json:"duration"`
	Client        string    `json:"client"`
	Challenge     string    `json:"challenge"`
	ChallengeHTML string    `json:"challenge_html,omitempty"`  // 挑战的 HTML 渲染版本
	Solution      string    `json:"solution"`
	SolutionHTML  string    `json:"solution_html,omitempty"`   // 解决方案的 HTML 渲染版本
	GalleryImages []string  `json:"gallery_images"`
}

// PortfolioListOptions 定义了项目作品集列表查询选项
type PortfolioListOptions struct {
	Page        int    `form:"page"`                     // 页码
	PageSize    int    `form:"page_size"`                // 每页数量
	Keyword     string `form:"keyword,omitempty"`        // 模糊搜索标题
	ProjectType string `form:"project_type,omitempty"`   // 按项目类型过滤
	Status      string `form:"status,omitempty"`         // 按状态过滤
	Featured    *bool  `form:"featured,omitempty"`       // 按是否精选过滤
}

// PortfolioListResponse 定义了项目作品集列表的 API 响应结构
type PortfolioListResponse struct {
	List  []Portfolio `json:"list"`
	Total int         `json:"total"`
}

// PortfolioStatsResponse 定义了项目作品集统计数据的响应结构
type PortfolioStatsResponse struct {
	Total           int              `json:"total"`            // 项目总数
	TopTechnologies []TechnologyCount `json:"top_technologies"` // 技术栈统计
	ByType          map[string]int   `json:"by_type"`          // 项目类型分布
	ByStatus        map[string]int   `json:"by_status"`        // 项目状态分布
}

// TechnologyCount 定义了技术栈使用统计项
type TechnologyCount struct {
	Name  string `json:"name"`  // 技术名称
	Count int    `json:"count"` // 使用次数
}

// CreatePortfolioParams 封装了创建项目作品集时需要持久化的所有数据
type CreatePortfolioParams struct {
	Title         string
	Description   string
	CoverURL      string
	ProjectType   string
	Status        string
	Technologies  []string
	DemoURL       string
	GithubURL     string
	Featured      bool
	SortOrder     int
	Overview      string
	Role          string
	Duration      string
	Client        string
	Challenge     string
	Solution      string
	GalleryImages []string
}
