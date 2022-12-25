package api

import (
	"fmt"
	"jdms/common"
	"jdms/utils"
	"net/url"
	"time"
)

// StartOrder 开始下单
func StartOrder() {

	fmt.Println("选中购物车商品")
	// 循环选中购物车
	for i := 0; i < 10; i++ {
		if len(SelectCart()) > 0 {
			break
		}
	}
	fmt.Println("准备下单")
	// 准备下单
	PrepareOrder()
	fmt.Println("准备提交订单")

	// 开始下单
	flg := make(chan bool)
	go start(flg)
	f := <-flg
	if f {
		fmt.Println("下单成功,请尽快付款")
	} else {
		fmt.Println("下单失败")
	}
}

// 循环下单
func start(flg chan bool) {
	// 循环5次
	for i := 0; i < 15; i++ {
		fmt.Print("尝试下单:", time.Now().Format("2006-01-02 15:04:05"))
		order := SubmitOrder()
		info := utils.Str2json(order)
		b := info["success"].(bool)
		if b {
			fmt.Println("下单成功============")
			flg <- true
			close(flg)
			return
		} else {
			// 下单失败
			fmt.Println("下单失败", info["message"])
		}
	}
	flg <- false
	close(flg)
}

// stop 超时关闭
func stop(flg chan bool) {
	// 20秒超时
	time.Sleep(20 * time.Second)
	_, ok := <-flg
	if ok {
		flg <- false
		close(flg)
	}
}

// PrepareOrder 准备下单
func PrepareOrder() string {
	data, _ := utils.Get(common.TradeApi + "/shopping/order/getOrderInfo.action")
	return data
}

// SubmitOrder 提交订单
func SubmitOrder() string {
	user := common.GetUserData()
	//data := "overseaPurchaseCookies=&vendorRemarks=[{\"venderId\":\"643684\",\"remark\":\"\"}]&submitOrderParam.sopNotPutInvoice=false&submitOrderParam.trackID=TestTrackId&presaleStockSign=1&submitOrderParam.ignorePriceChange=0&submitOrderParam.btSupport=0&submitOrderParam.jxj=1&submitOrderParam.zpjd=1"
	//data := "overseaPurchaseCookies=&vendorRemarks=[]&submitOrderParam.sopNotPutInvoice=false&submitOrderParam.trackID=TestTrackId&presaleStockSign=1&submitOrderParam.ignorePriceChange=0&submitOrderParam.btSupport=0&submitOrderParam.eid=" + user.Eid + "&submitOrderParam.fp=" + user.Fp + "&submitOrderParam.jxj=1&submitOrderParam.zpjd=1"
	//data := "overseaPurchaseCookies=&vendorRemarks=[]&submitOrderParam.sopNotPutInvoice=false&submitOrderParam.trackID=TestTrackId&presaleStockSign=1&submitOrderParam.ignorePriceChange=0&submitOrderParam.btSupport=0&submitOrderParam.eid=" + user.Eid + "&submitOrderParam.fp=" + user.Fp + "&submitOrderParam.jxj=1&submitOrderParam.zpjd=1"
	data := url.Values{}
	data.Add("overseaPurchaseCookies", "")
	data.Add("vendorRemarks", "[]")
	data.Add("submitOrderParam.sopNotPutInvoice", "false")
	data.Add("submitOrderParam.trackID", "TestTrackId")
	data.Add("presaleStockSign", "1")
	data.Add("submitOrderParam.ignorePriceChange", "0")
	data.Add("submitOrderParam.btSupport", "0")
	data.Add("submitOrderParam.eid", user.Eid)
	data.Add("submitOrderParam.fp", user.Fp)
	data.Add("submitOrderParam.jxj", "1")
	data.Add("submitOrderParam.zpjd", "1")

	return utils.PostForm(common.TradeApi+"/shopping/order/submitOrder.action?&presaleStockSign=1", data, "https://trade.jd.com/shopping/order/getOrderInfo.action")
}
