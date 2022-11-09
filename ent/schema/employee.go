package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// EMPLOYEE holds the schema definition for the EMPLOYEE entity.
type EMPLOYEE struct {
	ent.Schema
}

// Fields of the EMPLOYEE.
func (EMPLOYEE) Fields() []ent.Field {
	return []ent.Field{
		field.Int("id").
			SchemaType(map[string]string{
				dialect.MySQL: "INT",
			}),
		field.String("gid").
			SchemaType(map[string]string{
				dialect.MySQL: "VARCHAR(200)",
			}).
			Unique(),
		field.Int("employer_id").
			SchemaType(map[string]string{
				dialect.MySQL: "INT",
			}).
			Optional(),
		field.String("name").
			SchemaType(map[string]string{
				dialect.MySQL: "VARCHAR(200)",
			}),
		field.String("position").
			SchemaType(map[string]string{
				dialect.MySQL: "VARCHAR(200)",
			}),
		field.String("wallet").
			SchemaType(map[string]string{
				dialect.MySQL: "VARCHAR(200)",
			}),
		field.Float("payroll").
			SchemaType(map[string]string{
				dialect.MySQL: "FLOAT",
			}),
		field.Int("currency").
			SchemaType(map[string]string{
				dialect.MySQL: "INT",
			}).
			Optional(),
		field.Time("payday").
			SchemaType(map[string]string{
				dialect.MySQL: "DATETIME",
			}),
		field.Int("employ").
			SchemaType(map[string]string{
				dialect.MySQL: "INT",
			}).
			Optional(),
		field.String("email").
			SchemaType(map[string]string{
				dialect.MySQL: "VARCHAR(200)",
			}).Unique(),
		field.String("work_start").
			SchemaType(map[string]string{
				dialect.MySQL: "VARCHAR(200)",
			}),
		field.String("work_ends").
			SchemaType(map[string]string{
				dialect.MySQL: "VARCHAR(45)",
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
				dialect.MySQL: "VARCHAR(45)",
			}),
	}
}

// Edges of the EMPLOYEE.
func (EMPLOYEE) Edges() []ent.Edge {
	return []ent.Edge{
		// From
		edge.From("employee_currency", CRYPTO_CURRENCY.Type).
			Ref("currency_employee").
			Unique().
			Field("currency"),
		edge.From("employee_type_from", EMPLOY_TYPE.Type).
			Ref("employee_type_to").
			Unique().
			Field("employ"),
		edge.From("work_for", EMPLOYER_USER_INFO.Type).
			Ref("work_under").
			Unique().
			Field("employer_id"),
		// To
		edge.To("payment_history", PAYMENT_HISTORY.Type),
	}
}

func (EMPLOYEE) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "employee"},
	}
}
