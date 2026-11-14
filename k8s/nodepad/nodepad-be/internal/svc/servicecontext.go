// Code scaffolded by goctl. Safe to edit.
// goctl 1.10.1

package svc

import (
	"nodepad-be/internal/common/data"
	"nodepad-be/internal/config"
	"nodepad-be/internal/middleware"
	"nodepad-be/internal/repository"

	"github.com/zeromicro/go-zero/rest"
)

type ServiceContext struct {
	Config     config.Config
	Data       *data.Data
	UserRepo   repository.UserRepository
	NoteRepo   repository.NoteRepository
	CookieAuth rest.Middleware
}

func NewServiceContext(c config.Config) (*ServiceContext, func(), error) {
	d, cleanup, err := data.NewData(c.DatabaseDSN)
	if err != nil {
		return nil, nil, err
	}

	return &ServiceContext{
			Config:     c,
			Data:       d,
			UserRepo:   repository.NewUserRepository(d.Db),
			NoteRepo:   repository.NewNoteRepository(d.Db),
			CookieAuth: middleware.NewCookieAuthMiddleware(c.Auth.AccessSecret).Handle,
		},
		cleanup,
		nil
}
