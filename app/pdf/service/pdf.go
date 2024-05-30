package service

import (
	"context"
	"github.com/adnpa/gpdf/enums"
	"github.com/adnpa/gpdf/pkg/logger"
	"github.com/adnpa/gpdf/proto/pb"
	"github.com/pdfcpu/pdfcpu/pkg/api"
	"github.com/pdfcpu/pdfcpu/pkg/pdfcpu/model"
	"github.com/pdfcpu/pdfcpu/pkg/pdfcpu/types"
	"go.uber.org/zap"
	"os"
	"path"
	"strconv"
	"sync"
)

type PdfSrv struct{}

var PdfSrvIns *PdfSrv
var PdfSrvOnce sync.Once

func GetPdfSrv() *PdfSrv {
	PdfSrvOnce.Do(func() {
		PdfSrvIns = &PdfSrv{}
	})
	return PdfSrvIns
}

func (p *PdfSrv) Split(ctx context.Context, req *pb.SplitReq, resp *pb.SplitResp) (err error) {
	inputFile := path.Join(enums.InputPath, req.File)
	outputPath := path.Join(enums.OutputPath, req.File)
	err = os.MkdirAll(outputPath, 0755)
	if err != nil {
		return err
	}

	if req.Span != 0 { // 每n页拆分
		err := api.SplitFile(inputFile, outputPath, 2, &model.Configuration{})
		if err != nil {
			logger.Logger().Warn("Split error", zap.Error(err))
			resp.Code = enums.ERROR
			return err
		}
	} else if req.Page != nil { //指定页数拆分
		pageInt := make([]int, len(req.Page))
		for i, v := range req.Page {
			pageInt[i], _ = strconv.Atoi(v)
		}
		err := api.SplitByPageNrFile(inputFile, outputPath, pageInt, &model.Configuration{})
		if err != nil {
			logger.Logger().Warn("Split error", zap.Error(err))
			resp.Code = enums.ERROR
			return err
		}
	} else { // 每页拆分
		err := api.SplitFile(inputFile, outputPath, 2, &model.Configuration{})

		if err != nil {
			logger.Logger().Warn("Split error", zap.Error(err))
			resp.Code = enums.ERROR
			return err
		}
	}
	resp.Code = enums.SUCCESS

	return
}

func (p *PdfSrv) Merge(ctx context.Context, req *pb.MergeReq, resp *pb.MergeResp) (err error) {
	outfile := path.Join(enums.OutputPath, req.Files[0])
	for i := 0; i < len(req.Files); i++ {
		req.Files[i] = path.Join(enums.InputPath, req.Files[i])
	}
	err = api.MergeAppendFile(req.Files, outfile, false, nil)
	resp.Code = enums.SUCCESS
	return
}

func (p *PdfSrv) AddWaterMark(ctx context.Context, req *pb.AddWaterMarkReq, resp *pb.AddWaterMarkResp) (err error) {
	inFile := path.Join(enums.InputPath, req.File)
	outFile := path.Join(enums.OutputPath, req.File)

	wm, err := api.TextWatermark(req.Text, "", false, false, types.POINTS)
	if err != nil {
		return
	}

	err = api.AddWatermarksFile(inFile, outFile, nil, wm, nil)
	if err != nil {
		return
	}
	resp.Code = enums.SUCCESS
	return
}
