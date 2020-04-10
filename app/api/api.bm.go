// Code generated by protoc-gen-bm v0.1, DO NOT EDIT.
// source: api.proto

package api

import (
	"context"

	bm "github.com/go-kratos/kratos/pkg/net/http/blademaster"
	"github.com/go-kratos/kratos/pkg/net/http/blademaster/binding"
)
import goods_service_v1 "easymarket-go-server/common/goods/api"

// to suppressed 'imported but not used warning'
var _ *bm.Context
var _ context.Context
var _ binding.StructValidator

var PathAppGetGoodsList = "/app/getGoodsList"
var PathAppGetGoodsDetail = "/app/getGoodsDetail"

// AppBMServer is the server API for App service.
type AppBMServer interface {
	// 获取商品列表
	GetGoodsList(ctx context.Context, req *goods_service_v1.GoodsReq) (resp *goods_service_v1.GoodsListRes, err error)

	// 获取商品详情
	GetGoodsDetail(ctx context.Context, req *goods_service_v1.GoodsDetailReq) (resp *goods_service_v1.GoodsRes, err error)
}

var AppSvc AppBMServer

func appGetGoodsList(c *bm.Context) {
	p := new(goods_service_v1.GoodsReq)
	if err := c.BindWith(p, binding.Default(c.Request.Method, c.Request.Header.Get("Content-Type"))); err != nil {
		return
	}
	resp, err := AppSvc.GetGoodsList(c, p)
	c.JSON(resp, err)
}

func appGetGoodsDetail(c *bm.Context) {
	p := new(goods_service_v1.GoodsDetailReq)
	if err := c.BindWith(p, binding.Default(c.Request.Method, c.Request.Header.Get("Content-Type"))); err != nil {
		return
	}
	resp, err := AppSvc.GetGoodsDetail(c, p)
	c.JSON(resp, err)
}

// RegisterAppBMServer Register the blademaster route
func RegisterAppBMServer(e *bm.Engine, server AppBMServer) {
	AppSvc = server
	e.GET("/app/getGoodsList", appGetGoodsList)
	e.GET("/app/getGoodsDetail", appGetGoodsDetail)
}
