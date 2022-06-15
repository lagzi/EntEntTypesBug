package schema

import (
	"entgo.io/ent"
	edge "entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Picture holds the schema definition for the Picture entity.
type Picture struct {
	ent.Schema
}

// Fields of the Picture.
func (Picture) Fields() []ent.Field {
	return []ent.Field{
		field.String("url"),
	}
}

// Edges of the Picture.
func (Picture) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("user", User.Type).
			Ref("pic"),
	}
}
