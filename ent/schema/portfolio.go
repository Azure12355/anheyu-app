/*
 * @Description:
 * @Author: 安知鱼
 * @Date: 2025-07-25 09:51:07
 * @LastEditTime: 2025-08-13 19:01:58
 * @LastEditors: 安知鱼
 */
package schema

import (
	"time"

	"github.com/anzhiyu-c/anheyu-app/ent/schema/mixin"

	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
)

// Portfolio holds the schema definition for the Portfolio entity.
type Portfolio struct {
	ent.Schema
}

// Annotations of the Portfolio.
func (Portfolio) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.WithComments(true),
		schema.Comment("项目作品集表"),
	}
}

// Mixin of the Portfolio.
func (Portfolio) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.SoftDeleteMixin{},
	}
}

// Fields of the Portfolio.
func (Portfolio) Fields() []ent.Field {
	return []ent.Field{
		// --- 基础字段 ---
		field.Uint("id"),
		field.Time("created_at").
			Default(time.Now).
			Immutable().
			Comment("创建时间"),
		field.Time("updated_at").
			Default(time.Now).
			UpdateDefault(time.Now).
			Comment("更新时间"),

		// --- 基本信息 ---
		field.String("title").
			Comment("项目标题").
			NotEmpty().
			MaxLen(200),
		field.Text("description").
			Comment("项目简介").
			Optional(),
		field.String("cover_url").
			Comment("封面图URL").
			Optional(),

		// --- 层级与展示 ---
		field.Enum("tier").
			Values("normal", "recommended", "featured").
			Comment("项目层级: normal-普通, recommended-推荐, featured-精选").
			Default("normal"),

		// --- 类型与状态 ---
		field.Enum("project_type").
			Values(
				"frontend",
				"vibecoding",
				"fullstack",
				"miniprogram",
				"app",
				"uiux",
				"backend",
				"devops",
				"game",
				"3d-model",
				"illustration",
				"other",
			).
			Comment("项目类型").
			Default("other"),
		field.Enum("status").
			Values("developing", "completed", "archived").
			Comment("项目状态").
			Default("developing"),

		// --- 链接 ---
		field.String("demo_url").
			Comment("演示地址").
			Optional().
			MaxLen(500),
		field.String("github_url").
			Comment("GitHub地址").
			Optional().
			MaxLen(500),

		// --- 展示控制 ---
		field.Bool("featured").
			Comment("是否精选").
			Default(false),
		field.Int("sort_order").
			Comment("排序权重").
			Default(0),

		// --- 详细信息 ---
		field.Text("overview").
			Comment("项目概述").
			Optional(),
		field.String("role").
			Comment("担任角色").
			Optional().
			MaxLen(200),
		field.String("duration").
			Comment("项目周期").
			Optional().
			MaxLen(100),
		field.String("client").
			Comment("所属客户").
			Optional().
			MaxLen(200),
		field.Text("challenge").
			Comment("挑战描述").
			Optional(),
		field.Text("solution").
			Comment("解决方案").
			Optional(),

		// --- 图库图片 ---
		field.JSON("gallery_images", []string{}).
			Comment("图库图片URL列表").
			Optional(),
	}
}

// Edges of the Portfolio.
func (Portfolio) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("technologies", PortfolioTechnology.Type),
	}
}

// Indexes of the Portfolio.
func (Portfolio) Indexes() []ent.Index {
	return []ent.Index{
		// 常见查询：按项目类型和状态筛选
		index.Fields("project_type", "status"),

		// 精选查询：按精选状态和状态筛选
		index.Fields("featured", "status"),

		// 层级查询：按层级和状态筛选
		index.Fields("tier", "status"),

		// 排序查询：按排序权重排序
		index.Fields("sort_order"),
	}
}
