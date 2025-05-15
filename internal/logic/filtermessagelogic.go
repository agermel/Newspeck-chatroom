package logic

import (
	"context"

	"newspeak-chat/internal/svc"
	"newspeak-chat/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type FilterMessageLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewFilterMessageLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FilterMessageLogic {
	return &FilterMessageLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *FilterMessageLogic) FilterMessage(req *types.FilterRequest) (resp *types.FilterResponse, err error) {
	// todo: add your logic here and delete this line

	return
}
