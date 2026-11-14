// Code scaffolded by goctl. Safe to edit.
// goctl 1.10.1

package auth

import (
	"context"
	"net/mail"
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

type RegisterLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewRegisterLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RegisterLogic {
	return &RegisterLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *RegisterLogic) Register(req *types.RegisterRequest) (resp *types.AuthResponse, err error) {
	email := strings.ToLower(strings.TrimSpace(req.Email))
	if _, err := mail.ParseAddress(email); err != nil || len(req.Password) < 5 {
		return nil, apperror.BadRequest("Email không hợp lệ hoặc mật khẩu phải có ít nhất 5 ký tự")
	}
	if _, err := l.svcCtx.UserRepo.FindByEmail(l.ctx, email); err == nil {
		return nil, apperror.Conflict("Email đã được sử dụng")
	} else if !ent.IsNotFound(err) && !strings.Contains(err.Error(), "not found") {
		return nil, err
	}
	hash, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}
	item, err := l.svcCtx.UserRepo.Create(l.ctx, email, string(hash))
	if err != nil {
		return nil, err
	}
	token, err := helper.GenerateAccessToken(item.ID.String(), l.svcCtx.Config.Auth.AccessSecret, l.svcCtx.Config.Auth.AccessExpire)
	if err != nil {
		return nil, err
	}
	return &types.AuthResponse{AccessToken: token, TokenType: "Bearer", ExpiresIn: l.svcCtx.Config.Auth.AccessExpire, User: commonlogic.AuthUser(item)}, nil
}
