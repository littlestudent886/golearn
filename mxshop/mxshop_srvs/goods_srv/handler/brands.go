package handler

import (
	"context"
	"fmt"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
	"mxshop_srvs/goods_srv/global"
	"mxshop_srvs/goods_srv/model"
	"mxshop_srvs/goods_srv/proto"
)

func (s *GoodsServer) BrandList(ctx context.Context, req *proto.BrandFilterRequest) (*proto.BrandListResponse, error) {
	//展示品牌的列表
	brandListResonse := proto.BrandListResponse{}
	var brands []model.Brand
	var count int64
	//result := global.DB.Find(&brands)
	result := global.DB.Scopes(Paginate(int(req.Pages), int(req.PagePerNums))).Find(&brands)

	if result.Error != nil {
		return nil, result.Error
	}

	global.DB.Model(&model.Brand{}).Count(&count)
	fmt.Println(count)
	brandListResonse.Total = int32(count)

	brandResponses := []*proto.BrandInfoResponse{}
	for _, brand := range brands {
		brandResponse := proto.BrandInfoResponse{
			Id:   brand.ID,
			Name: brand.Name,
			Logo: brand.Logo,
		}
		brandResponses = append(brandResponses, &brandResponse)
	}
	brandListResonse.Data = brandResponses
	return &brandListResonse, nil
}

func (s *GoodsServer) CreateBrand(ctx context.Context, req *proto.BrandRequest) (*proto.BrandInfoResponse, error) {
	// 新建品牌
	if result := global.DB.First(&model.Brand{}, "name = ?", req.Name); result.RowsAffected > 0 {
		return nil, status.Errorf(codes.AlreadyExists, "品牌已存在")
	}

	brand := &model.Brand{
		Name: req.Name,
		Logo: req.Logo,
	}
	global.DB.Save(brand)
	return &proto.BrandInfoResponse{Id: brand.ID}, nil
}

func (s *GoodsServer) DeleteBrand(ctx context.Context, req *proto.BrandRequest) (*emptypb.Empty, error) {
	// 删除品牌
	if result := global.DB.Delete(&model.Brand{}, req.Id); result.RowsAffected == 0 {
		return nil, status.Errorf(codes.NotFound, "品牌不存在")
	}
	return &emptypb.Empty{}, nil
}

func (s *GoodsServer) UpdateBrand(ctx context.Context, req *proto.BrandRequest) (*emptypb.Empty, error) {
	brand := model.Brand{}
	if result := global.DB.First(&brand, "id = ?", req.Id); result.RowsAffected == 0 {
		return nil, status.Errorf(codes.NotFound, "品牌不存在")
	}
	if req.Name != "" {
		brand.Name = req.Name
	}
	if req.Logo != "" {
		brand.Logo = req.Logo
	}

	global.DB.Save(&brand)
	return &emptypb.Empty{}, nil
}
