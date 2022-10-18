package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect"
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
				dialect.MySQL: "VARCHAR(6)",
			}).
			Unique(),
		field.String("employer_gid").
			SchemaType(map[string]string{
				dialect.MySQL: "VARCHAR(6)",
			}),
		field.String("name").
			SchemaType(map[string]string{
				dialect.MySQL: "VARCHAR(200)",
			}),
		field.String("position").
			SchemaType(map[string]string{
				dialect.MySQL: "VARCHAR(45)",
			}),
		field.String("wallet").
			SchemaType(map[string]string{
				dialect.MySQL: "VARCHAR(45)",
			}),
		field.Float("payroll").
			SchemaType(map[string]string{
				dialect.MySQL: "FLOAT",
			}),
		field.Int("currency").
			SchemaType(map[string]string{
				dialect.MySQL: "INT",
			}),
		field.Time("payday").
			SchemaType(map[string]string{
				dialect.MySQL: "DATETIME",
			}),
		field.Int("employ").
			SchemaType(map[string]string{
				dialect.MySQL: "INT",
			}),
		field.String("email").
			SchemaType(map[string]string{
				dialect.MySQL: "VARCHAR(45)",
			}),
		field.String("work_start").
			SchemaType(map[string]string{
				dialect.MySQL: "VARCHAR(45)",
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

// Edges of the EMPLOYEE.
func (EMPLOYEE) Edges() []ent.Edge {
	return []ent.Edge{
		// From
		edge.From("employee_gets", CRYPTO_CURRENCY.Type).
			Ref("employee_paid").
			Unique(),
		edge.From("employee_type_from", EMPLOY_TYPE.Type).
			Ref("employee_type_to").
			Unique(),
		edge.From("work_for", EMPLOYER_USER_INFO.Type).
			Ref("work_under").
			Unique(),
		// To
		edge.To("payment_history", PAYMENT_HISTORY.Type),
	}
}
