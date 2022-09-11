package schema

import "entgo.io/ent"

// PAYMENT_HISTORY holds the schema definition for the PAYMENT_HISTORY entity.
type PAYMENT_HISTORY struct {
	ent.Schema
}

// Fields of the PAYMENT_HISTORY.
func (PAYMENT_HISTORY) Fields() []ent.Field {
	return nil
}

// Edges of the PAYMENT_HISTORY.
func (PAYMENT_HISTORY) Edges() []ent.Edge {
	return nil
}
