package schema

import "entgo.io/ent"

// EMPLOYEE holds the schema definition for the EMPLOYEE entity.
type EMPLOYEE struct {
	ent.Schema
}

// Fields of the EMPLOYEE.
func (EMPLOYEE) Fields() []ent.Field {
	return nil
}

// Edges of the EMPLOYEE.
func (EMPLOYEE) Edges() []ent.Edge {
	return nil
}
