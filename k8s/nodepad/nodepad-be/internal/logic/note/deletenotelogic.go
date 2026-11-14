// Code scaffolded by goctl. Safe to edit.
// goctl 1.10.1

package note

import (
	"context"
	"strings"
	"time"

	"nodepad-be/ent"
	"nodepad-be/internal/common/apperror"
	"nodepad-be/internal/common/helper"
	commonlogic "nodepad-be/internal/logic"
	"nodepad-be/internal/svc"
	"nodepad-be/internal/types"

	"github.com/google/uuid"
	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteNoteLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDeleteNoteLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteNoteLogic {
	return &DeleteNoteLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DeleteNoteLogic) DeleteNote(req *types.DeleteNoteRequest) (resp *types.NoteData, err error) {
	userID, err := helper.UserIDFromContext(l.ctx)
	if err != nil {
		return nil, apperror.Unauthorized()
	}
	id, err := uuid.Parse(req.Id)
	if err != nil {
		return nil, apperror.BadRequest("ID note không hợp lệ")
	}
	item, err := l.svcCtx.NoteRepo.Delete(l.ctx, id, userID)
	if err != nil {
		if ent.IsNotFound(err) || strings.Contains(err.Error(), "not found") {
			return nil, apperror.NotFound("Không tìm thấy note")
		}
		return nil, err
	}
	item.DeletedAt = time.Now()
	result := commonlogic.NoteData(item)
	return &result, nil
}
