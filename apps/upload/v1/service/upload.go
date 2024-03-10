package service

import (
	"context"
	"hrapi/apps/upload/v1/model"
	"hrapi/internal/utils"
	"io"
	"mime/multipart"
	"os"
)

func Upload(ctx context.Context) IUpload {
	return &upload{ctx: ctx}
}

type (
	IUpload interface {
		Upload(file *multipart.FileHeader) (out model.UploadRes, err error)
	}

	upload struct {
		ctx context.Context
	}
)

func (u *upload) Upload(file *multipart.FileHeader) (out model.UploadRes, err error) {
	var (
		src      multipart.File
		dir      = "assets/tmp"
		filePath = dir + "/" + file.Filename
		outFile  *os.File
	)
	src, err = file.Open()
	if err != nil {
		return
	}
	defer src.Close()

	// 建立資料夾, 如果不存在的話
	utils.MkdirIfNotExist(dir)
	// 建立文件
	outFile, err = os.Create(filePath)
	if err != nil {
		return
	}
	defer outFile.Close()

	// 將上傳的內容複製到剛剛建立的文件
	_, err = io.Copy(outFile, src)
	if err != nil {
		return
	}

	out.Url = filePath

	return out, nil
}
