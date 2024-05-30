package rpc

import (
	"context"
	"github.com/adnpa/gpdf/enums"
	"github.com/adnpa/gpdf/proto/pb"
)

func Split(ctx context.Context, req *pb.SplitReq) (resp *pb.SplitResp, err error) {
	resp, err = PdfService.Split(ctx, req)
	if err != nil {
		return
	}

	if resp.Code != enums.SUCCESS {
		return
	}

	return
}

func Merge(ctx context.Context, req *pb.MergeReq) (resp *pb.MergeResp, err error) {
	resp, err = PdfService.Merge(ctx, req)
	if err != nil {
		return
	}
	return
}

func AddWatermark(ctx context.Context, req *pb.AddWaterMarkReq) (resp *pb.AddWaterMarkResp, err error) {
	resp, err = PdfService.AddWaterMark(ctx, req)
	if err != nil {
		return
	}
	return
}
