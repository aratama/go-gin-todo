package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// TodoTask holds the schema definition for the TodoTask entity.
type TodoTask struct {
	ent.Schema
}

// Fields of the TodoTask.
func (TodoTask) Fields() []ent.Field {
	return []ent.Field{
		field.String("name"),
		field.Time("created_at").Default(time.Now),
	}
}

// Edges of the TodoTask.
func (TodoTask) Edges() []ent.Edge {
	return nil
}
