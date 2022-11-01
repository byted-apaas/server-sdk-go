// Copyright 2022 ByteDance Ltd. and/or its affiliates
// SPDX-License-Identifier: MIT

package file

import (
	"bytes"
	"context"
	"io"
	"os"
	"time"

	cExceptions "github.com/byted-apaas/server-common-go/exceptions"
	"github.com/byted-apaas/server-sdk-go/common/structs"
	"github.com/byted-apaas/server-sdk-go/request"
)

type IFile interface {
	// UploadByPath 精度是 second, 0 值不过期
	UploadByPath(ctx context.Context, fileName, path string, expire time.Duration) (*structs.Attachment, error)
	// UploadByReader 精度是 second, 0 值不过期
	UploadByReader(ctx context.Context, fileName string, reader io.Reader, expire time.Duration) (*structs.Attachment, error)
	// UploadByBuffer 精度是 second, 0 值不过期
	UploadByBuffer(ctx context.Context, fileName string, buffer []byte, expire time.Duration) (*structs.Attachment, error)
	Download(ctx context.Context, fileID string) ([]byte, error)
}

type File struct {
	appCtx *structs.AppCtx
}

func NewFile(s *structs.AppCtx) *File {
	return &File{appCtx: s}
}

func (f *File) UploadByPath(ctx context.Context, fileName, path string, expire time.Duration) (*structs.Attachment, error) {
	fileReader, err := os.Open(path)
	if err != nil {
		return nil, cExceptions.InvalidParamError("UploadByPath failed, err: %v", err)
	}
	return request.GetInstance(ctx).UploadFile(ctx, f.appCtx, fileName, fileReader, expire)
}

func (f *File) UploadByReader(ctx context.Context, fileName string, reader io.Reader, expire time.Duration) (*structs.Attachment, error) {
	return request.GetInstance(ctx).UploadFile(ctx, f.appCtx, fileName, reader, expire)
}

func (f *File) UploadByBuffer(ctx context.Context, fileName string, buffer []byte, expire time.Duration) (*structs.Attachment, error) {
	return request.GetInstance(ctx).UploadFile(ctx, f.appCtx, fileName, bytes.NewReader(buffer), expire)
}

func (f *File) Download(ctx context.Context, fileID string) ([]byte, error) {
	return request.GetInstance(ctx).DownloadFile(ctx, f.appCtx, fileID)
}
