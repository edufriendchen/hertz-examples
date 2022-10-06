/*
 * Copyright 2022 CloudWeGo Authors
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

// Code generated by hertz generator.

package user_gorm

import (
	"context"

	"github.com/cloudwego/hertz-examples/bizdemo/hertz_gorm/biz/pack"

	"github.com/cloudwego/hertz-examples/bizdemo/hertz_gorm/biz/dal/mysql"
	user_gorm "github.com/cloudwego/hertz-examples/bizdemo/hertz_gorm/biz/hertz_gen/user_gorm"
	"github.com/cloudwego/hertz-examples/bizdemo/hertz_gorm/biz/model"
	"github.com/cloudwego/hertz/pkg/app"
)

// UpdateUser .
// @router /v1/user/update/:user_id [POST]
func UpdateUser(ctx context.Context, c *app.RequestContext) {
	var err error
	var req user_gorm.UpdateUserRequest
	err = c.BindAndValidate(&req)
	if err != nil {
		c.JSON(200, &user_gorm.CreateUserResponse{Code: user_gorm.Code_ParamInvalid, Msg: err.Error()})
		return
	}

	u := &model.User{}
	u.ID = uint(req.UserID)
	u.Name = req.Name
	u.Gender = int64(req.Gender)
	u.Age = req.Age
	u.Introduce = req.Introduce

	if err = mysql.UpdateUser(u); err != nil {
		c.JSON(200, &user_gorm.CreateUserResponse{Code: user_gorm.Code_DBErr, Msg: err.Error()})
		return
	}

	c.JSON(200, &user_gorm.DeleteUserResponse{Code: user_gorm.Code_Success})
}

// DeleteUser .
// @router /v1/user/delete/:user_id [POST]
func DeleteUser(ctx context.Context, c *app.RequestContext) {
	var err error
	var req user_gorm.DeleteUserRequest
	err = c.BindAndValidate(&req)
	if err != nil {
		c.JSON(200, &user_gorm.CreateUserResponse{Code: user_gorm.Code_ParamInvalid, Msg: err.Error()})
		return
	}
	if err = mysql.DeleteUser(req.UserID); err != nil {
		c.JSON(200, &user_gorm.CreateUserResponse{Code: user_gorm.Code_DBErr, Msg: err.Error()})
		return
	}

	c.JSON(200, &user_gorm.DeleteUserResponse{Code: user_gorm.Code_Success})
}

// QueryUser .
// @router /v1/user/query/ [POST]
func QueryUser(ctx context.Context, c *app.RequestContext) {
	var err error
	var req user_gorm.QueryUserRequest
	err = c.BindAndValidate(&req)
	if err != nil {
		c.JSON(200, &user_gorm.CreateUserResponse{Code: user_gorm.Code_ParamInvalid, Msg: err.Error()})
		return
	}

	users, total, err := mysql.QueryUser(req.Keyword, req.Page, req.PageSize)
	if err != nil {
		c.JSON(200, &user_gorm.CreateUserResponse{Code: user_gorm.Code_DBErr, Msg: err.Error()})
		return
	}
	c.JSON(200, &user_gorm.QueryUserResponse{Code: user_gorm.Code_Success, Users: pack.Users(users), Totoal: total})
}

// CreateUser .
// @router /v1/user/create/ [POST]
func CreateUser(ctx context.Context, c *app.RequestContext) {
	var err error
	var req user_gorm.CreateUserRequest
	err = c.BindAndValidate(&req)
	if err != nil {
		c.JSON(200, &user_gorm.CreateUserResponse{Code: user_gorm.Code_ParamInvalid, Msg: err.Error()})
		return
	}
	if err = mysql.CreateUser([]*model.User{
		{
			Name:      req.Name,
			Gender:    int64(req.Gender),
			Age:       req.Age,
			Introduce: req.Introduce,
		},
	}); err != nil {
		c.JSON(200, &user_gorm.CreateUserResponse{Code: user_gorm.Code_DBErr, Msg: err.Error()})
		return
	}

	resp := new(user_gorm.UpdateUserResponse)
	resp.Code = user_gorm.Code_Success
	c.JSON(200, resp)
}