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
		field.Time("created_at").Default(time.Now),
		field.Time("updated_at").
			Default(time.Now).
			UpdateDefault(time.Now),

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

		// --- 类型与状态 ---
		field.Enum("project_type").
			Values("frontend", "vibecoding", "fullstack", "miniprogram", "app", "other").
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
	}
}

// Edges of the Portfolio.
func (Portfolio) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("technologies", PortfolioTechnology.Type),
	}
}
