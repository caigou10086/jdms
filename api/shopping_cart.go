package api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"jdms/common"
	"jdms/utils"
)

type Cart struct {
	Id   string `json:"id"`
	Name string `json:"name"`
	common.Resp
}

// GetCart 购物车商品选中
func (o Cart) GetCart(c *gin.Context) {
	o.MakeContext(c)

	o.Ok(SelectCart())
}

func SelectCart() map[string]string {
	o1 := utils.PostJson(common.CartApi+"/api?functionId=pcCart_jc_cartCheckAll&appid=JDC_mall_cart", map[string]any{}, "https://cart.jd.com/")
	str2json := utils.Str2json(o1)
	data := str2json["resultData"].(map[string]interface{})

	var skuIdMap = make(map[string]string)

	if data["cartInfo"] == nil {
		return skuIdMap
	}
	cartInfo := data["cartInfo"].(map[string]interface{})["plusFloorRequestParam"].(map[string]interface{})["preferentialGuidelinesInfo"].(map[string]interface{})["plusSkuInfo"].(map[string]interface{})
	// 获取选中的物品
	plusSkuList := cartInfo["plusSkuList"].([]interface{})

	for _, item := range plusSkuList {
		s := item.(map[string]interface{})
		skuId := s["skuId"].(string)
		skuIdMap[skuId] = skuId
	}
	return skuIdMap
}

// AddCart 添加购物车
func (o Cart) AddCart(c *gin.Context) {
	o.MakeContext(c)
	skuId := c.Query("skuId")
	if "" == skuId {
		o.Error("skuId不能为空")
	} else {
		AddCart(skuId)
		o.Ok(nil)
	}

}

func AddCart(skuId string) {
	skuIdMap := SelectCart()
	// 判断是否skuIdMap是否存在skuId
	if skuIdMap[skuId] != "" {
		fmt.Println("该商品已经存在")
		return
	}
	utils.Get(common.AddCardApi + "?pid=" + skuId + "&pcount=1&ptype=1")
}
