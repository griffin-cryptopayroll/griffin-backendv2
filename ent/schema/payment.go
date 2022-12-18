package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// PAYMENT holds the schema definition for the Payment entity.
type PAYMENT struct {
	ent.Schema
}

// Fields of the PAYMENT.
func (PAYMENT) Fields() []ent.Field {
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
		field.Time("payment_scheduled").
			SchemaType(map[string]string{
				dialect.MySQL: "DATETIME",
			}).Optional(),
		field.Time("payment_executed").
			SchemaType(map[string]string{
				dialect.MySQL: "DATETIME",
			}).Optional(),
		field.Float("payment_amount").
			SchemaType(map[string]string{
				dialect.MySQL: "FLOAT",
			}),
		field.Int("crypto_currency_id").
			SchemaType(map[string]string{
				dialect.MySQL: "INT",
			}).
			Optional(),
	}
}

// Edges of the PAYMENT.
func (PAYMENT) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("payment_from_employer", EMPLOYER.Type).
			Ref("employer_of_payment").
			Field("employer_id").
			Unique(),
		edge.From("payment_from_employee", EMPLOYEE.Type).
			Ref("employee_of_payment").
			Field("employee_id").
			Unique(),
		edge.From("payment_from_currency", CRYPTO_CURRENCY.Type).
			Ref("currency_of_payment").
			Field("crypto_currency_id").
			Unique(),
	}
}

func (PAYMENT) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "payment"},
	}
}
