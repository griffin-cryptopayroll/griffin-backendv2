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
		field.Int("crypto_currency_id").
			SchemaType(map[string]string{
				dialect.MySQL: "INT",
			}).
			Optional(),
		field.Time("payday").
			SchemaType(map[string]string{
				dialect.MySQL: "DATETIME",
			}),
		field.Int("employ_type_id").
			SchemaType(map[string]string{
				dialect.MySQL: "INT",
			}).
			Optional(),
		field.String("email").
			SchemaType(map[string]string{
				dialect.MySQL: "VARCHAR(200)",
			}).
			Unique(),
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
		edge.From("employee_from_currency", CRYPTO_CURRENCY.Type).
			Ref("currency_of_employee").
			Field("crypto_currency_id").
			Unique(),
		edge.From("employee_from_employ_type", EMPLOY_TYPE.Type).
			Ref("employ_type_of_employee").
			Field("employ_type_id").
			Unique(),
		edge.From("employee_from_employer", EMPLOYER.Type).
			Ref("employer_of_employee").
			Field("employer_id").
			Unique(),
		// To
		edge.To("employee_of_payment_history", PAYMENT_HISTORY.Type),
		edge.To("employee_of_payment", PAYMENT.Type),
	}
}

func (EMPLOYEE) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "employee"},
	}
}
