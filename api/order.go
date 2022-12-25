package api

import (
	"github.com/gin-gonic/gin"
	"jdms/common"
	"jdms/utils"
	"regexp"
)

type Order struct {
	Id   string `json:"id"`
	Name string `json:"name"`
	common.Resp
}

// MATCH 正则表达式
const MATCH = "<a href=\"//item.jd.com/(\\d+).html\" class=\"a-link\" .{0,100} target=\"_blank\" title=\"(.{1,100})\">"

// GetList 获取订单
func (o Order) GetList(c *gin.Context) {
	// 设置上下文
	o.MakeContext(c)
	o.Ok(GetOrderList())
}

// GetOrderList 获取订单列表
func GetOrderList() []Order {
	// 获取订单数据
	o1, _ := utils.Get(common.OrderApi + "/center/list.action?s=1")
	// 正则 解析订单
	pattern := regexp.MustCompile(MATCH)
	data := pattern.FindAllStringSubmatch(o1, -1)
	var orderList []Order
	for i := range data {
		order := Order{
			Id:   data[i][1],
			Name: data[i][2],
		}
		orderList = append(orderList, order)
	}
	return orderList
}
