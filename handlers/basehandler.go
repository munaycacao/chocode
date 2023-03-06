package handlers

import (
	"context"
	"github.com/munaycacao/chocode/repository"
)

type BaseHandler struct {
	Repo *repository.Repo
	Ctx context.Context
}

func NewHandler(repo *repository.Repo) *BaseHandler {
	return &BaseHandler{
		Repo: repo,
		Ctx:  context.Background(),
	}
}