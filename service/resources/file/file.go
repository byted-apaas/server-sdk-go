// Copyright 2022 ByteDance Ltd. and/or its affiliates
// SPDX-License-Identifier: MIT

package file

import (
	"bytes"
	"context"
	"io"
	"os"

	cExceptions "github.com/byted-apaas/server-common-go/exceptions"
	"github.com/byted-apaas/server-sdk-go/common/structs"
	"github.com/byted-apaas/server-sdk-go/request"
)

type IFile interface {
	UploadByPath(ctx context.Context, fileName, path string, expireSeconds int64) (*structs.Attachment, error)
	UploadByReader(ctx context.Context, fileName string, reader io.Reader, expireSeconds int64) (*structs.Attachment, error)
	UploadByBuffer(ctx context.Context, fileName string, buffer []byte, expireSeconds int64) (*structs.Attachment, error)
	Download(ctx context.Context, fileID string) ([]byte, error)
}

type File struct {
	appCtx *structs.AppCtx
}

func NewFile(s *structs.AppCtx) *File {
	return &File{appCtx: s}
}

func (f *File) UploadByPath(ctx context.Context, fileName, path string, expireSeconds int64) (*structs.Attachment, error) {
	fileReader, err := os.Open(path)
	if err != nil {
		return nil, cExceptions.InvalidParamError("UploadByPath failed, err: %v", err)
	}
	return request.GetInstance(ctx).UploadFile(ctx, f.appCtx, fileName, fileReader, expireSeconds)
}

func (f *File) UploadByReader(ctx context.Context, fileName string, reader io.Reader, expireSeconds int64) (*structs.Attachment, error) {
	return request.GetInstance(ctx).UploadFile(ctx, f.appCtx, fileName, reader, expireSeconds)
}

func (f *File) UploadByBuffer(ctx context.Context, fileName string, buffer []byte, expireSeconds int64) (*structs.Attachment, error) {
	return request.GetInstance(ctx).UploadFile(ctx, f.appCtx, fileName, bytes.NewReader(buffer), expireSeconds)
}

func (f *File) Download(ctx context.Context, fileID string) ([]byte, error) {
	return request.GetInstance(ctx).DownloadFile(ctx, f.appCtx, fileID)
}
