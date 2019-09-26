/*
 * @Descripttion:
 * @version:
 * @Author: yyeiei
 * @Date: 2019-08-13 18:03:16
 * @LastEditors: yyeiei
 * @LastEditTime: 2019-08-16 18:16:15
 */
package v1

import (
	"strconv"
	"time"

	"libra/controllers/api"
	"libra/models"
	"libra/models/constants"
	"libra/pkg/enums"
	"libra/pkg/mapper"
	"libra/pkg/random"

	"github.com/gin-gonic/gin"
)

// @Summary get adjudication
// @Tags adjudication
// @Produce json
// @Param token header string true "token"
// @Param pageSize query {integer} false "pageSize"
// @Success 200 {object} models.Result
// @Router /api/v1/adjudication/history [get]
func History_Get(context *gin.Context) {
	_, pageSize := api.GetPageInfo(context)
	wxAccount := context.MustGet(constants.SessionAccount).(models.WxAccount)

	out := []models.Adjudication{}
	models.X.Where("wx_account_id = ?", wxAccount.Id).Limit(pageSize, 0).Desc("id", "created").Find(&out)

	api.WJson(context, out)
}

// @Summary get adjudication
// @Tags adjudication
// @Produce json
// @Param token header string true "token"
// @Param id query {integer} false "id"
// @Success 200 {object} models.Result
// @Router /api/v1/adjudication [get]
func Adjudication_Get(context *gin.Context) {
	idStr := context.Query("id")
	if len(idStr) < 1 {
		api.WJsonCode(context, enums.ParamsError)
		return
	}
	id, _ := strconv.ParseInt(idStr, 10, 64)
	node := &models.Node{}
	models.X.ID(id).Get(node)

	var options = []models.Node{}
	models.X.Where("parent = ? AND node_type = ? AND status = ?", id, models.OptNode, enums.Normal).Find(&options)

	if len(options) < 1 || node == nil {
		api.WJsonCode(context, enums.DataBlank)
		return
	}
	wxAccount := context.MustGet(constants.SessionAccount).(models.WxAccount)

	option := adjudicate(options)

	entity := &models.Adjudication{
		WxAccountId:   wxAccount.Id,
		DirectoryId:   node.Id,
		DirectoryName: node.Name,
		OptionId:      option.Id,
		OptionName:    option.Name,
		Created:       time.Now(),
	}
	models.X.Insert(entity)

	out := &models.AdjudicationOut{}
	mapper.MapTo(entity, out)

	api.WJson(context, out)
}

func adjudicate(options []models.Node) models.Node {
	var sumWeight float64
	count := len(options)

	if count == 1 {
		return options[0]
	}

	weights := []float64{}
	maps := make(map[float64]models.Node)
	for _, v := range options {
		sumWeight += v.Weight
		weights = append(weights, v.Weight)
		maps[v.Weight] = v
	}

	if sumWeight == 0 {
		return options[random.Int(0, count)]
	}

	res := maps[random.Float64WithWeight(weights)]
	return res
}

// TODO: tracing directory
func tracingDirectory(ids []int64) []int64 {
	parent := &models.Node{Id: ids[0]}
	has, _ := models.X.ID(ids[0]).Get(parent)
	if has {
		ids = append(ids, parent.Id)

		if parent.Parent > 0 {
			ids = append(ids, parent.Parent)
			tracingDirectory(ids)
		}
	}

	return ids
}
