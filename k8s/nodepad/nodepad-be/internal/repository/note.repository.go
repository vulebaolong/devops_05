package repository

import (
	"context"
	"fmt"

	"nodepad-be/ent"
	entnote "nodepad-be/ent/note"
	"nodepad-be/internal/common/pagination"

	"entgo.io/ent/dialect/sql"
	"github.com/google/uuid"
)

type NoteRepository interface {
	Create(ctx context.Context, userID *uuid.UUID, shareKey string, title *string, content string) (*ent.Note, error)
	FindByShareKey(ctx context.Context, shareKey string) (*ent.Note, error)
	FindByIDAndUserID(ctx context.Context, id, userID uuid.UUID) (*ent.Note, error)
	ListByUserID(ctx context.Context, userID uuid.UUID, opts ...pagination.ListOption) ([]*ent.Note, error)
	CountByUserID(ctx context.Context, userID uuid.UUID) (int, error)
	Update(ctx context.Context, id, userID uuid.UUID, title, content *string) (*ent.Note, error)
	Delete(ctx context.Context, id, userID uuid.UUID) (*ent.Note, error)
}

type noteRepository struct {
	db *ent.Client
}

func NewNoteRepository(db *ent.Client) NoteRepository {
	return &noteRepository{db: db}
}

func (r *noteRepository) Create(ctx context.Context, userID *uuid.UUID, shareKey string, title *string, content string) (*ent.Note, error) {
	item, err := GetClient(ctx, r.db).Note.Create().
		SetNillableUserID(userID).
		SetShareKey(shareKey).
		SetNillableTitle(title).
		SetContent(content).
		Save(ctx)
	if err != nil {
		return nil, fmt.Errorf("create note: %w", err)
	}
	return item, nil
}

func (r *noteRepository) FindByShareKey(ctx context.Context, shareKey string) (*ent.Note, error) {
	item, err := GetClient(ctx, r.db).Note.Query().Where(entnote.ShareKeyEQ(shareKey)).Only(ctx)
	if err != nil {
		return nil, fmt.Errorf("find note by share key: %w", err)
	}
	return item, nil
}

func (r *noteRepository) FindByIDAndUserID(ctx context.Context, id, userID uuid.UUID) (*ent.Note, error) {
	item, err := GetClient(ctx, r.db).Note.Query().
		Where(entnote.IDEQ(id), entnote.UserIDEQ(userID)).
		Only(ctx)
	if err != nil {
		return nil, fmt.Errorf("find note: %w", err)
	}
	return item, nil
}

func (r *noteRepository) ListByUserID(ctx context.Context, userID uuid.UUID, opts ...pagination.ListOption) ([]*ent.Note, error) {
	options := pagination.ListOptions{}
	for _, opt := range opts {
		opt(&options)
	}
	items, err := GetClient(ctx, r.db).Note.Query().
		Where(entnote.UserIDEQ(userID)).
		Order(entnote.ByUpdatedAt(sql.OrderDesc())).
		Offset(options.Offset).
		Limit(options.Limit).
		All(ctx)
	if err != nil {
		return nil, fmt.Errorf("list notes: %w", err)
	}
	return items, nil
}

func (r *noteRepository) CountByUserID(ctx context.Context, userID uuid.UUID) (int, error) {
	total, err := GetClient(ctx, r.db).Note.Query().Where(entnote.UserIDEQ(userID)).Count(ctx)
	if err != nil {
		return 0, fmt.Errorf("count notes: %w", err)
	}
	return total, nil
}

func (r *noteRepository) Update(ctx context.Context, id, userID uuid.UUID, title, content *string) (*ent.Note, error) {
	item, err := r.FindByIDAndUserID(ctx, id, userID)
	if err != nil {
		return nil, err
	}
	update := GetClient(ctx, r.db).Note.UpdateOneID(item.ID)
	if title != nil {
		update.SetTitle(*title)
	}
	if content != nil {
		update.SetContent(*content)
	}
	item, err = update.Save(ctx)
	if err != nil {
		return nil, fmt.Errorf("update note: %w", err)
	}
	return item, nil
}

func (r *noteRepository) Delete(ctx context.Context, id, userID uuid.UUID) (*ent.Note, error) {
	item, err := r.FindByIDAndUserID(ctx, id, userID)
	if err != nil {
		return nil, err
	}
	if err := GetClient(ctx, r.db).Note.DeleteOneID(id).Exec(ctx); err != nil {
		return nil, fmt.Errorf("delete note: %w", err)
	}
	return item, nil
}
