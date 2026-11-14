package schema

import (
	"nodepad-be/ent/softdelete"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// Note holds the schema definition for the Note entity.
type Note struct {
	ent.Schema
}

// Fields of the Note.
func (Note) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).
			Default(func() uuid.UUID {
				id, _ := uuid.NewV7()
				return id
			}),

		// Nullable vì người chưa đăng nhập vẫn có thể tạo note.
		field.UUID("user_id", uuid.UUID{}).Optional().Nillable().StructTag(`json:"userId"`),

		// Chuỗi ngẫu nhiên dùng trong public URL:
		// /api/notes/public/:shareKey
		field.String("share_key").NotEmpty().MaxLen(20).Unique().Immutable().StructTag(`json:"shareKey"`),
		field.String("title").MaxLen(200).Optional().Nillable().StructTag(`json:"title"`),
		field.Text("content").Default("").StructTag(`json:"content"`),

		field.Time("createdAt").Default(time.Now).Immutable(),
		field.Time("updatedAt").Default(time.Now).UpdateDefault(time.Now),
	}
}

// Edges of the Note.
func (Note) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("user", User.Type).Ref("Note").Field("user_id").Unique(),
	}
}

func (Note) Mixin() []ent.Mixin {
	return []ent.Mixin{
		softdelete.SoftDeleteMixin{},
	}
}
