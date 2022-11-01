package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// About holds the schema definition for the About entity.
type About struct {
	ent.Schema
}

// Fields of the About.
func (About) Fields() []ent.Field {
	return []ent.Field{
		field.String("title").NotEmpty(),
		field.String("content").Optional(),
	}
}

// Edges of the About.
func (About) Edges() []ent.Edge {
	return nil
}
