package utils

import (
	"bytes"
	"encoding/json"
	"io"
	"io/ioutil"
	"jdms/common"
	"log"
	"net/http"
	"net/url"
	"strings"
	"unsafe"
)

func Get(url string) (string, error) {
	user := common.GetUserData()
	if !user.IsLogin {
		panic("当前用户没有登陆")
	}
	headers := map[string]string{
		"cookie":     user.Cookie,
		"user-agent": "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/108.0.0.0 Safari/537.36",
	}
	return SendGet(url, headers)
}

// PostJson JSON提交
func PostJson(url string, data map[string]any, referer string) string {
	user := common.GetUserData()
	if !user.IsLogin {
		panic("当前用户没有登陆")
	}
	headers := map[string]string{
		"cookie":       user.Cookie,
		"user-agent":   "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/108.0.0.0 Safari/537.36",
		"Content-Type": "application/json;charset=UTF-8",
	}
	headers["referer"] = referer
	return SendPost(url, data, headers)
}

// PostForm From表单提交
func PostForm(url1 string, data url.Values, referer string) string {
	user := common.GetUserData()
	if !user.IsLogin {
		panic("当前用户没有登陆")
	}
	headers := map[string]string{
		"cookie":       user.Cookie,
		"user-agent":   "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/108.0.0.0 Safari/537.36",
		"Content-Type": "application/x-www-form-urlencoded",
	}
	headers["referer"] = referer
	return SendPost(url1, data, headers)
}

// GetDomain 获取Domain
func GetDomain(url string) string {
	return url[0:strings.Index(url, "/")]
}

// SendGet 发送GET请求
func SendGet(url string, headers map[string]string) (string, error) {
	client := &http.Client{}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Printf("请求报错%v\n", err)
		return "", err
	}

	// 设置请求头
	for key, header := range headers {
		req.Header.Set(key, header)
	}

	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}

	//关闭连接
	defer resp.Body.Close()
	all, _ := io.ReadAll(resp.Body)
	return string(all), nil
}

// SendPost POST请求
func SendPost(url string, info any, headers map[string]string) string {

	bytesData, err := json.Marshal(info)
	if err != nil {
		panic(err)
	}

	reader := bytes.NewReader(bytesData)

	request, err := http.NewRequest("POST", url, reader)
	defer request.Body.Close() //程序在使用完回复后必须关闭回复的主体
	if err != nil {
		panic(err)
	}
	// 请求头设置
	for key, header := range headers {
		request.Header.Set(key, header)
	}

	client := http.Client{}
	resp, err := client.Do(request) //Do 方法发送请求，返回 HTTP 回复
	if err != nil {
		panic(err)
	}

	respBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	str := (*string)(unsafe.Pointer(&respBytes))
	return *str
}
