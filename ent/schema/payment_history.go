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
		field.Int("employee_id").
			SchemaType(map[string]string{
				dialect.MySQL: "INT",
			}).
			Optional(),
		field.Int("employer_id").
			SchemaType(map[string]string{
				dialect.MySQL: "INT",
			}).
			Optional(),
		field.Int("currency_id").
			SchemaType(map[string]string{
				dialect.MySQL: "INT",
			}).
			Optional(),
		field.Float("amount"),
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
		edge.From("payment_history_from_employee", EMPLOYEE.Type).
			Ref("employee_of_payment_history").
			Field("employee_id").
			Unique(),
		edge.From("payment_history_from_employer", EMPLOYER.Type).
			Ref("employer_of_payment_history").
			Field("employer_id").
			Unique(),
		edge.From("payment_history_from_currency_id", CRYPTO_CURRENCY.Type).
			Ref("currency_of_payment_history").
			Field("currency_id").
			Unique(),
	}
}

func (PAYMENT_HISTORY) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "payment_history"},
	}
}
