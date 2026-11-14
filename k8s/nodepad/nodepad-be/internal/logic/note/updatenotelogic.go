// Code scaffolded by goctl. Safe to edit.
// goctl 1.10.1

package note

import (
	"context"
	"strings"

	"nodepad-be/ent"
	"nodepad-be/internal/common/apperror"
	"nodepad-be/internal/common/helper"
	commonlogic "nodepad-be/internal/logic"
	"nodepad-be/internal/svc"
	"nodepad-be/internal/types"

	"github.com/google/uuid"
	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateNoteLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdateNoteLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateNoteLogic {
	return &UpdateNoteLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdateNoteLogic) UpdateNote(req *types.UpdateNoteRequest) (resp *types.NoteData, err error) {
	if req.Title == nil && req.Content == nil {
		return nil, apperror.BadRequest("Cần cung cấp title hoặc content")
	}
	userID, err := helper.UserIDFromContext(l.ctx)
	if err != nil {
		return nil, apperror.Unauthorized()
	}
	id, err := uuid.Parse(req.Id)
	if err != nil {
		return nil, apperror.BadRequest("ID note không hợp lệ")
	}
	if req.Title != nil {
		value := strings.TrimSpace(*req.Title)
		req.Title = &value
	}
	item, err := l.svcCtx.NoteRepo.Update(l.ctx, id, userID, req.Title, req.Content)
	if err != nil {
		if ent.IsNotFound(err) || strings.Contains(err.Error(), "not found") {
			return nil, apperror.NotFound("Không tìm thấy note")
		}
		return nil, err
	}
	result := commonlogic.NoteData(item)
	return &result, nil
}
