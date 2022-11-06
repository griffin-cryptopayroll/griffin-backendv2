package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
)

// Tr_log holds the schema definition for the Tr_log entity.
type Tr_log struct {
	ent.Schema
}

// Fields of the Tr_log.
func (Tr_log) Fields() []ent.Field {
	return []ent.Field{
		field.Int("id").
			SchemaType(map[string]string{
				dialect.MySQL: "INT",
			}),
		field.String("tr_type").
			SchemaType(map[string]string{
				dialect.MySQL: "VARCHAR(200)",
			}),
		field.Time("created_at").
			SchemaType(map[string]string{
				dialect.MySQL: "DATETIME",
			}),
	}
}

// Edges of the Tr_log.
func (Tr_log) Edges() []ent.Edge {
	return nil
}

func (Tr_log) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "tr_log"},
	}
}
