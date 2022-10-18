package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// CRYPTO_PRC_SOURCE holds the schema definition for the CRYPTO_PRC_SOURCE entity.
type CRYPTO_PRC_SOURCE struct {
	ent.Schema
}

// Fields of the CRYPTO_PRC_SOURCE.
func (CRYPTO_PRC_SOURCE) Fields() []ent.Field {
	return []ent.Field{
		field.Int("id").
			SchemaType(map[string]string{
				dialect.MySQL: "INT",
			}),
		field.String("name").
			SchemaType(map[string]string{
				dialect.MySQL: "VARCHAR(200)",
			}),
	}
}

// Edges of the CRYPTO_PRC_SOURCE.
func (CRYPTO_PRC_SOURCE) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("price_of", CRYPTO_CURRENCY.Type),
	}
}
