package chrome

import (
	"context"
	"fmt"
	"github.com/chromedp/cdproto/network"
	"github.com/chromedp/chromedp"
	"jdms/common"
	"jdms/utils"
	"log"
	"os"
	"regexp"
	"time"
)

func RunChrome() {
	// 禁用chrome headless
	opts := append(chromedp.DefaultExecAllocatorOptions[:],
		chromedp.Flag("headless", true),
		chromedp.UserAgent(`Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/108.0.0.0 Safari/537.36`),
	)
	allocCtx, cancel := chromedp.NewExecAllocator(context.Background(), opts...)
	defer cancel()

	ctx, cancel := chromedp.NewContext(
		allocCtx,
		chromedp.WithLogf(log.Printf),
	)
	defer cancel()

	ctx, cancel = context.WithTimeout(ctx, 40*time.Second)
	defer cancel()

	err := chromedp.Run(ctx,
		chromedp.Navigate(`https://passport.jd.com/uc/login`),
		// 等待三秒
		chromedp.Sleep(3*time.Second),
		GetQr(),
	)
	if err != nil {
		log.Fatal(err)
	}

}

// GetQr 获取二维码
func GetQr() chromedp.Tasks {
	return chromedp.Tasks{
		chromedp.WaitVisible(`.qrcode-img`, chromedp.ByQuery),
		JdLogin(),
	}
}

// JdLogin 登陆操作
func JdLogin() chromedp.ActionFunc {
	return func(ctx context.Context) (err error) {
		// 存储数据
		var code []byte
		var eid string
		var fp string
		var userName string

		fmt.Println("获取登陆二维码")
		// 二维码获取
		if err = chromedp.Screenshot(`.qrcode-img`, &code, chromedp.ByQuery).Do(ctx); err != nil {
			return
		}
		fmt.Println("二维码获取成功....")
		chromedp.Value(`#eid`, &eid, chromedp.ByID).Do(ctx)
		chromedp.Value(`#sessionId`, &fp, chromedp.ByID).Do(ctx)
		fmt.Println(utils.PrintQRCode(code))

		c := make(chan string)
		// 获取cookie
		go func() {
			chromedp.WaitVisible(`#ttbar-myjd`, chromedp.ByID).Do(ctx)
			c <- getCookies(ctx)
			close(c)
		}()
		// 超时关闭
		go func() {
			// 20秒超时
			time.Sleep(20 * time.Second)
			_, ok := <-c
			if ok {
				c <- ""
				close(c)
			}
		}()

		var cookie string
		fmt.Println("等待扫码登陆....")
		cookie = <-c
		if cookie == "" {
			fmt.Println("长时间未登陆，程序退出....")
			os.Exit(-1)
		}
		// 获取用户名称
		chromedp.OuterHTML(`#ttbar-login > div.dt.cw-icon > a`, &userName, chromedp.ByQuery).Do(ctx)
		const MATCH = ">(.{1,30})<"
		pattern := regexp.MustCompile(MATCH)
		p := pattern.FindAllStringSubmatch(userName, -1)
		userName = p[0][1]
		fmt.Println("登陆成功：" + userName)

		// 存入缓存
		cache := common.GetUserData()
		cache.Cookie = cookie
		cache.IsLogin = true
		cache.Name = userName
		cache.Eid = eid
		cache.Fp = fp
		return
	}
}

// 获取cookie
func getCookies(ctx context.Context) string {
	// cookies的获取对应是在devTools的network面板中
	// 1. 获取cookies
	cookies, err := network.GetAllCookies().Do(ctx)
	if err != nil {
		//return nil
	}
	var cookie string
	for i, item := range cookies {
		cookie += item.Name + "=" + item.Value
		if len(cookie)-1 != i {
			cookie += "; "
		}
	}
	return cookie
}
