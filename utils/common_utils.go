package utils

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/makiuchi-d/gozxing"
	"github.com/makiuchi-d/gozxing/qrcode"
	goQrcode "github.com/skip2/go-qrcode"
	"image"
	"strings"
)

func Str2json(s string) map[string]interface{} {
	str := strings.Replace(s, "'", "\"", -1)
	str = strings.Replace(str, "\n", "", -1)

	var dat map[string]interface{}
	if err := json.Unmarshal([]byte(str), &dat); err != nil {
		fmt.Println(dat)
	}
	return dat
}

// PrintQRCode 将图片输出到终端
func PrintQRCode(code []byte) string {
	// 解码
	img, _, err := image.Decode(bytes.NewReader(code))
	if err != nil {
		panic(err)
	}
	// 转换
	bmp, err := gozxing.NewBinaryBitmapFromImage(img)
	if err != nil {
		panic(err)
	}
	// 获取对象
	res, err := qrcode.NewQRCodeReader().Decode(bmp, nil)
	if err != nil {
		panic(err)
	}
	qr, err := goQrcode.New(res.String(), goQrcode.High)
	if err != nil {
		panic(err)
	}
	// 输出数据
	return qr.ToSmallString(false)
}
