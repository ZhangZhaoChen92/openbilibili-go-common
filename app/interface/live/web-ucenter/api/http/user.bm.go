// Code generated by protoc-gen-bm v0.1, DO NOT EDIT.
// source: user.proto

/*
Package http is a generated blademaster stub package.
This code was generated with go-common/app/tool/bmgen/protoc-gen-bm v0.1.

It is generated from these files:
	user.proto
*/
package http

import (
	"context"

	bm "go-common/library/net/http/blademaster"
	"go-common/library/net/http/blademaster/binding"
)

// to suppressed 'imported but not used warning'
var _ *bm.Context
var _ context.Context
var _ binding.StructValidator

var PathUserGetUserInfo = "/live.webucenter.User/get_user_info"

// ==============
// User Interface
// ==============

type UserBMServer interface {
	// 根据uid查询用户信息
	// `midware:"auth"`
	GetUserInfo(ctx context.Context, req *GetInfoReq) (resp *GetInfoResp, err error)
}

var UserSvc UserBMServer

func userGetUserInfo(c *bm.Context) {
	p := new(GetInfoReq)
	if err := c.BindWith(p, binding.Default(c.Request.Method, c.Request.Header.Get("Content-Type"))); err != nil {
		return
	}
	resp, err := UserSvc.GetUserInfo(c, p)
	c.JSON(resp, err)
}

// RegisterUserService Register the blademaster route with middleware map
// midMap is the middleware map, the key is defined in proto
func RegisterUserService(e *bm.Engine, svc UserBMServer, midMap map[string]bm.HandlerFunc) {
	auth := midMap["auth"]
	UserSvc = svc
	e.GET("/xlive/web-ucenter/user/get_user_info", auth, userGetUserInfo)
}

// RegisterUserBMServer Register the blademaster route
func RegisterUserBMServer(e *bm.Engine, server UserBMServer) {
	e.GET("/live.webucenter.User/get_user_info", userGetUserInfo)
}
