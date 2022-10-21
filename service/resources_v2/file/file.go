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
	UploadByPath(ctx context.Context, fileName, path string) (*structs.FileUploadResult, error)
	UploadByReader(ctx context.Context, fileName string, reader io.Reader) (*structs.FileUploadResult, error)
	UploadByBuffer(ctx context.Context, fileName string, buffer []byte) (*structs.FileUploadResult, error)
	UploadAvatarByPath(ctx context.Context, fileName, path string) (*structs.Avatar, error)
	UploadAvatarByReader(ctx context.Context, fileName string, reader io.Reader) (*structs.Avatar, error)
	UploadAvatarByBuffer(ctx context.Context, fileName string, buffer []byte) (*structs.Avatar, error)
	Download(ctx context.Context, fileID string) ([]byte, error)
	DownloadAvatar(ctx context.Context, imageID string) ([]byte, error)
}

type File struct {
	appCtx *structs.AppCtx
}

func NewFile(s *structs.AppCtx) *File {
	return &File{appCtx: s}
}

func (f *File) UploadByPath(ctx context.Context, fileName, path string) (*structs.FileUploadResult, error) {
	fileReader, err := os.Open(path)
	if err != nil {
		return nil, cExceptions.InvalidParamError("UploadByPath failed, err: %v", err)
	}
	return request.GetInstance(ctx).UploadFileV2(ctx, f.appCtx, fileName, fileReader)
}

func (f *File) UploadAvatarByPath(ctx context.Context, fileName, path string) (*structs.Avatar, error) {
	fileReader, err := os.Open(path)
	if err != nil {
		return nil, cExceptions.InvalidParamError("UploadAvatarByPath failed, err: %v", err)
	}
	return request.GetInstance(ctx).UploadAvatar(ctx, f.appCtx, fileName, fileReader)
}

func (f *File) UploadByReader(ctx context.Context, fileName string, reader io.Reader) (*structs.FileUploadResult, error) {
	return request.GetInstance(ctx).UploadFileV2(ctx, f.appCtx, fileName, reader)
}

func (f *File) UploadAvatarByReader(ctx context.Context, fileName string, reader io.Reader) (*structs.Avatar, error) {
	return request.GetInstance(ctx).UploadAvatar(ctx, f.appCtx, fileName, reader)
}

func (f *File) UploadByBuffer(ctx context.Context, fileName string, buffer []byte) (*structs.FileUploadResult, error) {
	return request.GetInstance(ctx).UploadFileV2(ctx, f.appCtx, fileName, bytes.NewReader(buffer))
}

func (f *File) UploadAvatarByBuffer(ctx context.Context, fileName string, buffer []byte) (*structs.Avatar, error) {
	return request.GetInstance(ctx).UploadAvatar(ctx, f.appCtx, fileName, bytes.NewReader(buffer))
}

func (f *File) Download(ctx context.Context, fileID string) ([]byte, error) {
	return request.GetInstance(ctx).DownloadFile(ctx, f.appCtx, fileID)
}

func (f *File) DownloadAvatar(ctx context.Context, imageID string) ([]byte, error) {
	return request.GetInstance(ctx).DownloadAvatar(ctx, f.appCtx, imageID)
}
