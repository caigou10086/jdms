package main

import (
	"bufio"
	"fmt"
	"github.com/robfig/cron/v3"
	"jdms/api"
	"jdms/chrome"
	"jdms/common"
	"os"
	"strconv"
	"strings"
	"time"
)

func init1() {

	chrome.RunChrome()

	user := common.GetUserData()
	//user.Cookie = "__jdu=16457740896051002181490; pinId=o_IRM5iweoNFvBVqWivTFbV9-x-f3wj7; shshshfpa=258283c7-3118-b3e1-2e5d-94645c9aec16-1636905733; pin=jd_67c22fc00eb96; unick=jd_176089ule; _tp=8PjFcphNQ31r6RMZ8XcDx6chH1w3Jc5GtO9skwctKKc%3D; _pst=jd_67c22fc00eb96; PCSYCityID=CN_330000_330100_0; user-key=32b0fd47-cd0d-48c8-b914-5cf845d84ae0; ipLocation=%u6d77%u5357; answer-code=\"\"; unpl=JF8EAKxnNSttXkNVVhwGHkEVTlpSW1RYG0QGazdWVQ8PTlYGTAMYRxh7XlVdXhRLFB9sZRRUVVNLVA4bBisSEXtfVVdcDkgWA25XNVNdUQYVV1YyGBIgS1xkXloNTxEFamYMXFpZT1QEGgEbGxBOVVVuXDhMFwpfVzVRXVlKUAYdBBgaIEptVl9cDUMVC2xuBGQWNkoZBRwHHxQWTlxdVloJTxcCbmQFXV1dQ1U1GjIY; __jdv=122270672|kong|t_308072010_|tuiguang|790b654b446748aab45ab0cf432f02d9|1671793876343; mt_xid=V2_52007VwMVV1leVF8eQBFbBGMDE1NeWFtaGkEYbFBhAhAAVFxXRhceTggZYlcSAkFRW11IVU4OAGUHQFZYUFNYSHkaXQZkHxJRQVtXSx9OElwAbAYSYl1oUmoXTBldBGQCG1VaaFdZGUo%3D; areaId=15; TrackID=1y9X2w8TLA8EyjfhzK9qbJ1zJx_XhpHYc7ls0CkCNWPek-mJjkY-QdqpXV9yoxniVkCckB4gouZjRx6IZ8tRrCuo5dUtbKDWOXjRFYyt28bPIvF0HxGWNN0403MpiC96d; ceshi3.com=103; ipLoc-djd=23-2121-22467-54685.6092928932; shshshfpb=r7myqL%2F8%2FoqeHBf1JcYPC8A%3D%3D; __jdc=122270672; ip_cityCode=1213; shshshfp=a5fa9023f1816f26688f5ca71952ba7a; cn=1; detailedAdd_coord=20.0204,110.302; token=d949e3e2b67cf196cf4b6befe1af2a81,3,928868; __tk=2wfD2zSCqAqnruanrAWW1wx5rDbuKUgW2zMArc2u2za,3,928868; jsavif=1; __jda=122270672.16457740896051002181490.1645774090.1671952868.1671963061.54; joyya=0.1671964232.30.0e0acyv; thor=77FC6469580DBFC6D043FF56D90C878F67705AEDB689AE99A680A806AF911715403BA184896B037F5762A16B325A775CA4B9D08440DD7129CD7C6274DAC6F3C345558BE695F07ECCFBECCE7E8436B40C278E9960D37DB4FE6D7EDD9BACA84FCEFFC6ECF97433AA731BAF51A40A1BEE327B621BDAED5A055DE750DA29089621DA9041E8AE5517EA6AC30D477D60C9F1832FA28B8980EBD74B0EA274ED07C9D563; shshshsID=642f9490090f1e2935fbb296b3514c09_2_1671964267201; __jdb=122270672.3.16457740896051002181490|54.1671963061; JSESSIONID=D8AED87D02CC6506ADDA9B95844F44D3.s1; 3AB9D23F7A4B3C9B=RSPAXW5H2DMCQLB63IPNPA4IKQ27AF63EGQWWHCKXJITFYH3OGFCG5QF2QIFE4G3ZUKFYZZLFHMPNTNUFTIWIHBB2I"
	//user.IsLogin = true
	//user.Eid = "RSPAXW5H2DMCQLB63IPNPA4IKQ27AF63EGQWWHCKXJITFYH3OGFCG5QF2QIFE4G3ZUKFYZZLFHMPNTNUFTIWIHBB2I"
	//user.Fp = "2c295e94b3755f778b09394efa1ef0ff"
	var stkId string
	fmt.Print("需要抢购的商品代码：")
	fmt.Scanln(&stkId)
	fmt.Println("开始添加购物车")
	api.AddCart(stkId)

	fmt.Print("请输入秒杀时间(例：2022-01-02 00:00:00)：")
	reader := bufio.NewReader(os.Stdin) // 标准输入输出
	t, _ := reader.ReadString('\n')     // 回车结束
	t = strings.TrimSpace(t)            // 去除最后一个空格

	strTime := t
	millisecond := "000"
	if len(t) == 23 {
		strTime = t[0 : len(t)-4]
		millisecond = t[len(t)-3 : len(t)]
	} else if len(t) != 19 {
		fmt.Println("时间格式不正确")
		return
	}
	local, _ := time.LoadLocation("Asia/Shanghai")
	formatTime, _ := time.ParseInLocation("2006-01-02 15:04:05", strTime, local)

	user.StartTime = formatTime
	user.Millisecond = millisecond
	test := func() {
		// 做毫秒级别的判断
		nanosecond := time.Now().Nanosecond()
		user := common.GetUserData()
		m, _ := strconv.Atoi(user.Millisecond)
		m = m * 1000000
		m = (m - nanosecond) / 1000000
		if m > 0 {
			fmt.Println("等待毫秒级别的时间", m)
			time.Sleep(time.Duration(m))
		}

		print("定时任务开始执行")
		api.StartOrder()
	}
	// 创建定时任务
	c := cron.New(cron.WithSeconds())
	cronStr := strconv.Itoa(formatTime.Second()) + " " + strconv.Itoa(formatTime.Minute()) + " " + strconv.Itoa(formatTime.Hour()) + " " + strconv.Itoa(formatTime.Day()) + " * ?"
	c.AddFunc(cronStr, test)
	fmt.Println("定时任务创建成功，下次执行时间：", formatTime)
	c.Start()
	select {}
}
func main() {

	// 运行时初始化
	init1()
	//chrome.RunChrome()

	//order := api.Order{}
	//cart := api.Cart{}
	//r := gin.New()
	//r.Use(handler.Recover)
	//r.GET("/getOrder", order.GetList)
	//r.GET("/getCart", cart.GetCart)
	//r.GET("/addCart", cart.AddCart)
	//r.Run()

}
