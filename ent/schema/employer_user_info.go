package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// EMPLOYER_USER_INFO holds the schema definition for the EMPLOYER_USER_INFO entity.
type EMPLOYER_USER_INFO struct {
	ent.Schema
}

// Fields of the EMPLOYER_USER_INFO.
func (EMPLOYER_USER_INFO) Fields() []ent.Field {
	return []ent.Field{
		field.Int("id").
			SchemaType(map[string]string{
				dialect.MySQL: "INT",
			}),
		field.String("username").
			SchemaType(map[string]string{
				dialect.MySQL: "VARCHAR(20)",
			}),
		field.String("password").
			SchemaType(map[string]string{
				dialect.MySQL: "VARCHAR(20)",
			}),
		field.String("gid").
			SchemaType(map[string]string{
				dialect.MySQL: "VARCHAR(6)",
			}).
			Unique(),
		field.String("corp_name").
			SchemaType(map[string]string{
				dialect.MySQL: "VARCHAR(45)",
			}),
		field.String("corp_email").
			SchemaType(map[string]string{
				dialect.MySQL: "VARCHAR(45)",
			}),
		field.String("wallet").
			SchemaType(map[string]string{
				dialect.MySQL: "VARCHAR(45)",
			}),
		field.Time("created_at").
			SchemaType(map[string]string{
				dialect.MySQL: "DATETIME",
			}),
		field.String("created_by").
			SchemaType(map[string]string{
				dialect.MySQL: "VARCHAR(45)",
			}),
		field.Time("updated_at").
			SchemaType(map[string]string{
				dialect.MySQL: "DATETIME",
			}),
		field.String("updated_by").
			SchemaType(map[string]string{
				dialect.MySQL: "VARCHAR(45)",
			}),
	}
}

// Edges of the EMPLOYER_USER_INFO.
func (EMPLOYER_USER_INFO) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("work_under", EMPLOYEE.Type),
	}
}
