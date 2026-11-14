// Code scaffolded by goctl. Safe to edit.
// goctl 1.10.1

package note

import (
	"context"
	"strings"

	"nodepad-be/ent"
	"nodepad-be/internal/common/apperror"
	commonlogic "nodepad-be/internal/logic"
	"nodepad-be/internal/svc"
	"nodepad-be/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetPublicNoteLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetPublicNoteLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetPublicNoteLogic {
	return &GetPublicNoteLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetPublicNoteLogic) GetPublicNote(req *types.GetPublicNoteRequest) (resp *types.NoteData, err error) {
	item, err := l.svcCtx.NoteRepo.FindByShareKey(l.ctx, strings.TrimSpace(req.ShareKey))
	if err != nil {
		if ent.IsNotFound(err) || strings.Contains(err.Error(), "not found") {
			return nil, apperror.NotFound("Không tìm thấy note")
		}
		return nil, err
	}
	result := commonlogic.NoteData(item)
	return &result, nil
}
