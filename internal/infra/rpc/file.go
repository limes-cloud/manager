package rpc

import (
	"github.com/limes-cloud/kratosx"
	"github.com/limes-cloud/manager/api/manager/errors"
	file "github.com/limes-cloud/resource/api/resource/file/v1"
)

const (
	Resource = "Resource"
)

type File struct {
}

func NewFile() *File {
	return &File{}
}

func (i File) client(ctx kratosx.Context) (file.FileClient, error) {
	conn, err := kratosx.MustContext(ctx).GrpcConn(Resource)
	if err != nil {
		return nil, errors.ResourceServerError()
	}
	return file.NewFileClient(conn), nil
}

func (i File) GetFileURL(ctx kratosx.Context, sha string) string {
	client, err := i.client(ctx)
	if err != nil {
		ctx.Logger().Warnw("msg", "connect resource server error", "err", err.Error())
		return ""
	}
	reply, err := client.GetFile(ctx, &file.GetFileRequest{Sha: &sha})
	if err != nil {
		ctx.Logger().Warnw("msg", "get resource file error", "err", err.Error())
		return ""
	}
	return reply.Url
}
