package controllers

import (
	"base/client/pkg/comm"
	"base/client/pkg/middlewares"
	"base/client/pkg/validator"
	"base/model/client"
	"base/model/imp/serverCo"
	"base/pkg/app"
	"client/request"
	"client/response"
	"encoding/json"
	"github.com/gin-gonic/gin"
)

// @Tags co_benefit
// @Summary 获取企业拥有的权益列表
// @Security ApiKeyAuth
// @accept json
// @Produce json
// @Param coid header string true "企业ID"
// @Success 200 {object} comm.Response
// @Router /co/benefit/own [post]
func UserBenefitOwn(ctx *gin.Context) {
	coId, _ := middlewares.GetFastCoIdUid(ctx)
	res, err := client.ServerCoUser().UserBenefitOwn(serverCo.Uid{Coid: coId})
	if err != nil {
		comm.TipResponse(ctx, err, 0, res)
	} else {
		comm.NewResponse(app.Success, res).End(ctx)
	}
	return
}

// @Tags co_benefit
// @Summary 获取企业拥有的权益分页
// @Description 包含自定义
// @Security ApiKeyAuth
// @accept json
// @Produce json
// @Param coid header string true "企业ID"
// @Param data body request.CoUserBenefitDiyPage true "date"
// @Success 200 {object} comm.Response
// @Router /co/benefit/page [post]
func UserBenefitPage(ctx *gin.Context) {
	var form request.CoUserBenefitDiyPage
	if err := validator.Bind(ctx, &form); !err.IsValid() {
		comm.ValidatorResponse(ctx, err)
		return
	}
	var Form serverCo.ReqCoUserBenefitDiy
	if err := app.UnmarshalJson(form, &Form.Page); err != nil {
		comm.TipResponse(ctx, err, 0, nil)
		return
	}
	coId, _ := middlewares.GetFastCoIdUid(ctx)
	Form.Req = &serverCo.CoUserBenefitDiy{Coid: coId, Type: int32(form.Type)}
	res, err := client.ServerCoUser().UserBenefitDiyPage(Form)
	if err != nil {
		comm.TipResponse(ctx, err, 0, nil)
	} else {
		var v interface{}
		err = json.Unmarshal(res.List, &v)
		comm.NewResponse(app.Success, response.PageResult{
			List:     v,
			Total:    res.Count,
			Page:     form.Page,
			PageSize: form.PageSize,
		}).End(ctx)
	}
	return
}

// @Tags co_benefit
// @Summary 获取自定义权益数据
// @Security ApiKeyAuth
// @accept json
// @Produce json
// @Param coid header string true "企业ID"
// @Param data body request.Id true "date"
// @Success 200 {object} comm.Response
// @Router /co/benefit/get [post]
func UserBenefitData(ctx *gin.Context) {
	var form request.Id
	if err := validator.Bind(ctx, &form); !err.IsValid() {
		comm.ValidatorResponse(ctx, err)
		return
	}
	coId, _ := middlewares.GetFastCoIdUid(ctx)
	res, err := client.ServerCoUser().UserBenefitDiyData(serverCo.CoUserBenefitDiy{Id: form.ID, Coid: coId})
	if err != nil {
		comm.TipResponse(ctx, err, 0, res)
	} else {
		comm.NewResponse(app.Success, &res).End(ctx)
	}
}

// 公用修改 删除操作
func BenefitSava(ctx *gin.Context, way string) {
	var form request.CoUserBenefitDiy
	if err := validator.Bind(ctx, &form); !err.IsValid() {
		comm.ValidatorResponse(ctx, err)
		return
	}
	if way == "edit" {
		if err := validator.Get().Var(form.ID, "required"); err != nil {
			comm.TipResponse(ctx, app.Err("ID值异常"), 0, nil)
			return
		}
	}
	var Form serverCo.CoUserBenefitDiy
	if err := app.UnmarshalJson(form, &Form); err != nil {
		comm.TipResponse(ctx, err, 0, nil)
		return
	}
	coId, _ := middlewares.GetFastCoIdUid(ctx)
	Form.Coid = coId
	res, err := client.ServerCoUser().UserBenefitDiySave(Form)
	comm.TipResponse(ctx, err, 0, res)
}

// @Tags co_benefit
// @Summary 添加自定义权益
// @Security ApiKeyAuth
// @accept json
// @Produce json
// @Param coid header string true "企业ID"
// @Param data body request.CoUserBenefitDiy true "date"
// @Success 200 {object} comm.Response
// @Router /co/benefit/add [put]
func UserBenefitAdd(ctx *gin.Context) {
	BenefitSava(ctx, "add")
	return
}

// @Tags co_benefit
// @Summary 修改自定义权益
// @Security ApiKeyAuth
// @accept json
// @Produce json
// @Param coid header string true "企业ID"
// @Param data body request.CoUserBenefitDiy true "date"
// @Success 200 {object} comm.Response
// @Router /co/benefit/edit [put]
func UserBenefitEdit(ctx *gin.Context) {
	BenefitSava(ctx, "edit")
	return
}

// @Tags co_benefit
// @Summary 删除自定义权益
// @Security ApiKeyAuth
// @accept json
// @Produce json
// @Param coid header string true "企业ID"
// @Param data body request.Id true "date"
// @Success 200 {object} comm.Response
// @Router /co/benefit/delete [delete]
func UserBenefitDelete(ctx *gin.Context) {
	var form request.Id
	if err := validator.Bind(ctx, &form); !err.IsValid() {
		comm.ValidatorResponse(ctx, err)
		return
	}
	coId, _ := middlewares.GetFastCoIdUid(ctx)
	res, err := client.ServerCoUser().UserBenefitDiyDelete(serverCo.CoUserBenefitDiy{Id: form.ID, Coid: coId})
	comm.TipResponse(ctx, err, 0, res)
}

// @Tags co_benefitCard
// @Summary 获取企业权益卡 分页
// @Security ApiKeyAuth
// @accept json
// @Produce json
// @Param coid header string true "企业ID"
// @Param data body request.CoUserBenefitPage true "date"
// @Success 200 {object} comm.Response
// @Router /co/benefitCard/page [post]
func UserBenefitCardPage(ctx *gin.Context) {
	var form request.CoUserBenefitPage
	if err := validator.Bind(ctx, &form); !err.IsValid() {
		comm.ValidatorResponse(ctx, err)
		return
	}
	var Form serverCo.ReqCoUserBenefit
	if err := app.UnmarshalJson(form, &Form.Page); err != nil {
		comm.TipResponse(ctx, err, 0, nil)
		return
	}
	coId, _ := middlewares.GetFastCoIdUid(ctx)
	Form.Req = &serverCo.CoUserBenefit{Coid: coId, Name: form.Name}
	res, err := client.ServerCoUser().UserBenefitPage(Form)
	if err != nil {
		comm.TipResponse(ctx, err, 0, nil)
	} else {
		var v struct {
			Count int32       `json:"count"`
			List  interface{} `json:"list"`
		}
		err = json.Unmarshal(res.Body, &v)
		if err != nil {
			comm.TipResponse(ctx, err, 0, res)
		} else {
			comm.NewResponse(app.Success, response.PageResult{
				List:     v.List,
				Total:    v.Count,
				Page:     form.Page,
				PageSize: form.PageSize,
			}).End(ctx)
		}
	}
	return
}

// @Tags co_benefitCard
// @Summary 获取企业权益卡 数据
// @Security ApiKeyAuth
// @accept json
// @Produce json
// @Param coid header string true "企业ID"
// @Param data body request.Id true "date"
// @Success 200 {object} comm.Response
// @Router /co/benefitCard/get [post]
func UserBenefitCardData(ctx *gin.Context) {
	var form request.Id
	if err := validator.Bind(ctx, &form); !err.IsValid() {
		comm.ValidatorResponse(ctx, err)
		return
	}
	coId, _ := middlewares.GetFastCoIdUid(ctx)
	res, err := client.ServerCoUser().UserBenefitData(serverCo.CoUserBenefit{Id: form.ID, Coid: coId})
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

// 公用修改 删除操作
func BenefitCardSave(ctx *gin.Context, way string) {
	var form request.CoUserBenefit
	if err := validator.Bind(ctx, &form); !err.IsValid() {
		comm.ValidatorResponse(ctx, err)
		return
	}
	if way == "edit" {
		if err := validator.Get().Var(form.ID, "required"); err != nil {
			comm.TipResponse(ctx, app.Err("ID值异常"), 0, nil)
			return
		}
	}
	var Form serverCo.SaveCoUserBenefit
	if err := app.UnmarshalJson(form, &Form.Info); err != nil {
		comm.TipResponse(ctx, err, 0, nil)
		return
	}
	if err := app.UnmarshalJson(form.MeetBag, &Form.MeetBag); err != nil {
		comm.TipResponse(ctx, err, 0, nil)
		return
	}
	coId, _ := middlewares.GetFastCoIdUid(ctx)
	Form.Info.Coid = coId
	Form.BenefitBag = app.Struct2Json(form.BenefitBag)
	res, err := client.ServerCoUser().UserBenefitSave(Form)
	comm.TipResponse(ctx, err, 0, res)
}

// @Tags co_benefitCard
// @Summary 添加企业权益卡
// @Security ApiKeyAuth
// @accept json
// @Produce json
// @Param coid header string true "企业ID"
// @Param data body request.CoUserBenefit true "date"
// @Success 200 {object} comm.Response
// @Router /co/benefitCard/add [put]
func UserBenefitCardAdd(ctx *gin.Context) {
	BenefitCardSave(ctx, "add")
	return
}

// @Tags co_benefitCard
// @Summary 修改企业权益卡
// @Security ApiKeyAuth
// @accept json
// @Produce json
// @Param coid header string true "企业ID"
// @Param data body request.CoUserBenefit true "date"
// @Success 200 {object} comm.Response
// @Router /co/benefitCard/edit [put]
func UserBenefitCardEdit(ctx *gin.Context) {
	BenefitCardSave(ctx, "edit")
	return
}

// @Tags co_benefitCard
// @Summary 删除企业权益卡
// @Security ApiKeyAuth
// @accept json
// @Produce json
// @Param coid header string true "企业ID"
// @Param data body request.Id true "date"
// @Success 200 {object} comm.Response
// @Router /co/benefitCard/delete [delete]
func UserBenefitCardDelete(ctx *gin.Context) {
	var form request.Id
	if err := validator.Bind(ctx, &form); !err.IsValid() {
		comm.ValidatorResponse(ctx, err)
		return
	}
	coId, _ := middlewares.GetFastCoIdUid(ctx)
	res, err := client.ServerCoUser().UserBenefitDiyDelete(serverCo.CoUserBenefitDiy{Id: form.ID, Coid: coId})
	comm.TipResponse(ctx, err, 0, res)
}
