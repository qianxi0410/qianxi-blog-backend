package blogRpc

import (
	"context"
	"errors"
	"qianxi-blog/service/blog/rpc/blogclient"

	"qianxi-blog/service/admin/api/internal/svc"
	"qianxi-blog/service/admin/api/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
)

type DeletePostLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDeletePostLogic(ctx context.Context, svcCtx *svc.ServiceContext) DeletePostLogic {
	return DeletePostLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DeletePostLogic) DeletePost(req types.DeleteReq) (*types.Reply, error) {
	id, page, size := req.Id, req.Page, req.Size

	if id < 1 {
		return nil, errors.New("文章id不合法")
	}

	if (page < 0 || size < 0) && size != -1 {
		return nil, errors.New("文章分页参数不合法")
	}

	_, err := l.svcCtx.BlogRpc.Delete(l.ctx, &blogclient.DeleteReq{Id: id, Page: page, Size: size})
	if err != nil {
		return nil, err
	}

	return &types.Reply{
		Code: 666,
	}, nil
}
