package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// EMPLOYER holds the schema definition for the EMPLOYER_USER_INFO entity.
type EMPLOYER struct {
	ent.Schema
}

// Fields of the EMPLOYER_USER_INFO.
func (EMPLOYER) Fields() []ent.Field {
	return []ent.Field{
		field.Int("id").
			SchemaType(map[string]string{
				dialect.MySQL: "INT",
			}),
		field.String("username").
			SchemaType(map[string]string{
				dialect.MySQL: "VARCHAR(200)",
			}),
		field.String("password").
			SchemaType(map[string]string{
				dialect.MySQL: "VARCHAR(200)",
			}),
		field.String("gid").
			SchemaType(map[string]string{
				dialect.MySQL: "VARCHAR(200)",
			}).
			Unique(),
		field.String("corp_name").
			SchemaType(map[string]string{
				dialect.MySQL: "VARCHAR(200)",
			}),
		field.String("corp_email").
			SchemaType(map[string]string{
				dialect.MySQL: "VARCHAR(200)",
			}).Unique(),
		field.String("wallet").
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
		field.Time("updated_at").
			SchemaType(map[string]string{
				dialect.MySQL: "DATETIME",
			}),
		field.String("updated_by").
			SchemaType(map[string]string{
				dialect.MySQL: "VARCHAR(200)",
			}),
	}
}

// Edges of the EMPLOYER_USER_INFO.
func (EMPLOYER) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("employer_of_employee", EMPLOYEE.Type),
		edge.To("employer_of_payment_history", PAYMENT_HISTORY.Type),
	}
}

func (EMPLOYER) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "employer"},
	}
}
