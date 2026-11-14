// Code scaffolded by goctl. Safe to edit.
// goctl 1.10.1

package note

import (
	"context"
	"crypto/rand"
	"encoding/base64"
	"strings"

	"nodepad-be/internal/common/apperror"
	"nodepad-be/internal/common/helper"
	commonlogic "nodepad-be/internal/logic"
	"nodepad-be/internal/svc"
	"nodepad-be/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateNoteLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCreateNoteLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateNoteLogic {
	return &CreateNoteLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CreateNoteLogic) CreateNote(req *types.CreateNoteRequest) (resp *types.NoteData, err error) {
	if strings.TrimSpace(req.Content) == "" {
		return nil, apperror.BadRequest("Nội dung note không được để trống")
	}
	if req.Title != nil {
		value := strings.TrimSpace(*req.Title)
		req.Title = &value
	}
	// Authentication is optional on this endpoint. A missing, expired or malformed
	// cookie must not prevent a visitor from creating a public note. We only attach
	// the note to an account when the access token can be verified successfully.
	userID, _ := helper.ParseUserIDFromToken(req.Authorization, l.svcCtx.Config.Auth.AccessSecret)
	random := make([]byte, 9)
	if _, err := rand.Read(random); err != nil {
		return nil, err
	}
	shareKey := base64.RawURLEncoding.EncodeToString(random)
	item, err := l.svcCtx.NoteRepo.Create(l.ctx, userID, shareKey, req.Title, req.Content)
	if err != nil {
		return nil, err
	}
	result := commonlogic.NoteData(item)
	return &result, nil
}
