package toolkit

import (
	"context"
	"encoding/json"
	"fmt"
	"net/url"
	"path"

	"github.com/androidsr/sc-go/model"
	"github.com/androidsr/sc-go/shttp"
)

type HttpClient struct {
	ctx     context.Context
	baseUrl string
	headers map[string]string
}

func NewHttpClient(baseUrl string) *HttpClient {
	headers := make(map[string]string, 0)
	return &HttpClient{baseUrl: baseUrl, headers: headers}
}

func (a *HttpClient) Startup(ctx context.Context) {
	a.ctx = ctx
}

func (a *HttpClient) SetHeader(headers map[string]string) {
	a.headers = headers
	fmt.Println(a.headers)
}

func (a *HttpClient) Post(url string, data interface{}) interface{} {
	fmt.Println(url, data)
	bs, err := shttp.Post(a.baseUrl+url, shttp.JSON, a.headers, data)
	if err != nil {
		return model.NewFailDefault(err.Error())
	}
	return model.NewOK(a.toJSON(bs))
}

func (a *HttpClient) Put(url string, data interface{}) interface{} {
	fmt.Println(url, data)
	bs, err := shttp.Put(path.Join(a.baseUrl, url), shttp.JSON, a.headers, data)
	if err != nil {
		return model.NewFailDefault(err.Error())
	}
	return model.NewOK(a.toJSON(bs))
}

func (a *HttpClient) Delete(url string, data interface{}) interface{} {
	fmt.Println(url, data)
	bs, err := shttp.Delete(path.Join(a.baseUrl, url), shttp.JSON, a.headers, data)
	if err != nil {
		return model.NewFailDefault(err.Error())
	}
	return model.NewOK(a.toJSON(bs))
}

func (a *HttpClient) PostForm(url string, data interface{}) interface{} {
	fmt.Println(url, data)
	bs, err := shttp.PostForm(path.Join(a.baseUrl, url), shttp.JSON, a.headers, data)
	if err != nil {
		return model.NewFailDefault(err.Error())
	}
	return model.NewOK(a.toJSON(bs))
}

func (a *HttpClient) Get(url string, data map[string]string) interface{} {
	paramsStr := mapToURLParams(data)
	bs, err := shttp.Get(a.baseUrl+url+"?"+paramsStr, shttp.JSON, a.headers)
	if err != nil {
		return model.NewFailDefault(err.Error())
	}
	return model.NewOK(a.toJSON(bs))
}

func (a *HttpClient) toJSON(data []byte) interface{} {
	var raw json.RawMessage
	err := json.Unmarshal(data, &raw)
	if err != nil {
		return nil
	}

	// 尝试解析为对象
	var obj map[string]interface{}
	if err := json.Unmarshal(raw, &obj); err == nil {
		return obj
	}

	// 尝试解析为数组
	var arr []interface{}
	if err := json.Unmarshal(raw, &arr); err == nil {
		return arr
	}

	return nil
}

func mapToURLParams(params map[string]string) string {
	var result string
	for key, value := range params {
		if result != "" {
			result += "&" // 添加连接符
		}
		result += fmt.Sprintf("%s=%s", url.QueryEscape(key), url.QueryEscape(value))
	}
	return result
}
