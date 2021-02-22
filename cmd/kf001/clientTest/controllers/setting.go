package controllers

import (
	"base/client/pkg/comm"
	"base/client/pkg/middlewares"
	"base/client/pkg/validator"
	"base/model/client"
	"base/model/imp/serverCo"
	"base/pkg/app"
	"base/tools"
	"client/request"
	"encoding/json"
	"github.com/gin-gonic/gin"
)

// @Tags co_setting
// @Summary 根据配置ID 获取企业配置数据
// @Security ApiKeyAuth
// @accept json
// @Produce json
// @Param coid header string true "企业ID"
// @Param setting_id query int true "int"
// @Success 200 {object} comm.Response
// @Router /co/setting/info [get]
func SettingInfo(ctx *gin.Context) {
	var Form serverCo.CoUserSetting
	Form.Id = tools.StringToInt64(ctx.DefaultQuery("setting_id", ""))
	if err := validator.Get().Var(Form.Id, "required"); err != nil {
		comm.TipResponse(ctx, app.Err("配置ID不存在"), 0, nil)
		return
	}
	coId, _ := middlewares.GetFastCoIdUid(ctx)
	Form.Coid = coId
	res, err := client.ServerCoUser().UserSettingInfo(Form)
	if err != nil {
		comm.TipResponse(ctx, err, 0, res)
	} else {
		var v interface{}
		err = json.Unmarshal(res.Body, &v)
		if err != nil {
			comm.TipResponse(ctx, err, 0, res)
		} else {
			comm.NewResponse(app.Success, &v).End(ctx)
		}
	}
}

// @Tags co_setting
// @Summary 根据菜单ID 获取企业配置数据
// @Security ApiKeyAuth
// @accept json
// @Produce json
// @Param coid header string true "企业ID"
// @Param menu_id query int true "int"
// @Success 200 {object} comm.Response
// @Router /co/setting/menuId [get]
func SettingMenu(ctx *gin.Context) {
	var Form serverCo.CoUserSetting
	Form.MenuId = tools.StringToInt64(ctx.DefaultQuery("menu_id", ""))
	if err := validator.Get().Var(Form.MenuId, "required"); err != nil {
		comm.TipResponse(ctx, app.Err("菜单ID不存在"), 0, nil)
		return
	}
	coId, _ := middlewares.GetFastCoIdUid(ctx)
	Form.Coid = coId
	res, err := client.ServerCoUser().UserSettingInfo(Form)
	if err != nil {
		comm.TipResponse(ctx, err, 0, res)
	} else {
		var v interface{}
		err = json.Unmarshal(res.Body, &v)
		if err != nil {
			comm.TipResponse(ctx, err, 0, res)
		} else {
			comm.NewResponse(app.Success, &v).End(ctx)
		}
	}
}

// @Tags co_setting
// @Summary 修改企业配置数据
// @Security ApiKeyAuth
// @accept json
// @Produce json
// @Param coid header string true "企业ID"
// @Param data body request.CoSysSetting true "date"
// @Success 200 {object} comm.Response
// @Router /co/setting/save [put]
func SettingSave(ctx *gin.Context) {
	var form request.CoSysSetting
	if err := validator.Bind(ctx, &form); !err.IsValid() {
		comm.ValidatorResponse(ctx, err)
		return
	}
	var Form serverCo.SaveCoUserSetting
	coId, _ := middlewares.GetFastCoIdUid(ctx)
	Form.Coid = coId
	Form.SettingData = app.Struct2Json(form.Rule)
	res, err := client.ServerCoUser().UserSettingSave(Form)
	comm.TipResponse(ctx, err, 0, res)
}
