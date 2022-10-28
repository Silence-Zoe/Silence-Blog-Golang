package schema

import "entgo.io/ent"

// About holds the schema definition for the About entity.
type About struct {
	ent.Schema
}

// Fields of the About.
func (About) Fields() []ent.Field {
	return nil
}

// Edges of the About.
func (About) Edges() []ent.Edge {
	return nil
}
