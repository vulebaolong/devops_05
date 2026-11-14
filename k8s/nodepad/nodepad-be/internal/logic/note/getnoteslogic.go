// Code scaffolded by goctl. Safe to edit.
// goctl 1.10.1

package note

import (
	"context"

	"nodepad-be/internal/common/apperror"
	"nodepad-be/internal/common/helper"
	"nodepad-be/internal/common/pagination"
	commonlogic "nodepad-be/internal/logic"
	"nodepad-be/internal/svc"
	"nodepad-be/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetNotesLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetNotesLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetNotesLogic {
	return &GetNotesLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetNotesLogic) GetNotes(req *types.GetNotesRequest) (resp *types.NoteListResponse, err error) {
	userID, err := helper.UserIDFromContext(l.ctx)
	if err != nil {
		return nil, apperror.Unauthorized()
	}
	page, pageSize, offset := pagination.Normalize(req.Page, req.PageSize)
	if pageSize > 100 {
		pageSize = 100
		offset = (page - 1) * pageSize
	}
	total, err := l.svcCtx.NoteRepo.CountByUserID(l.ctx, userID)
	if err != nil {
		return nil, err
	}
	items, err := l.svcCtx.NoteRepo.ListByUserID(l.ctx, userID, pagination.ListOffset(offset), pagination.ListLimit(pageSize))
	if err != nil {
		return nil, err
	}
	resp = &types.NoteListResponse{
		Page: page, PageSize: pageSize, TotalItems: total,
		TotalPages: (total + pageSize - 1) / pageSize,
		Items:      make([]types.NoteData, 0, len(items)),
	}
	for _, item := range items {
		resp.Items = append(resp.Items, commonlogic.NoteData(item))
	}
	return resp, nil
}
