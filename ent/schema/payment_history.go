package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// PAYMENT_HISTORY holds the schema definition for the PAYMENT_HISTORY entity.
type PAYMENT_HISTORY struct {
	ent.Schema
}

// Fields of the PAYMENT_HISTORY.
func (PAYMENT_HISTORY) Fields() []ent.Field {
	return []ent.Field{
		field.Int("id").
			SchemaType(map[string]string{
				dialect.MySQL: "INT",
			}),
		field.String("employee_gid").
			SchemaType(map[string]string{
				dialect.MySQL: "VARCHAR(200)",
			}),
		field.Time("created_at").
			SchemaType(map[string]string{
				dialect.MySQL: "DATETIME",
			}),
		field.String("created_by").
			SchemaType(map[string]string{
				dialect.MySQL: "VARCHAR(200)",
			}),
	}
}

// Edges of the PAYMENT_HISTORY.
func (PAYMENT_HISTORY) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("payment_history_rec", EMPLOYEE.Type).
			Ref("payment_history").
			Unique(),
	}
}

func (PAYMENT_HISTORY) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "payment_history"},
	}
}
