// Code scaffolded by goctl. Safe to edit.
// goctl 1.10.1

package auth

import (
	"context"

	"nodepad-be/internal/common/apperror"
	"nodepad-be/internal/common/helper"
	commonlogic "nodepad-be/internal/logic"
	"nodepad-be/internal/svc"
	"nodepad-be/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type MeLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewMeLogic(ctx context.Context, svcCtx *svc.ServiceContext) *MeLogic {
	return &MeLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *MeLogic) Me() (resp *types.MeResponse, err error) {
	userID, err := helper.UserIDFromContext(l.ctx)
	if err != nil {
		return nil, apperror.Unauthorized()
	}
	item, err := l.svcCtx.UserRepo.FindByID(l.ctx, userID)
	if err != nil {
		return nil, apperror.Unauthorized()
	}
	user := commonlogic.AuthUser(item)
	return &types.MeResponse{Id: user.Id, Email: user.Email, CreatedAt: user.CreatedAt, UpdatedAt: user.UpdatedAt, DeletedAt: user.DeletedAt}, nil
}
