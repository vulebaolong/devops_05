package schema

import (
	"nodepad-be/ent/softdelete"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// User holds the schema definition for the User entity.
type User struct {
	ent.Schema
}

// Fields of the User.
func (User) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).
			Default(func() uuid.UUID {
				id, _ := uuid.NewV7()
				return id
			}),

		field.String("email").NotEmpty().MaxLen(255).Unique().StructTag(`json:"email"`),
		field.String("password").NotEmpty().Sensitive(),

		field.Time("createdAt").Default(time.Now).Immutable(),
		field.Time("updatedAt").Default(time.Now).UpdateDefault(time.Now),
	}
}

// Edges of the User.
func (User) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("Note", Note.Type),
	}
}

func (User) Mixin() []ent.Mixin {
	return []ent.Mixin{
		softdelete.SoftDeleteMixin{},
	}
}
