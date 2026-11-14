// Code scaffolded by goctl. Safe to edit.
// goctl 1.10.1

package auth

import (
	"context"
	"strings"

	"nodepad-be/ent"
	"nodepad-be/internal/common/apperror"
	"nodepad-be/internal/common/helper"
	commonlogic "nodepad-be/internal/logic"
	"nodepad-be/internal/svc"
	"nodepad-be/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
	"golang.org/x/crypto/bcrypt"
)

type LoginLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoginLogic {
	return &LoginLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *LoginLogic) Login(req *types.LoginRequest) (resp *types.AuthResponse, err error) {
	item, err := l.svcCtx.UserRepo.FindByEmail(l.ctx, strings.ToLower(strings.TrimSpace(req.Email)))
	if err != nil {
		if ent.IsNotFound(err) || strings.Contains(err.Error(), "not found") {
			return nil, apperror.Unauthorized("Email hoặc mật khẩu không đúng")
		}
		return nil, err
	}
	if bcrypt.CompareHashAndPassword([]byte(item.Password), []byte(req.Password)) != nil {
		return nil, apperror.Unauthorized("Email hoặc mật khẩu không đúng")
	}
	token, err := helper.GenerateAccessToken(item.ID.String(), l.svcCtx.Config.Auth.AccessSecret, l.svcCtx.Config.Auth.AccessExpire)
	if err != nil {
		return nil, err
	}
	return &types.AuthResponse{AccessToken: token, TokenType: "Bearer", ExpiresIn: l.svcCtx.Config.Auth.AccessExpire, User: commonlogic.AuthUser(item)}, nil
}
