package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// EMPLOY_TYPE holds the schema definition for the EMPLOY_TYPE entity.
type EMPLOY_TYPE struct {
	ent.Schema
}

// Fields of the EMPLOY_TYPE.
func (EMPLOY_TYPE) Fields() []ent.Field {
	return []ent.Field{
		field.Int("id").
			SchemaType(map[string]string{
				dialect.MySQL: "INT",
			}),
		field.String("is_permanent").
			SchemaType(map[string]string{
				dialect.MySQL: "VARCHAR(200)",
			}),
		field.String("pay_freq").
			SchemaType(map[string]string{
				dialect.MySQL: "VARCHAR(200)", // D. W. M
			}),
	}
}

// Edges of the EMPLOY_TYPE.
func (EMPLOY_TYPE) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("employee_type_to", EMPLOYEE.Type),
	}
}

func (EMPLOY_TYPE) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "employ_type"},
	}
}
