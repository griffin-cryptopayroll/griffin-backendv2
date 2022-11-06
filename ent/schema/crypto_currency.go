package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// CRYPTO_CURRENCY holds the schema definition for the CRYPTO_CURRENCY entity.
type CRYPTO_CURRENCY struct {
	ent.Schema
}

// Fields of the CRYPTO_CURRENCY.
func (CRYPTO_CURRENCY) Fields() []ent.Field {
	return []ent.Field{
		field.Int("id").
			SchemaType(map[string]string{
				dialect.MySQL: "INT",
			}),
		field.String("ticker").
			SchemaType(map[string]string{
				dialect.MySQL: "VARCHAR(200)",
			}),
		field.Int("source").
			SchemaType(map[string]string{
				dialect.MySQL: "INT",
			}),
	}
}

// Edges of the CRYPTO_CURRENCY.
func (CRYPTO_CURRENCY) Edges() []ent.Edge {
	return []ent.Edge{
		// From
		// Unique to ensure that crypto-currency ticker can have one price source
		edge.From("source_of", CRYPTO_PRC_SOURCE.Type).
			Ref("price_of").
			Unique(),
		// To
		edge.To("employee_paid", EMPLOYEE.Type),
	}
}

func (CRYPTO_CURRENCY) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "crypto_currency"},
	}
}
