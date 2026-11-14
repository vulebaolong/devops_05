// Code scaffolded by goctl. Safe to edit.
// goctl 1.10.1

package note

import (
	"net/http"

	"nodepad-be/internal/common/apperror"
	"nodepad-be/internal/logic/note"
	"nodepad-be/internal/svc"
	"nodepad-be/internal/types"

	"github.com/zeromicro/go-zero/rest/httpx"
)

func CreateNoteHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.CreateNoteRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, apperror.BadRequest(err.Error()))
			return
		}

		l := note.NewCreateNoteLogic(r.Context(), svcCtx)
		resp, err := l.CreateNote(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
