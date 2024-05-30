package handler

import (
	"fmt"
	"github.com/adnpa/gpdf/app/gateway/rpc"
	"github.com/adnpa/gpdf/enums"
	"github.com/adnpa/gpdf/pkg/utils"
	"github.com/adnpa/gpdf/proto/pb"
	"github.com/gin-gonic/gin"
	"go-micro.dev/v4/logger"
	"mime/multipart"
	"os"
	"path"
	"strconv"
)

// SplitHandler 拆分
func SplitHandler(ctx *gin.Context) {
	var req pb.SplitReq
	req.Page = ctx.PostFormArray("page")
	span, _ := strconv.Atoi(ctx.PostForm("span"))
	req.Span = int64(span)
	formFile, err := ctx.FormFile("file")
	if err != nil {
		ResponseError(ctx, enums.CodeServerBusy)
		return
	}
	id, _ := utils.GetID()
	req.File = path.Join(fmt.Sprint(strconv.Itoa(int(id)), ".pdf"))
	filePath := path.Join(enums.InputPath, req.File)
	file, err := os.Create(filePath)
	if err != nil {
		return
	}
	file.Close()
	err = ctx.SaveUploadedFile(formFile, filePath)
	if err != nil {
		ResponseError(ctx, enums.CodeServerBusy)
		return
	}

	Resp, err := rpc.Split(ctx, &req)
	if err != nil {
		ResponseError(ctx, enums.CodeServerBusy)
		return
	}

	logger.Info("<<<<<<<<<<<<<<拆分成功<<<<<<<<<<<<<<")
	ResponseSuccess(ctx, Resp)
}

// MergeHandler 合并
func MergeHandler(ctx *gin.Context) {
	var req pb.MergeReq

	form, _ := ctx.MultipartForm()
	files := form.File["file[]"]

	for _, file := range files {
		filepath, err := SaveFileWithFileHeader(ctx, file)
		if err != nil {
			return
		}
		req.Files = append(req.Files, filepath)
	}

	Resp, err := rpc.Merge(ctx, &req)
	if err != nil {
		ResponseError(ctx, enums.CodeServerBusy)
		return
	}

	logger.Info("<<<<<<<<<<<<<<合并成功<<<<<<<<<<<<<<")
	ResponseSuccess(ctx, Resp)
}

// AddWatermarkHandler 加水印
func AddWatermarkHandler(ctx *gin.Context) {
	var req pb.AddWaterMarkReq
	req.Pages = ctx.PostFormArray("page")
	req.Text = ctx.PostForm("text")

	filepath, err := SaveFile(ctx)
	req.File = filepath
	if err != nil {
		ResponseError(ctx, enums.CodeServerBusy)
		return
	}

	Resp, err := rpc.AddWatermark(ctx, &req)
	if err != nil {
		ResponseError(ctx, enums.CodeServerBusy)
		return
	}

	logger.Info("<<<<<<<<<<<<<<添加水印成功<<<<<<<<<<<<<<")
	ResponseSuccess(ctx, Resp)
}

// 辅助函数
func SaveFile(ctx *gin.Context) (filename string, err error) {
	formFile, err := ctx.FormFile("file")
	return SaveFileWithFileHeader(ctx, formFile)
}

func SaveFileWithFileHeader(ctx *gin.Context, fh *multipart.FileHeader) (filename string, err error) {

	id, _ := utils.GetID()
	filename = fmt.Sprint(strconv.Itoa(int(id)), ".pdf")
	filepath := path.Join(enums.InputPath, filename)
	file, err := os.Create(filepath)
	if err != nil {
		return
	}
	file.Close()
	err = ctx.SaveUploadedFile(fh, filepath)
	return
}
