package base

import (
	"base/client/pkg/comm"
	"base/client/pkg/validator"
	"base/model/client"
	"base/model/imp/serverUser"
	"base/pkg/app"
	"client/request"
	"client/response"
	"github.com/gin-gonic/gin"
)

// @Tags base
// @Summary 获取国内地区数据列表
// @accept json
// @Produce json
// @Param id query string true "-1 省 其他为上级ID"
// @Success 200 {object} comm.Response
// @Router /base/region [get]
func GetRegion(ctx *gin.Context) {
	var form request.Id
	err := ctx.ShouldBind(&form)
	if err != nil {
		comm.TipResponse(ctx, err, 0, nil)
		return
	}
	if form.ID == -1 {
		form.ID = 0
	}
	res, err := client.ServerUserSys().UserRegionList(serverUser.UserRegion{ParentId: int32(form.ID)})
	if err != nil {
		comm.TipResponse(ctx, err, 0, res)
	} else {
		comm.NewResponse(app.Success, &res).End(ctx)
	}
}

// @Tags base
// @Summary 根据经纬度返回所在城市信息
// @accept json
// @Produce json
// @Param data body request.UserRegionLngLat true "date"
// @Success 200 {object} comm.Response
// @Router /base/regionLngLat [post]
func GetRegionLngLat(ctx *gin.Context) {
	var form request.UserRegionLngLat
	if err := validator.Bind(ctx, &form); !err.IsValid() {
		comm.ValidatorResponse(ctx, err)
		return
	}
	res, err := client.ServerUserSys().UserRegionLngLat(serverUser.UserRegion{Loc: []float64{form.Lng, form.Lat}})
	if err != nil {
		comm.TipResponse(ctx, err, 0, res)
	} else {
		var Roles []response.UserRegion
		err = app.Unmarshal(res.List, &Roles)
		treeMap := make(map[int32][]response.UserRegion)
		for _, v := range Roles {
			treeMap[v.ParentId] = append(treeMap[v.ParentId], v)
		}
		menus := treeMap[0]
		for i := 0; i < len(menus); i++ {
			err = getChildrenGetUserMenuList(&menus[i], treeMap)
		}
		comm.NewResponse(app.Success, &menus).End(ctx)
	}
}

func getChildrenGetUserMenuList(menu *response.UserRegion, treeMap map[int32][]response.UserRegion) (err error) {
	menu.Children = treeMap[menu.ID]
	for i := 0; i < len(menu.Children); i++ {
		err = getChildrenGetUserMenuList(&menu.Children[i], treeMap)
	}
	return err
}
