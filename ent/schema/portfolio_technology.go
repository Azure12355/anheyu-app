/*
 * @Description:
 * @Author: 安知鱼
 * @Date: 2025-07-25 11:32:06
 * @LastEditTime: 2025-08-05 10:13:00
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

// PortfolioTechnology holds the schema definition for the PortfolioTechnology entity.
type PortfolioTechnology struct {
	ent.Schema
}

// Annotations of the PortfolioTechnology.
func (PortfolioTechnology) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.WithComments(true),
		schema.Comment("项目技术栈表"),
	}
}

// Mixin of the PortfolioTechnology.
func (PortfolioTechnology) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.SoftDeleteMixin{},
	}
}

// Fields of the PortfolioTechnology.
func (PortfolioTechnology) Fields() []ent.Field {
	return []ent.Field{
		// --- 手动定义基础字段 ---
		field.Uint("id"),

		field.Time("created_at").
			Default(time.Now).
			Immutable().
			Comment("创建时间"),

		field.Time("updated_at").
			Default(time.Now).
			UpdateDefault(time.Now).
			Comment("更新时间"),

		field.String("technology").
			Comment("技术名称").
			NotEmpty().
			MaxLen(100).
			Unique(),
	}
}

// Edges of the PortfolioTechnology.
func (PortfolioTechnology) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("portfolio", Portfolio.Type).
			Ref("technologies").
			Unique(),
	}
}
