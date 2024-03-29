/*
 * @Descripttion:
 * @version:
 * @Author: yyeiei
 * @Date: 2019-08-13 18:03:16
 * @LastEditors: yyeiei
 * @LastEditTime: 2019-08-16 18:11:33
 */
package v1

import (
	"strconv"
	"strings"
	"time"

	"libra/controllers/api"
	"libra/models"
	"libra/models/constants"
	"libra/pkg/enums"
	"libra/pkg/mapper"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

// @Summary get nodes
// @Tags node
// @Produce json
// @Param token header string true "token"
// @Param id query {integer} false "id"
// @Success 200 {array} models.Node
// @Router /api/v1/node [get]
func Node_Get(context *gin.Context) {
	idStr := context.DefaultQuery("id", "0")
	id, _ := strconv.ParseInt(idStr, 10, 64)
	wxAccount := context.MustGet(constants.SessionAccount).(models.WxAccount)

	entitys := []models.Node{}
	models.X.
		Where("wx_account_id = ? AND parent = ? AND status = ?", wxAccount.Id, id, enums.Normal).
		Find(&entitys)

	out := []models.NodeOut{}
	mapper.MapTo(entitys, &out)
	api.WJson(context, out)
}

// @Summary add node
// @Tags node
// @Produce json
// @Param token header string true "token"
// @Param in body models.NodeIn true "in"
// @Success 200 {boolean} true
// @Router /api/v1/node [post]
func Node_Post(context *gin.Context) {
	var in models.NodeIn
	if err := context.ShouldBindBodyWith(&in, binding.JSON); err != nil {
		api.WJsonCodeMsg(context, enums.ParamsError, err.Error())
		return
	}
	in.Name = strings.TrimSpace(in.Name)

	entity := &models.Node{
		Name:     in.Name,
		Parent:   in.Parent,
		NodeType: in.NodeType,
	}
	has, _ := models.X.Exist(entity)
	if has {
		api.WJsonCode(context, enums.DataRepeat)
		return
	}
	wxAccount := context.MustGet(constants.SessionAccount).(models.WxAccount)

	now := time.Now()
	entity.Weight = in.Weight
	entity.WxAccountId = wxAccount.Id
	entity.Status = enums.Normal
	entity.Created = now
	models.X.Insert(entity)

	out := models.NodeOut{}
	mapper.MapTo(entity, &out)

	api.WJson(context, out)
}

// @Summary modify node
// @Tags node
// @Produce json
// @Param token header string true "token"
// @Param in body models.NodePutIn true "in"
// @Success 200 {boolean} true
// @Router /api/v1/node [put]
func Node_Put(context *gin.Context) {
	var in models.NodePutIn
	if err := context.ShouldBindBodyWith(&in, binding.JSON); err != nil {
		api.WJsonCodeMsg(context, enums.ParamsError, err.Error())
		return
	}
	in.Name = strings.TrimSpace(in.Name)

	affected, err := models.X.ID(in.Id).Cols("weight, parent, name, updated").Update(&models.Node{
		Weight:  in.Weight,
		Parent:  in.Parent,
		Name:    in.Name,
		Updated: time.Now(),
	})
	if err != nil {
		api.WJsonCodeMsg(context, enums.Fail, err.Error())
		return
	}
	if affected < 1 {
		api.WJsonCode(context, enums.DataBlank)
		return
	}

	api.WJson(context, true)
}

// @Summary delete node
// @Tags node
// @Produce json
// @Param token header string true "token"
// @Param id query {integer} false "id"
// @Success 200 {boolean} true
// @Router /api/v1/node [delete]
func Node_Delete(context *gin.Context) {
	idStr := context.Query("id")
	if len(idStr) < 1 {
		api.WJsonCode(context, enums.ParamsInvalid)
		return
	}

	id, _ := strconv.ParseInt(idStr, 10, 64)

	affected, err := models.X.ID(id).Cols("status, deleted").Update(&models.Node{
		Status:  enums.Deleted,
		Deleted: time.Now(),
	})
	if err != nil {
		api.WJsonCodeMsg(context, enums.Fail, err.Error())
		return
	}
	if affected < 1 {
		api.WJsonCode(context, enums.DataBlank)
		return
	}

	api.WJson(context, true)
}
